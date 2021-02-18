package tests

import (
	"github.com/syncfuture/amzadv/campaigns"
	"github.com/syncfuture/amzadv/core"
	"github.com/syncfuture/amzadv/reports"
	config "github.com/syncfuture/go/sconfig"
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
	_reportsClient   *reports.ReportsAPI
	_campaignsClient *campaigns.CampaignsAPI
	_profiles        = map[string]string{
		"lf": "2022686179132559",
		"ff": "2418655185698010 ",
		"hh": "2496076506396355 ",
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
	tokenStore := core.NewTokenStore("lf", redisConfig)

	_reportsClient = reports.NewReportsAPI(c, tokenStore)
	_reportsClient.AdvURL = AdvURL
	_reportsClient.ProfileID = _profiles["lf"]

	_campaignsClient = campaigns.NewCampaignsAPI(c, tokenStore)
	_campaignsClient.AdvURL = AdvURL
	_campaignsClient.ProfileID = _profiles["lf"]
}
