package tests

import (
	"testing"

	"github.com/syncfuture/amzadv/protoc/campaignsmodel"

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
	id := "111"

	r, err := _campaignsClient.GetSBCampaigns(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%d\n", len(r))
	t.Logf("%v\n", r[0])
}

func TestCreateSPCampaigns(t *testing.T) {
	in := []*campaignsmodel.SPCampaignDTO{
		&campaignsmodel.SPCampaignDTO{
			// ID:                   42441275323797,
			PortfolioID:          8956684801033,
			Name:                 "test-campaign - Auto",
			CampaignType:         "sponsoredProducts",
			TargetingType:        "auto",
			PremiumBidAdjustment: false,
			State:                "enabled",
			DailyBudget:          1.00,
			StartDate:            "20210131",
			EndDate:              "",
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
