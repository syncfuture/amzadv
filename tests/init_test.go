package tests

import (
	"github.com/syncfuture/amzadv"
	"github.com/syncfuture/go/config"
	"github.com/syncfuture/go/sredis"
	"golang.org/x/oauth2"
)

const (
	AdvURL       = "https://advertising-api.amazon.com"
	TokenURL     = "https://api.amazon.com/auth/o2/token"
	RedirectURI  = "https://da.armos.ai/connect/amazon"
	ClientID     = "amzn1.application-oa2-client.cb6856e698c74eb480ce7804f428b3dc"
	ClientSecret = "58ba5905c68a3b00d2cf3b100a5c2ead7c9ade45620e0a74d188f772de69541c"
)

var (
	_apiClient *amzadv.APIClient
	_profiles  = map[string]string{
		"lf": "2022686179132559",
		"ff": "2418655185698010 ",
	}
)

type TestTokenStore struct {
	AccessToken  string
	RefreshToken string
}

func init() {
	// config
	endpoint := oauth2.Endpoint{
		TokenURL: TokenURL,
	}
	c := &oauth2.Config{
		Endpoint:     endpoint,
		RedirectURL:  RedirectURI,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
	}

	// token store
	var redisConfig *sredis.RedisConfig
	configProvider := config.NewJsonConfigProvider()
	configProvider.GetStruct("Redis", &redisConfig)
	tokenStore := amzadv.NewTokenStore(redisConfig)

	_apiClient = amzadv.NewAPIClient(c, tokenStore)
	_apiClient.AdvURL = AdvURL
	_apiClient.ProfileID = _profiles["lf"]
}
