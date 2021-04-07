package campaignsmodel

import (
	"encoding/json"

	"github.com/syncfuture/go/sconv"
	"github.com/syncfuture/go/serr"
)

// #region SPCampaignDTO

type AmazonSPCampaignDTO struct {
	SPCampaignDTO
}

func (x *SPCampaignDTO) ToAmazonSPCampaignDTO() *AmazonSPCampaignDTO {
	return &AmazonSPCampaignDTO{SPCampaignDTO: (*x)}
}

func (t *AmazonSPCampaignDTO) MarshalJSON() ([]byte, error) {
	type ourDTOType AmazonSPCampaignDTO // create new type with same structure as T but without its method set!
	type amazonDTOType struct {
		ourDTOType        // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}

	amazonDTO := &amazonDTOType{
		ourDTOType: ourDTOType(*t),
	}

	amazonDTO.ID = sconv.ToInt64(t.ID)
	amazonDTO.PortfolioID = sconv.ToInt64(t.PortfolioID)

	return json.Marshal(amazonDTO)
}

func (x *AmazonSPCampaignDTO) ToSPCampaignDTO() *SPCampaignDTO {
	return &x.SPCampaignDTO
}

func (t *AmazonSPCampaignDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonSPCampaignDTO // create new type with same structure as T but without its method set!
	x := struct {
		T2                // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &x)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonSPCampaignDTO(x.T2)
	t.ID = sconv.ToString(x.ID)
	t.PortfolioID = sconv.ToString(x.PortfolioID)
	return nil
}

// #endregion

// #region SBCampaignDTO

type AmazonSBCampaignDTO struct {
	SBCampaignDTO
}

func (x *SBCampaignDTO) ToAmazonSBCampaignDTO() *AmazonSBCampaignDTO {
	return &AmazonSBCampaignDTO{SBCampaignDTO: (*x)}
}

func (t *AmazonSBCampaignDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonSBCampaignDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2                // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.ID = sconv.ToInt64(t.ID)
	rs.PortfolioID = sconv.ToInt64(t.PortfolioID)
	return json.Marshal(rs)
}

func (x *AmazonSBCampaignDTO) ToSBCampaignDTO() *SBCampaignDTO {
	return &x.SBCampaignDTO
}

func (t *AmazonSBCampaignDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonSBCampaignDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2                // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonSBCampaignDTO(newT.T2)
	t.ID = sconv.ToString(newT.ID)
	t.PortfolioID = sconv.ToString(newT.PortfolioID)
	return nil
}

// #endregion

// #region BidRecommendationsDTO

type AmazonBidRecommendationsDTO struct {
	BidRecommendationsDTO
}

func (x *BidRecommendationsDTO) ToAmazonBidRecommendationsDTO() *AmazonBidRecommendationsDTO {
	return &AmazonBidRecommendationsDTO{BidRecommendationsDTO: (*x)}
}

func (t *AmazonBidRecommendationsDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonBidRecommendationsDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2              // embed
		AdGroupID int64 `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	return json.Marshal(rs)
}

func (x *AmazonBidRecommendationsDTO) ToBidRecommendationsDTO() *BidRecommendationsDTO {
	return &x.BidRecommendationsDTO
}

func (t *AmazonBidRecommendationsDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonBidRecommendationsDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2              // embed
		AdGroupID int64 `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonBidRecommendationsDTO(newT.T2)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	return nil
}

// #endregion

// #region AdGroupDTO

type AmazonAdGroupDTO struct {
	AdGroupDTO
}

func (x *AdGroupDTO) ToAmazonAdGroupDTO() *AmazonAdGroupDTO {
	return &AmazonAdGroupDTO{AdGroupDTO: (*x)}
}

func (t *AmazonAdGroupDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonAdGroupDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2               // embed
		ID         int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"adGroupId,omitempty"`
		CampaignID int64 `protobuf:"varint,2,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.ID = sconv.ToInt64(t.ID)
	rs.CampaignID = sconv.ToInt64(t.CampaignID)
	return json.Marshal(rs)
}

func (x *AmazonAdGroupDTO) ToAdGroupDTO() *AdGroupDTO {
	return &x.AdGroupDTO
}

func (t *AmazonAdGroupDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonAdGroupDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		ID         int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"adGroupId,omitempty"`
		CampaignID int64 `protobuf:"varint,2,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonAdGroupDTO(newT.T2)
	t.ID = sconv.ToString(newT.ID)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	return nil
}

// #endregion

// #region ProductAdDTO

type AmazonProductAdDTO struct {
	ProductAdDTO
}

func (x *ProductAdDTO) ToAmazonProductAdDTO() *AmazonProductAdDTO {
	return &AmazonProductAdDTO{ProductAdDTO: (*x)}
}

func (t *AmazonProductAdDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonProductAdDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2               // embed
		AdID       int64 `protobuf:"varint,1,opt,name=AdID,proto3" json:"adId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		CampaignID int64 `protobuf:"varint,3,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.AdID = sconv.ToInt64(t.AdID)
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	rs.CampaignID = sconv.ToInt64(t.CampaignID)
	return json.Marshal(rs)
}

func (x *AmazonProductAdDTO) ToProductAdDTO() *ProductAdDTO {
	return &x.ProductAdDTO
}

func (t *AmazonProductAdDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonProductAdDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		AdID       int64 `protobuf:"varint,1,opt,name=AdID,proto3" json:"adId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		CampaignID int64 `protobuf:"varint,3,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonProductAdDTO(newT.T2)
	t.AdID = sconv.ToString(newT.AdID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	return nil
}

// #endregion

// #region KeywordDTO

type AmazonKeywordDTO struct {
	KeywordDTO
}

func (x *KeywordDTO) ToAmazonKeywordDTO() *AmazonKeywordDTO {
	return &AmazonKeywordDTO{KeywordDTO: (*x)}
}

func (t *AmazonKeywordDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonKeywordDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		KeywordID  int64 `protobuf:"varint,3,opt,name=KeywordID,proto3" json:"keywordId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.CampaignID = sconv.ToInt64(t.CampaignID)
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	rs.KeywordID = sconv.ToInt64(t.KeywordID)
	return json.Marshal(rs)
}

func (x *AmazonKeywordDTO) ToKeywordDTO() *KeywordDTO {
	return &x.KeywordDTO
}

func (t *AmazonKeywordDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonKeywordDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		KeywordID  int64 `protobuf:"varint,3,opt,name=KeywordID,proto3" json:"keywordId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonKeywordDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.KeywordID = sconv.ToString(newT.KeywordID)
	return nil
}

// #endregion

// #region TargetDTO

type AmazonTargetDTO struct {
	TargetDTO
}

func (x *TargetDTO) ToAmazonTargetDTO() *AmazonTargetDTO {
	return &AmazonTargetDTO{TargetDTO: (*x)}
}

func (t *AmazonTargetDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonTargetDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		TargetID   int64 `protobuf:"varint,3,opt,name=TargetID,proto3" json:"targetId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.CampaignID = sconv.ToInt64(t.CampaignID)
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	rs.TargetID = sconv.ToInt64(t.TargetID)
	return json.Marshal(rs)
}

func (x *AmazonTargetDTO) ToTargetDTO() *TargetDTO {
	return &x.TargetDTO
}

func (t *AmazonTargetDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonTargetDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
		TargetID   int64 `protobuf:"varint,3,opt,name=TargetID,proto3" json:"targetId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonTargetDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.TargetID = sconv.ToString(newT.TargetID)
	return nil
}

// #endregion

// #region ResponseDTO

type AmazonResponseDTO struct {
	ResponseDTO
}

func (x *ResponseDTO) ToAmazonTargetDTO() *AmazonResponseDTO {
	return &AmazonResponseDTO{ResponseDTO: (*x)}
}

func (t *AmazonResponseDTO) MarshalJSON() ([]byte, error) {
	type T2 AmazonResponseDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.CampaignID = sconv.ToInt64(t.CampaignID)
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	return json.Marshal(rs)
}

func (x *AmazonResponseDTO) ToResponseDTO() *ResponseDTO {
	return &x.ResponseDTO
}

func (t *AmazonResponseDTO) UnmarshalJSON(d []byte) error {
	type T2 AmazonResponseDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AmazonResponseDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	return nil
}

// #endregion
