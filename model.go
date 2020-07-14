package amzadv

type TokenStoreDTO struct {
	// TokenType    string
	AccessToken  string
	RefreshToken string
	Expiry       string
}

type SponseredReportsQuery struct {
	RequestURL string
	ReportDate string
	Metrics    string
}

type SponseredProductsReportDTO struct {
	SKU  string  `json:"sku"`
	Cost float32 `json:"cost"`
}

type SponseredBrandReportDTO struct {
	Cost float32 `json:"cost"`
}

type ReportResponse struct {
	ReportID string `json:"reportId"`
}
