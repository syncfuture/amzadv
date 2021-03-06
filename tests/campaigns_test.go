package tests

import (
	"encoding/json"
	"testing"

	"github.com/syncfuture/amzadv/protoc/campaignsmodel"
	"github.com/syncfuture/go/u"

	"github.com/syncfuture/amzadv/campaigns"

	"github.com/stretchr/testify/assert"
)

func TestGetPortfolios(t *testing.T) {
	r, err := _campaignsClient.GetPortfolios()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r[0])
}

func TestGetSPCampaigns(t *testing.T) {
	// id := "111"

	r, err := _campaignsClient.GetSPCampaigns("")
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%d\n", len(r))
	t.Logf("%v\n", r[0])
}

func TestCreateSPCampaigns(t *testing.T) {
	in := []*campaignsmodel.SPCampaignDTO{
		&campaignsmodel.SPCampaignDTO{
			// ID:                   42441275323797,
			// PortfolioID:          254949460323282,
			Name:                 "test-campaign - Auto",
			CampaignType:         "sponsoredProducts",
			TargetingType:        "auto",
			PremiumBidAdjustment: false,
			State:                "enabled",
			DailyBudget:          1.00,
			// StartDate:            "20210131",
			// EndDate:              "",
			Bidding: &campaignsmodel.BiddingDTO{
				Strategy: "legacyForSales",
				Adjustments: []*campaignsmodel.AdjustmentDTO{
					&campaignsmodel.AdjustmentDTO{
						Predicate:  "placementTop",
						Percentage: 50,
					},
					&campaignsmodel.AdjustmentDTO{
						Predicate:  "placementProductPage",
						Percentage: 50,
					},
				},
			},
		},
	}

	r, err := _campaignsClient.CreateSPCampaigns(in)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r)
}

func TestGetAdGroupBid(t *testing.T) {
	in := &campaigns.TargetsBidsQuery{
		AdGroupID: 126059537267120,
		Expressions: [][]*campaigns.Expression{
			{&campaigns.Expression{Type: "queryBroadRelMatches"}},
			{&campaigns.Expression{Type: "queryHighRelMatches"}},
			{&campaigns.Expression{Type: "asinSubstituteRelated"}},
			{&campaigns.Expression{Type: "asinAccessoryRelated"}},
		},
	}

	r, err := _campaignsClient.GetSPTargetsBids(in)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r)
}

func TestGetBrands(t *testing.T) {
	r, err := _campaignsClient.GetStores()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Log(r)
}

func TestUploadImages(t *testing.T) {
	in := &campaignsmodel.CreateStoreAssetCommand{
		ImageName: "12345",
		ImageType: "image/jpeg",
		AssetInfo: &campaignsmodel.StoreAssetDTO{
			BrandEntityId: "1234",
			MediaType:     "brandLogo",
		},
		Asset: make([]byte, 0),
	}

	r, err := _campaignsClient.CreateStoreAssets(in)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Log(r)
}

func TestSerialize(t *testing.T) {
	a := new(campaignsmodel.SPCampaignDTO)
	a.ID = "144371381093537413"
	a.PortfolioID = "144371381093555555"

	data, err := json.Marshal(a)
	assert.NoError(t, err)
	t.Log(u.BytesToStr(data))

	var b *campaignsmodel.SPCampaignDTO
	err = json.Unmarshal(data, &b)
	assert.NoError(t, err)
}
