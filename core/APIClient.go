package core

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

type APIClient struct {
	OAuth2Config *oauth2.Config
	TokenStore   ITokenStore
	AdvURL       string
	ProfileID    string
}

// func NewAPIClient(config *oauth2.Config, tokenStore ITokenStore) (r *APIClient) {
// 	r = &APIClient{
// 		OAuth2Config: config,
// 		TokenStore:   tokenStore,
// 	}
// 	return r
// }

// #region Token

func (x *APIClient) GetTokenSource() (oauth2.TokenSource, error) {
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

// #region Profiles

func (x *APIClient) GetProfiles() (r []*ProfileDTO, err error) {
	r = make([]*ProfileDTO, 0)

	ts, err := x.GetTokenSource()
	if err != nil {
		return
	}

	url := x.AdvURL + "/v2/profiles"
	client := oauth2.NewClient(oauth2.NoContext, ts)
	resp, err := x.Request("GET", client, url, nil)
	if resp.StatusCode != 200 || u.LogError(err) {
		return
	}
	defer func() { resp.Body.Close() }()

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

// #endregion
