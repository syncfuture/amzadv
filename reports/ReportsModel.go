package reports

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
	Details    string `json:"details"`
}
