package core

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

func (x *APIClient) Send(action string, uri string, body []byte) (r []byte, err error) {
	r = make([]byte, 0)

	if x.AdvURL == "" || x.ProfileID == "" {
		err = errors.New("Missing AdvURL or ProfileID")
		return
	}

	// token
	ts, err := x.GetTokenSource()
	if err != nil {
		return
	}

	// client
	url := x.AdvURL + uri
	client := oauth2.NewClient(oauth2.NoContext, ts)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// send request
	resp, err := x.Request(action, client, url, body)
	if u.LogError(err) {
		// resp.StatusCode != 200 ||
		return
	}
	defer resp.Body.Close()

	r, err = ioutil.ReadAll(resp.Body)
	u.LogError(err)
	return
}

func (x *APIClient) Request(action string, client *http.Client, url string, body []byte) (r *http.Response, err error) {
	r = new(http.Response)

	// init request
	req, err := http.NewRequest(action, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// headers
	contentType := "application/json"
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Amazon-Advertising-API-ClientId", x.OAuth2Config.ClientID)
	req.Header.Set("Amazon-Advertising-API-Scope", x.ProfileID)
	if strings.ToUpper(action) == "GET" {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	// send request
	r, err = client.Do(req)
	return
}
