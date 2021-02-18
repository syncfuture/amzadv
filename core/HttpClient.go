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

func (x *APIClient) NewHttpClient() (r *http.Client, err error) {
	r = new(http.Client)

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
	r = oauth2.NewClient(oauth2.NoContext, ts)
	r.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return
}

func (x *APIClient) NewHttpRequest(action string, url string, body []byte) (r *http.Request) {
	r = new(http.Request)

	// init request
	r, err := http.NewRequest(action, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// headers
	contentType := "application/json"
	r.Header.Set("Content-Type", contentType)
	r.Header.Set("Amazon-Advertising-API-ClientId", x.OAuth2Config.ClientID)
	r.Header.Set("Amazon-Advertising-API-Scope", x.ProfileID)
	if strings.ToUpper(action) == "GET" {
		r.Header.Set("Accept-Encoding", "gzip")
	}

	return
}

func (x *APIClient) HttpSend(req *http.Request) (r []byte, err error) {
	r = make([]byte, 0)

	// client
	client, err := x.NewHttpClient()
	if err != nil {
		return
	}

	// send request
	resp, err := client.Do(req)
	if u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	r, err = ioutil.ReadAll(resp.Body)
	u.LogError(err)
	return
}
