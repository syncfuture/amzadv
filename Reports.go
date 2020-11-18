package amzadv

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

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
