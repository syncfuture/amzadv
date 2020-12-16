package campaigns

import (
	"encoding/json"
	"errors"

	"github.com/syncfuture/amzadv/core"
	"github.com/syncfuture/amzadv/protoc/campaignsmodel"
	"golang.org/x/oauth2"

	"github.com/syncfuture/go/u"
)

type CampaignsAPI struct {
	core.APIClient
}

func NewCampaignsAPI(config *oauth2.Config, tokenStore core.ITokenStore) (r *CampaignsAPI) {
	r = new(CampaignsAPI)
	r.OAuth2Config = config
	r.TokenStore = tokenStore
	return r
}

// #region Get

func (x *CampaignsAPI) GetPortfolios() (r []*core.PortfolioDTO, err error) {
	r = make([]*core.PortfolioDTO, 0)

	// Send request
	bytes, err := x.Send("GET", "/v2/portfolios?portfolioStateFilter=enabled", nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSPCampaigns(portfolioID string) (r []*campaignsmodel.SPCampaignDTO, err error) {
	r = make([]*campaignsmodel.SPCampaignDTO, 0)

	if portfolioID == "" {
		err = errors.New("Missing PortfolioID")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/campaigns?&portfolioIdFilter="+portfolioID, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSBCampaigns(portfolioID string) (r []*campaignsmodel.SBCampaignDTO, err error) {
	r = make([]*campaignsmodel.SBCampaignDTO, 0)

	if portfolioID == "" {
		err = errors.New("Missing PortfolioID")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/sb/campaigns?&portfolioIdFilter="+portfolioID, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Ad Groups

func (x *CampaignsAPI) GetSPAdGroups(campaignId string) (r *campaignsmodel.AdGroupDTO, err error) {
	r = new(campaignsmodel.AdGroupDTO)

	if campaignId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/adGroups?&campaignIdFilter="+campaignId, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSPAdGroupsBid(adGroupId string) (r *campaignsmodel.BidRecommendationsDTO, err error) {
	r = new(campaignsmodel.BidRecommendationsDTO)

	if adGroupId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/adGroups/"+adGroupId+"/bidRecommendations", nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Targets

func (x *CampaignsAPI) GetSPTargets(adGroupId string) (r []*campaignsmodel.TargetDTO, err error) {
	r = make([]*campaignsmodel.TargetDTO, 0)

	if adGroupId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/targets?&adGroupIdFilter="+adGroupId, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSPTargetsBids(target *campaignsmodel.TargetDTO) (r []*campaignsmodel.BidRecommendationsDTO, err error) {
	r = make([]*campaignsmodel.BidRecommendationsDTO, 0)

	if target.AdGroupID == 0 {
		err = errors.New("Missing AdGroupId")
		return
	}

	// Send request
	body, _ := json.Marshal(target)
	bytes, err := x.Send("POST", "/v2/sp/targets/bidRecommendations", body)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Create & Update

func (x *CampaignsAPI) CreateSPCampaigns(entries []*campaignsmodel.SPCampaignDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.Send("POST", "/sb/campaigns", body)
}

func (x *CampaignsAPI) CreateSPAdGroups(entries []*campaignsmodel.AdGroupDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.Send("POST", "/sb/adGroups", body)
}

func (x *CampaignsAPI) CreateSPKeywords(entries []*campaignsmodel.KeywordDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.Send("POST", "/sb/keywords", body)
}

func (x *CampaignsAPI) CreateSPProductAds(entries []*campaignsmodel.ProductAdDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.Send("POST", "/sb/productAds", body)
}

func (x *CampaignsAPI) UpdateSPTargets(entries []*campaignsmodel.TargetDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.Send("PUT", "/sb/targets/", body)
}

// #endregion
