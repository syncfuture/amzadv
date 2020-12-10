package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/syncfuture/amzadv/reports"
)

func TestSponseredProducts(t *testing.T) {
	q := new(reports.SponseredReportsQuery)
	q.ReportDate = "20201026"
	q.Metrics = "sku,cost"
	q.RequestURL = "/v2/sp/productAds/report"

	// request report
	r, err := _reportsClient.RequestSponseredReports(q)
	assert.NoError(t, err)
	assert.NotNil(t, r.ReportID)
	t.Log(r.ReportID)

	if r.ReportID != "" {
		time.Sleep(3 * time.Second)

		// get report
		rs, err := _reportsClient.GetReport(r.ReportID)
		assert.NoError(t, err)
		assert.NotNil(t, rs)

		// decode
		v := make([]*reports.SponseredProductsReportDTO, 0)
		_ = json.Unmarshal(rs, &v)
		t.Log(len(v))
	}
}

func TestSponseredBrand(t *testing.T) {
	q := new(reports.SponseredReportsQuery)
	q.ReportDate = "20201001"
	q.Metrics = "cost"
	q.RequestURL = "/v2/hsa/campaigns/report"

	// request report
	r, err := _reportsClient.RequestSponseredReports(q)
	assert.NoError(t, err)
	assert.NotNil(t, r.ReportID)
	t.Log(r.ReportID)

	if r.ReportID != "" {
		time.Sleep(2 * time.Second)

		// get report
		rs, err := _reportsClient.GetReport(r.ReportID)
		assert.NoError(t, err)
		assert.NotNil(t, rs)

		// decode
		v := make([]*reports.SponseredBrandReportDTO, 0)
		_ = json.Unmarshal(rs, &v)
		t.Logf("%+v", v)
	}
}

func TestGetProfiles(t *testing.T) {
	r, err := _reportsClient.GetProfiles()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Log(r)
}
