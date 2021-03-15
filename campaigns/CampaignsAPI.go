package campaigns

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/syncfuture/amzadv/core"
	"github.com/syncfuture/amzadv/protoc/campaignsmodel"
	"golang.org/x/oauth2"

	"github.com/syncfuture/go/u"
)

// var (
// 	AdGroupMap = map[string]string{
// 		"LF-US": "126059537267120",
// 		"LF-CA": "190293692100357",
// 		"LF-MX": "125142947684660",
// 		"FF-US": "244523535626646",
// 		"FF-CA": "244523535626646",
// 		"FF-MX": "244523535626646",
// 		"HH-US": "164315009197103",
// 		"HH-CA": "164315009197103",
// 		"HH-MX": "164315009197103",
// 	}
// )

type CampaignsAPI struct {
	core.APIClient
}

func NewCampaignsAPI(config *oauth2.Config, tokenStore core.ITokenStore) (r *CampaignsAPI) {
	r = new(CampaignsAPI)
	r.OAuth2Config = config
	r.TokenStore = tokenStore
	return r
}

func (x *CampaignsAPI) GetPortfolios() (r []*core.PortfolioDTO, err error) {
	r = make([]*core.PortfolioDTO, 0)

	// Send request
	bytes, err := x.Send("GET", "/v2/portfolios?portfolioStateFilter=enabled", nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #region Campaigns

func (x *CampaignsAPI) GetSPCampaigns(portfolioID string) (r []*campaignsmodel.SPCampaignDTO, err error) {
	r = make([]*campaignsmodel.SPCampaignDTO, 0)

	if portfolioID == "" {
		err = errors.New("Missing PortfolioID")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/campaigns?&portfolioIdFilter="+portfolioID, nil)
	if err != nil || len(bytes) == 0 {
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
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSPCampaignsByName(portfolioID string, name string) (r []*campaignsmodel.SPCampaignDTO, err error) {
	r = make([]*campaignsmodel.SPCampaignDTO, 0)

	if name == "" {
		err = errors.New("Missing Name")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/campaigns?&portfolioIdFilter="+portfolioID+"&name="+url.QueryEscape(name), nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) CreateSPCampaigns(entries []*campaignsmodel.SPCampaignDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("POST", "/v2/sp/campaigns", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) CreateSBCampaigns(entries []*campaignsmodel.SBCampaignDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("POST", "/sb/campaigns", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Ad Groups

func (x *CampaignsAPI) GetSPAdGroupsByName(campaignId string, name string) (r []*campaignsmodel.AdGroupDTO, err error) {
	r = make([]*campaignsmodel.AdGroupDTO, 0)

	if name == "" {
		err = errors.New("Missing Name")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/v2/sp/adGroups?campaignIdFilter="+campaignId+"&name="+url.QueryEscape(name), nil)
	if err != nil || len(bytes) == 0 {
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
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) CreateSPAdGroups(entries []*campaignsmodel.AdGroupDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("POST", "/v2/sp/adGroups", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Product Ads

func (x *CampaignsAPI) GetSPProductAds(adGroupId string, state string) (r []*campaignsmodel.ProductAdDTO, err error) {
	r = make([]*campaignsmodel.ProductAdDTO, 0)

	if adGroupId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	uri := "/v2/sp/productAds?&adGroupIdFilter=" + adGroupId
	if state != "" {
		uri += "&stateFilter=" + state
	}

	// Send request
	bytes, err := x.Send("GET", uri, nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) CreateSPProductAds(entries []*campaignsmodel.ProductAdDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("POST", "/v2/sp/productAds", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) UpdateSPProductAds(entries []*campaignsmodel.ProductAdDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("PUT", "/v2/sp/productAds", body)
	if err != nil || len(bytes) == 0 {
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
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetSPTargetsBids(in *TargetsBidsQuery) (r *campaignsmodel.BidRecommendationsDTO, err error) {
	r = new(campaignsmodel.BidRecommendationsDTO)

	if in.AdGroupID == 0 {
		err = errors.New("Missing AdGroupId")
		return
	}

	// Send request
	body, _ := json.Marshal(in)
	bytes, err := x.Send("POST", "/v2/sp/targets/bidRecommendations", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) UpdateSPTargets(entries []*campaignsmodel.TargetDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("PUT", "/v2/sp/targets", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region Keywords

func (x *CampaignsAPI) CreateSPKeywords(entries []*campaignsmodel.KeywordDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("POST", "/v2/sp/keywords", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) UpdateSBKeywords(entries []*campaignsmodel.KeywordDTO) (r []*campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(entries)
	bytes, err := x.Send("PUT", "/sb/keywords", body)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion

// #region SB Brands

func (x *CampaignsAPI) GetBrands() (r []*campaignsmodel.BrandDTO, err error) {
	r = make([]*campaignsmodel.BrandDTO, 0)

	// Send request
	bytes, err := x.Send("GET", "/brands", nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetStores() (r []*campaignsmodel.StoreDTO, err error) {
	r = make([]*campaignsmodel.StoreDTO, 0)

	// Send request
	bytes, err := x.Send("GET", "/v2/stores", nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetStoreAssets(brandEntityId string) (r []*campaignsmodel.StoreAssetDTO, err error) {
	r = make([]*campaignsmodel.StoreAssetDTO, 0)

	if brandEntityId == "" {
		err = errors.New("Missing brandEntityId")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/stores/assets?&brandEntityId="+brandEntityId, nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) GetLandingPageASINs(pageUrl string) (r *campaignsmodel.LandingPageASINDTO, err error) {
	r = new(campaignsmodel.LandingPageASINDTO)

	if pageUrl == "" {
		err = errors.New("Missing url")
		return
	}

	// Send request
	bytes, err := x.Send("GET", "/pageAsins?pageUrl="+pageUrl, nil)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *CampaignsAPI) CreateStoreAssets(in *campaignsmodel.CreateStoreAssetCommand) (r *campaignsmodel.ResponseDTO, err error) {
	body, _ := json.Marshal(in)

	req := x.NewHttpRequest("PUT", "/stores/assets", body)
	req.Header.Set("Content-Disposition", in.ImageName)
	req.Header.Set("Content-Type", in.ImageType)

	bytes, err := x.HttpSend(req)
	if err != nil || len(bytes) == 0 {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

// #endregion
