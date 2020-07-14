package amzadv

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/syncfuture/go/soauth2"
	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

type APIClient struct {
	OAuth2Config *oauth2.Config
	TokenStore   soauth2.ITokenStore
	AdvURL       string
	ProfileID    string
}

func NewAPIClient(config *oauth2.Config, tokenStore soauth2.ITokenStore) (r *APIClient) {
	r = &APIClient{
		OAuth2Config: config,
		TokenStore:   tokenStore,
	}
	return r
}

// #region Token

func (x *APIClient) getTokenSource() (oauth2.TokenSource, error) {
	token, err := x.TokenStore.GetToken()
	tokenSource := x.OAuth2Config.TokenSource(oauth2.NoContext, token)
	newToken, err := tokenSource.Token()
	if u.LogError(err) {
		return tokenSource, err
	}

	// Save new token if it get refreshed
	if newToken.AccessToken != token.AccessToken {
		x.TokenStore.SaveToken(newToken)
	}
	return tokenSource, err
}

func (x *APIClient) ExchangeToken(authCode string) (r *oauth2.Token, err error) {
	// Todo: use authCode to exchange tokens
	r, err = x.OAuth2Config.Exchange(context.Background(), authCode)
	if u.LogError(err) {
		return
	}

	err = x.TokenStore.SaveToken(r)
	if u.LogError(err) {
		return
	}

	return
}

// #endregion

// #region Reports

func (x *APIClient) RequestSponseredReports(query *SponseredReportsQuery) (r *ReportResponse, err error) {
	r = new(ReportResponse)

	if x.AdvURL == "" || x.ProfileID == "" || query.RequestURL == "" {
		err = errors.New("Missing URL or ProfileID")
		return
	}

	ts, err := x.getTokenSource()
	if err != nil {
		return
	}

	// body
	values := map[string]string{
		"reportDate": query.ReportDate,
		"metrics":    query.Metrics,
	}
	body, _ := json.Marshal(values)
	url := x.AdvURL + query.RequestURL

	client := oauth2.NewClient(oauth2.NoContext, ts)
	resp, err := x.request("POST", client, url, body)
	if resp.StatusCode != 202 || u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if u.LogError(err) {
		return
	}

	err = json.Unmarshal(bytes, &r)
	if u.LogError(err) {
		return
	}

	return
}

func (x *APIClient) GetReport(reportID string) (r []byte, err error) {
	r = make([]byte, 0)

	if x.AdvURL == "" || x.ProfileID == "" {
		err = errors.New("Missing AdvURL or ProfileID")
		return
	}

	// token
	ts, err := x.getTokenSource()
	if err != nil {
		return
	}

	// client
	url := x.AdvURL + "/v2/reports/" + reportID + "/download"
	client := oauth2.NewClient(oauth2.NoContext, ts)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// send request
	resp, err := x.request("GET", client, url, nil)
	if resp.StatusCode != 307 || u.LogError(err) {
		return
	}

	// get zip file
	dwonloadURL := resp.Header.Get("Location")
	resp, err = http.Get(dwonloadURL)
	if resp.StatusCode != 200 || u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	// unzip
	if resp.Header.Get("Content-Length") != "" {
		reader, err := gzip.NewReader(resp.Body)
		if u.LogError(err) {
			return r, err
		}
		defer reader.Close()

		bytes, err := ioutil.ReadAll(reader)
		if u.LogError(err) {
			return r, err
		}

		r = bytes
	}

	return
}

// #endregion

// #region Utilities

func (x *APIClient) request(action string, client *http.Client, url string, body []byte) (r *http.Response, err error) {
	r = new(http.Response)

	// init request
	req, err := http.NewRequest(action, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// headers
	contentType := "application/json"
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Amazon-Advertising-API-Scope", x.ProfileID)
	req.Header.Set("Amazon-Advertising-API-ClientId", x.OAuth2Config.ClientID)
	if strings.ToUpper(action) == "GET" {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	// send request
	r, err = client.Do(req)
	return
}

// #endregion
