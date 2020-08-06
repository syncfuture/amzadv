package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/syncfuture/amzadv"
)

func TestSponseredProducts(t *testing.T) {
	q := new(amzadv.SponseredReportsQuery)
	q.ReportDate = "20200630"
	q.Metrics = "sku,cost"
	q.RequestURL = "/v2/sp/productAds/report"

	// request report
	r, err := _apiClient.RequestSponseredReports(q)
	assert.NoError(t, err)
	assert.NotNil(t, r.ReportID)
	t.Log(r.ReportID)

	if r.ReportID != "" {
		time.Sleep(3 * time.Second)

		// get report
		rs, err := _apiClient.GetReport(r.ReportID)
		assert.NoError(t, err)
		assert.NotNil(t, rs)

		// decode
		v := make([]*amzadv.SponseredProductsReportDTO, 0)
		_ = json.Unmarshal(rs, &v)
		t.Log(v)
	}
}

func TestSponseredBrand(t *testing.T) {
	q := new(amzadv.SponseredReportsQuery)
	q.ReportDate = "20200630"
	q.Metrics = "cost"
	q.RequestURL = "/v2/hsa/campaigns/report"

	// request report
	r, err := _apiClient.RequestSponseredReports(q)
	assert.NoError(t, err)
	assert.NotNil(t, r.ReportID)
	t.Log(r.ReportID)

	if r.ReportID != "" {
		time.Sleep(2 * time.Second)

		// get report
		rs, err := _apiClient.GetReport(r.ReportID)
		assert.NoError(t, err)
		assert.NotNil(t, rs)

		// decode
		v := make([]*amzadv.SponseredBrandReportDTO, 0)
		_ = json.Unmarshal(rs, &v)
		t.Logf("%+v", v)
	}
}

func TestGetProfiles(t *testing.T) {
	r, err := _apiClient.GetProfiles()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Log(r)
}
