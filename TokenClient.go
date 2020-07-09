package amzadv

import (
	"golang.org/x/oauth2"
)

type TokenClient struct {
	ClientID     string
	ClientSecret string
}

func NewTokenClient(clientID, clientSecret string) (r *TokenClient) {
	r = &TokenClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	return
}

func (x *TokenClient) ExchangeToken(authCode string) (r *oauth2.Token, err error) {
	// Todo: use authCode to exchange tokens
	return
}

func (x *TokenClient) RefreshToken(refreshToken string) (r *oauth2.Token, err error) {
	// Todo: use refresh token to get new tokens
	return
}
