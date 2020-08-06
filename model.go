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
	ReportID   string `json:"reportId"`
	RecordType string `json:"recordType"`
	Status     string `json:"status"`
}

type ProfileDTO struct {
	ProfileID   int64          `json:"profileId"`
	CountryCode string         `json:"countryCode"`
	AccountInfo AccountInfoDTO `json:"accountInfo"`
}

type AccountInfoDTO struct {
	ID string `json:"id"`
}
