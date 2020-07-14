package amzadv

type TokenStoreDTO struct {
	AccessToken  string
	RefreshToken string
}

type SponseredReportsQuery struct {
	// MerchantID string
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
