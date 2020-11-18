package amzadv

import (
	"encoding/json"
	"errors"

	"github.com/syncfuture/go/u"
)

// #region Get

func (x *APIClient) GetPortfolios() (r []*PortfolioDTO, err error) {
	r = make([]*PortfolioDTO, 0)

	// send request
	bytes, err := x.send("GET", "/v2/portfolios?portfolioStateFilter=enabled", nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *APIClient) GetSPCampaigns(portfolioID string) (r []*SPCampaignDTO, err error) {
	r = make([]*SPCampaignDTO, 0)

	if portfolioID == "" {
		err = errors.New("Missing PortfolioID")
		return
	}

	// send request
	bytes, err := x.send("GET", "/v2/sp/campaigns?&portfolioIdFilter="+portfolioID, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *APIClient) GetSBCampaigns(portfolioID string) (r []*SBCampaignDTO, err error) {
	r = make([]*SBCampaignDTO, 0)

	if portfolioID == "" {
		err = errors.New("Missing PortfolioID")
		return
	}

	// send request
	bytes, err := x.send("GET", "/sb/campaigns?&portfolioIdFilter="+portfolioID, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *APIClient) GetSPTargets(adGroupId string) (r []*TargetDTO, err error) {
	r = make([]*TargetDTO, 0)

	if adGroupId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	// send request
	bytes, err := x.send("GET", "/v2/sp/targets?&adGroupIdFilter="+adGroupId, nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *APIClient) GetSPAdGroupsBid(adGroupId string) (r *BidRecommendationsDTO, err error) {
	r = new(BidRecommendationsDTO)

	if adGroupId == "" {
		err = errors.New("Missing AdGroupId")
		return
	}

	// send request
	bytes, err := x.send("GET", "/v2/sp/adGroups/"+adGroupId+"/bidRecommendations", nil)
	if err != nil {
		return
	}

	// decode
	err = json.Unmarshal(bytes, &r)
	u.LogError(err)
	return
}

func (x *APIClient) GetSPTargetsBids(target *TargetDTO) (r []*BidRecommendationsDTO, err error) {
	r = make([]*BidRecommendationsDTO, 0)

	if target.AdGroupID == 0 {
		err = errors.New("Missing AdGroupId")
		return
	}

	// send request
	body, _ := json.Marshal(target)
	bytes, err := x.send("POST", "/v2/sp/targets/bidRecommendations", body)
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

func (x *APIClient) CreateSPCampaigns(entries []*SPCampaignDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.send("POST", "/sb/campaigns", body)
}

func (x *APIClient) CreateSPAdGroups(entries []*AdGroupDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.send("POST", "/sb/adGroups", body)
}

func (x *APIClient) CreateSPKeywords(entries []*KeywordDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.send("POST", "/sb/keywords", body)
}

func (x *APIClient) CreateSPProductAds(entries []*ProductAdDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.send("POST", "/sb/productAds", body)
}

func (x *APIClient) UpdateSPTargets(entries []*TargetDTO) (r []byte, err error) {
	body, _ := json.Marshal(entries)
	return x.send("PUT", "/sb/targets/", body)
}

// #endregion
