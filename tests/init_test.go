package tests

import (
	"github.com/syncfuture/amzadv"
	"github.com/syncfuture/go/soauth2"
	"golang.org/x/oauth2"
)

const (
	TokenURL     = "https://api.amazon.com/auth/o2/token"
	RedirectURI  = "https://da.armos.ai/connect/amazon"
	ClientID     = "amzn1.application-oa2-client.cb6856e698c74eb480ce7804f428b3dc"
	ClientSecret = "58ba5905c68a3b00d2cf3b100a5c2ead7c9ade45620e0a74d188f772de69541c"
	AccessToken  = "Atza|IwEBIG2tb40VTl7fdoEjTU8xR-d8HuAtDUhfm3Pr5VGdT9R4ILHdqJU8pDZye-1PRw0EW_8kIuvDVmUYwGUjK6R2lkWsriPIKj5-UxiiV5ZnzPjtOtEa1v36Ql_XHg_8bM4_EP6FJ-Ow3dCgpUOTmGBraZke6KRcCx468KByN_QgJzZ6I9F_PchvipF1GfDeIx1mkiDWBsn1syJYt-LArvWGjC5F_D4nN9_u7IjDMG-4v6UxawCTHc6lv8eBfuCDRWEpf1qDL-dPFsngznraoMrEe224MxpA_Bhx9YsRKicZEh4PF8scr_r8HWCnHh9EN8A1gLwq_S_hpj-T5GRUC-L0oWediPZvRPqQFL6J1nVz3YJ29u10nDeBrMKvQE6aoJu-aK2uGvVhx_Vgu4YxG2SSNPAoJcTEwLInDVpmNpUHDnJvtJpaX81_8nZsaAbGIREsHrgFLUWSVHzZVQ6QjBFRjMvu"
	RefreshToken = "Atzr|IwEBIFyTj-1NzZcQLtJ1JDr-UDZ13NGNE1L7xR3FBgcNReJvRg0_pblocR9gPbp9ZSR5SIZD1QzYA6feUR1lZuOjQWV0h6A1Zme6DdIa-ZKnzXTNrXbAmEv429a13u9k8OLo5dq720BTssgtyPD08hQgdhgnQLMrFWRwQ_sXmy8b7Xm5EbGnNm7jptpE57rlbqoQhT50Ta_BQKmNvnJyS-oHOAT5cIzUfu5-2RSY6aKTDK_nVg-JNEjpAbtQJmJqFjBbfEa1kJvDf5M56gKR6GdcUxxMq7v2DMYCb0eAKU_kVH0R7ti89OCoC4-gQ7qajJLxugpKX1wf0Es3xomfOIJUTs_lyynwy2gFQlbODOtPr1fWBZQKM__S6tBLnj2dKqnecE4a_fBw13oEbCixOCPE9_khNwmEUNq7JaLZpYf_Xp_ySzugtopliy12ozIURS8oCvw"
)

var (
	_apiClient *amzadv.APIClient
	Profiles   = map[string]string{
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
	ts := NewTokenStore()

	_apiClient = amzadv.NewAPIClient(c, ts)
}

func NewTokenStore() soauth2.ITokenStore {
	r := new(TestTokenStore)
	return r
}

func (x *TestTokenStore) GetToken(args ...interface{}) (r *oauth2.Token, err error) {
	r = new(oauth2.Token)
	r.AccessToken = AccessToken
	r.RefreshToken = RefreshToken

	return
}

func (x *TestTokenStore) SaveToken(token *oauth2.Token, args ...interface{}) (err error) {
	return
}
