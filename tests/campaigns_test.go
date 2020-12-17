package tests

import (
	"testing"

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
	id := "99020166512274"

	r, err := _campaignsClient.GetSPAdGroupsBid(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r)
}
