package campaigns

// // type SponseredReportsQuery struct {
// // 	RequestURL string
// // 	ReportDate string
// // 	Metrics    string
// // }

// type PortfolioDTO struct {
// 	ID       int    `json:"portfolioId"`
// 	Name     string `json:"name"`
// 	InBudget bool   `json:"inBudget"`
// 	State    string `json:"state"`
// }

// type AdjustmentDTO struct {
// 	Predicate  string  `json:"predicate"`
// 	Percentage float32 `json:"percentage"`
// }

// type BiddingDTO struct {
// 	Strategy    string           `json:"strategy"`
// 	Adjustments []*AdjustmentDTO `json:"adjustments"`
// }

// type CreativeDTO struct {
// 	BrandName           string   `json:"brandName"`
// 	BrandLogoAssetID    string   `json:"brandLogoAssetID"`
// 	Headline            string   `json:"headline"`
// 	ShouldOptimizeAsins bool     `json:"shouldOptimizeAsins"`
// 	ASINs               []string `json:"asins"`
// 	BrandLogoURL        string   `json:"brandLogoUrl"`
// }

// type LandingPageDTO struct {
// 	PageType string `json:"pageType"`
// 	URL      string `json:"url"`
// }

// type ExpressionDTO struct {
// 	Value string `json:"value"`
// 	Type  string `json:"type"`
// }

// type SuggestedBidDTO struct {
// 	End       string `json:"rangeEnd"`
// 	Start     string `json:"rangeStart"`
// 	Suggested string `json:"suggested"`
// }

// type RecommendationsDTO struct {
// 	Expression   []*ExpressionDTO `json:"expression"`
// 	SuggestedBid SuggestedBidDTO  `json:"suggestedBid"`
// }

// type SPCampaignDTO struct {
// 	ID                   int        `json:"campaignId"`
// 	PortfolioID          int        `json:"portfolioId"`
// 	Name                 string     `json:"name"`
// 	CampaignType         string     `json:"campaignType"`
// 	TargetingType        string     `json:"targetingType"`
// 	PremiumBidAdjustment bool       `json:"premiumBidAdjustment"`
// 	DailyBudget          float32    `json:"dailyBudget"`
// 	StartDate            string     `json:"startDate"`
// 	EndDate              string     `json:"endDate"`
// 	State                string     `json:"state"`
// 	Bidding              BiddingDTO `json:"bidding"`
// }

// type SBCampaignDTO struct {
// 	ID              int            `json:"campaignId"`
// 	PortfolioID     int            `json:"portfolioId"`
// 	Name            string         `json:"name"`
// 	Budget          float32        `json:"budget"`
// 	BidOptimization bool           `json:"bidOptimization"`
// 	AdFormat        string         `json:"adFormat"`
// 	BudgetType      string         `json:"	budgetType"`
// 	StartDate       string         `json:"startDate"`
// 	EndDate         string         `json:"endDate"`
// 	State           string         `json:"state"`
// 	ServingStatus   string         `json:"servingStatus"`
// 	Creative        CreativeDTO    `json:"creative"`
// 	LandingPage     LandingPageDTO `json:"landingPage"`
// }

// type AdGroupDTO struct {
// 	CampaignID int     `json:"campaignId"`
// 	Name       string  `json:"name"`
// 	State      string  `json:"state"`
// 	DefaultBid float32 `json:"defaultBid"`
// }

// type ProductAdDTO struct {
// 	AdGroupID  int    `json:"adGroupId"`
// 	CampaignID int    `json:"campaignId"`
// 	SKU        string `json:"sku"`
// 	State      string `json:"state"`
// }

// type KeywordDTO struct {
// 	CampaignID  int     `json:"campaignId"`
// 	AdGroupID   int     `json:"adGroupId"`
// 	State       string  `json:"state"`
// 	KeywordText string  `json:"keywordText"`
// 	MatchType   string  `json:"matchType"`
// 	Bid         float32 `json:"bid"`
// }

// type TargetDTO struct {
// 	CampaignID     int              `json:"campaignId"`
// 	AdGroupID      int              `json:"adGroupId"`
// 	TargetID       int              `json:"targetId"`
// 	ExpressionType string           `json:"expressionType"`
// 	Expression     []*ExpressionDTO `json:"expression"`
// 	Bid            float32          `json:"bid"`
// 	State          string           `json:"state"`
// }

// type BidRecommendationsDTO struct {
// 	AdGroupID       int                   `json:"adGroupId"`
// 	SuggestedBid    SuggestedBidDTO       `json:"suggestedBid"`
// 	Recommendations []*RecommendationsDTO `json:"recommendations"`
// }
