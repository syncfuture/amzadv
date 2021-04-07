package campaignsmodel

import (
	"encoding/json"

	"github.com/syncfuture/go/sconv"
	"github.com/syncfuture/go/serr"
)

// #region SPCampaignDTO

func (t *SPCampaignDTO) MarshalJSON() ([]byte, error) {
	type ourDTOType SPCampaignDTO // create new type with same structure as T but without its method set!
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

func (t *SPCampaignDTO) UnmarshalJSON(d []byte) error {
	type T2 SPCampaignDTO // create new type with same structure as T but without its method set!
	x := struct {
		T2                // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &x)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = SPCampaignDTO(x.T2)
	t.ID = sconv.ToString(x.ID)
	t.PortfolioID = sconv.ToString(x.PortfolioID)
	return nil
}

// #endregion

// #region SBCampaignDTO

func (t *SBCampaignDTO) MarshalJSON() ([]byte, error) {
	type T2 SBCampaignDTO // create new type with same structure as T but without its method set!
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

func (t *SBCampaignDTO) UnmarshalJSON(d []byte) error {
	type T2 SBCampaignDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2                // embed
		ID          int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"campaignId,omitempty"`
		PortfolioID int64 `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"portfolioId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = SBCampaignDTO(newT.T2)
	t.ID = sconv.ToString(newT.ID)
	t.PortfolioID = sconv.ToString(newT.PortfolioID)
	return nil
}

// #endregion

// #region BidRecommendationsDTO

func (t *BidRecommendationsDTO) MarshalJSON() ([]byte, error) {
	type T2 BidRecommendationsDTO // create new type with same structure as T but without its method set!
	type newT struct {
		T2              // embed
		AdGroupID int64 `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}

	rs := &newT{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them
	rs.AdGroupID = sconv.ToInt64(t.AdGroupID)
	return json.Marshal(rs)
}

func (t *BidRecommendationsDTO) UnmarshalJSON(d []byte) error {
	type T2 BidRecommendationsDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2              // embed
		AdGroupID int64 `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = BidRecommendationsDTO(newT.T2)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	return nil
}

// #endregion

// #region AdGroupDTO

func (t *AdGroupDTO) MarshalJSON() ([]byte, error) {
	type T2 AdGroupDTO // create new type with same structure as T but without its method set!
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

func (t *AdGroupDTO) UnmarshalJSON(d []byte) error {
	type T2 AdGroupDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		ID         int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"adGroupId,omitempty"`
		CampaignID int64 `protobuf:"varint,2,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = AdGroupDTO(newT.T2)
	t.ID = sconv.ToString(newT.ID)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	return nil
}

// #endregion

// #region ProductAdDTO

func (t *ProductAdDTO) MarshalJSON() ([]byte, error) {
	type T2 ProductAdDTO // create new type with same structure as T but without its method set!
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

func (t *ProductAdDTO) UnmarshalJSON(d []byte) error {
	type T2 ProductAdDTO // create new type with same structure as T but without its method set!
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

	*t = ProductAdDTO(newT.T2)
	t.AdID = sconv.ToString(newT.AdID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	return nil
}

// #endregion

// #region KeywordDTO

func (t *KeywordDTO) MarshalJSON() ([]byte, error) {
	type T2 KeywordDTO // create new type with same structure as T but without its method set!
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

func (t *KeywordDTO) UnmarshalJSON(d []byte) error {
	type T2 KeywordDTO // create new type with same structure as T but without its method set!
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

	*t = KeywordDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.KeywordID = sconv.ToString(newT.KeywordID)
	return nil
}

// #endregion

// #region TargetDTO

func (t *TargetDTO) MarshalJSON() ([]byte, error) {
	type T2 TargetDTO // create new type with same structure as T but without its method set!
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

func (t *TargetDTO) UnmarshalJSON(d []byte) error {
	type T2 TargetDTO // create new type with same structure as T but without its method set!
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

	*t = TargetDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	t.TargetID = sconv.ToString(newT.TargetID)
	return nil
}

// #endregion

// #region ResponseDTO

func (t *ResponseDTO) MarshalJSON() ([]byte, error) {
	type T2 ResponseDTO // create new type with same structure as T but without its method set!
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

func (t *ResponseDTO) UnmarshalJSON(d []byte) error {
	type T2 ResponseDTO // create new type with same structure as T but without its method set!
	newT := struct {
		T2               // embed
		CampaignID int64 `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"campaignId,omitempty"`
		AdGroupID  int64 `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"adGroupId,omitempty"`
	}{T2: T2(*t)} // don't forget this, if you do and 't' already has some fields set you would lose them

	err := json.Unmarshal(d, &newT)
	if err != nil {
		return serr.WithStack(err)
	}

	*t = ResponseDTO(newT.T2)
	t.CampaignID = sconv.ToString(newT.CampaignID)
	t.AdGroupID = sconv.ToString(newT.AdGroupID)
	return nil
}

// #endregion
