package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPortfolios(t *testing.T) {
	r, err := _apiClient.GetPortfolios()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r[0])
}

func TestGetSPCampaigns(t *testing.T) {
	id := "42441275323797"

	r, err := _apiClient.GetSPCampaigns(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("%v\n", r[0])
}
