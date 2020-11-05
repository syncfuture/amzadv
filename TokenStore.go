package amzadv

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/syncfuture/go/sredis"

	"github.com/syncfuture/go/u"
	"golang.org/x/oauth2"
)

const (
	_key = "amzadv:CONFIGS"
)

type ITokenStore interface {
	GetToken(args ...interface{}) (*oauth2.Token, error)
	SaveToken(token *oauth2.Token, args ...interface{}) error
}

type TokenStore struct {
	Client     redis.Cmdable
	MerchantID string
}

func NewTokenStore(merchantId string, config *sredis.RedisConfig) ITokenStore {
	r := new(TokenStore)
	r.MerchantID = merchantId
	r.Client = sredis.NewClient(config)
	return r
}

func (x *TokenStore) GetToken(args ...interface{}) (r *oauth2.Token, err error) {
	r = new(oauth2.Token)

	jsonStr, err := x.Client.HGet(_key, x.MerchantID).Result()
	if u.LogError(err) {
		return
	}

	dto := new(TokenStoreDTO)
	err = json.Unmarshal([]byte(jsonStr), &dto)
	if u.LogError(err) {
		return
	}

	t, err := time.Parse(time.RFC3339, dto.Expiry)
	u.LogError(err)

	r.AccessToken = dto.AccessToken
	r.RefreshToken = dto.RefreshToken
	r.Expiry = t
	return
}

func (x *TokenStore) SaveToken(token *oauth2.Token, args ...interface{}) (err error) {
	in := &TokenStoreDTO{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry.Format(time.RFC3339),
	}

	jsonBytes, err := json.Marshal(in)
	if u.LogError(err) {
		return err
	}

	x.Client.HSet(_key, x.MerchantID, string(jsonBytes))
	return nil
}
