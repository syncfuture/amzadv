package amzadv

type SponseredReportsQuery struct {
	RequestURL string
	ProfileID  string
	ReportDate string
	Metrics    string
}

type SponseredProductsReportDTO struct {
	SKU  string  `json:"sku"`
	COST float32 `json:"cost"`
}

type SponseredBrandReportDTO struct {
	COST float32 `json:"cost"`
}

type ReportResponse struct {
	ReportID string `json:"reportId"`
}
