package core

type TokenStoreDTO struct {
	// TokenType    string
	AccessToken  string
	RefreshToken string
	Expiry       string
}

type ProfileDTO struct {
	ProfileID   int64          `json:"profileId"`
	CountryCode string         `json:"countryCode"`
	AccountInfo AccountInfoDTO `json:"accountInfo"`
}

type AccountInfoDTO struct {
	ID            string `json:"id"`
	MarketplaceID string `json:"marketplaceStringId"`
	Type          string `json:"type"`
}

type PortfolioDTO struct {
	ID       int    `json:"portfolioId"`
	Name     string `json:"name"`
	InBudget bool   `json:"inBudget"`
	State    string `json:"state"`
}
