package tests

import (
	"testing"

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
	id := "42441275323797"

	r, err := _campaignsClient.GetSPCampaigns(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%d\n", len(r))
	t.Logf("%v\n", r[0])
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
