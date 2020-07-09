package amzadv

import (
	"context"

	"github.com/syncfuture/go/soauth2"
	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

type APIClient struct {
	OAuth2Config *oauth2.Config
	TokenStore   soauth2.ITokenStore
}

func NewAPIClient(config *oauth2.Config, tokenStore soauth2.ITokenStore) (r *APIClient) {
	r = &APIClient{
		OAuth2Config: config,
		TokenStore:   tokenStore,
	}
	return r
}

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

func (x *APIClient) XXXXXXApi() (err error) {
	ts, err := x.getTokenSource()
	if err != nil {
		return err
	}

	client := oauth2.NewClient(oauth2.NoContext, ts)
	resp, err := client.Get("")
	if u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	return
}

// Todo: create more api as needed.
