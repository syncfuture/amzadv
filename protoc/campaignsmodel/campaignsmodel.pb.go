// Code generated by protoc-gen-go. DO NOT EDIT.
// source: campaignsmodel/campaignsmodel.proto

package campaignsmodel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SPCampaignDTO struct {
	ID                   int64       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PortfolioID          int64       `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"PortfolioID,omitempty"`
	Name                 string      `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	CampaignType         string      `protobuf:"bytes,4,opt,name=CampaignType,proto3" json:"CampaignType,omitempty"`
	TargetingType        string      `protobuf:"bytes,5,opt,name=TargetingType,proto3" json:"TargetingType,omitempty"`
	PremiumBidAdjustment bool        `protobuf:"varint,6,opt,name=PremiumBidAdjustment,proto3" json:"PremiumBidAdjustment,omitempty"`
	DailyBudget          float32     `protobuf:"fixed32,7,opt,name=DailyBudget,proto3" json:"DailyBudget,omitempty"`
	StartDate            string      `protobuf:"bytes,8,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string      `protobuf:"bytes,9,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	State                string      `protobuf:"bytes,10,opt,name=State,proto3" json:"State,omitempty"`
	Bidding              *BiddingDTO `protobuf:"bytes,11,opt,name=Bidding,proto3" json:"Bidding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SPCampaignDTO) Reset()         { *m = SPCampaignDTO{} }
func (m *SPCampaignDTO) String() string { return proto.CompactTextString(m) }
func (*SPCampaignDTO) ProtoMessage()    {}
func (*SPCampaignDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{0}
}

func (m *SPCampaignDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPCampaignDTO.Unmarshal(m, b)
}
func (m *SPCampaignDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPCampaignDTO.Marshal(b, m, deterministic)
}
func (m *SPCampaignDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPCampaignDTO.Merge(m, src)
}
func (m *SPCampaignDTO) XXX_Size() int {
	return xxx_messageInfo_SPCampaignDTO.Size(m)
}
func (m *SPCampaignDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_SPCampaignDTO.DiscardUnknown(m)
}

var xxx_messageInfo_SPCampaignDTO proto.InternalMessageInfo

func (m *SPCampaignDTO) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SPCampaignDTO) GetPortfolioID() int64 {
	if m != nil {
		return m.PortfolioID
	}
	return 0
}

func (m *SPCampaignDTO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SPCampaignDTO) GetCampaignType() string {
	if m != nil {
		return m.CampaignType
	}
	return ""
}

func (m *SPCampaignDTO) GetTargetingType() string {
	if m != nil {
		return m.TargetingType
	}
	return ""
}

func (m *SPCampaignDTO) GetPremiumBidAdjustment() bool {
	if m != nil {
		return m.PremiumBidAdjustment
	}
	return false
}

func (m *SPCampaignDTO) GetDailyBudget() float32 {
	if m != nil {
		return m.DailyBudget
	}
	return 0
}

func (m *SPCampaignDTO) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *SPCampaignDTO) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *SPCampaignDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *SPCampaignDTO) GetBidding() *BiddingDTO {
	if m != nil {
		return m.Bidding
	}
	return nil
}

type BiddingDTO struct {
	Strategy             string           `protobuf:"bytes,1,opt,name=Strategy,proto3" json:"Strategy,omitempty"`
	Adjustments          []*AdjustmentDTO `protobuf:"bytes,2,rep,name=Adjustments,proto3" json:"Adjustments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BiddingDTO) Reset()         { *m = BiddingDTO{} }
func (m *BiddingDTO) String() string { return proto.CompactTextString(m) }
func (*BiddingDTO) ProtoMessage()    {}
func (*BiddingDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{1}
}

func (m *BiddingDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingDTO.Unmarshal(m, b)
}
func (m *BiddingDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingDTO.Marshal(b, m, deterministic)
}
func (m *BiddingDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingDTO.Merge(m, src)
}
func (m *BiddingDTO) XXX_Size() int {
	return xxx_messageInfo_BiddingDTO.Size(m)
}
func (m *BiddingDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingDTO.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingDTO proto.InternalMessageInfo

func (m *BiddingDTO) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *BiddingDTO) GetAdjustments() []*AdjustmentDTO {
	if m != nil {
		return m.Adjustments
	}
	return nil
}

type AdjustmentDTO struct {
	Predicate            string   `protobuf:"bytes,1,opt,name=Predicate,proto3" json:"Predicate,omitempty"`
	Percentage           float32  `protobuf:"fixed32,2,opt,name=Percentage,proto3" json:"Percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdjustmentDTO) Reset()         { *m = AdjustmentDTO{} }
func (m *AdjustmentDTO) String() string { return proto.CompactTextString(m) }
func (*AdjustmentDTO) ProtoMessage()    {}
func (*AdjustmentDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{2}
}

func (m *AdjustmentDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdjustmentDTO.Unmarshal(m, b)
}
func (m *AdjustmentDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdjustmentDTO.Marshal(b, m, deterministic)
}
func (m *AdjustmentDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdjustmentDTO.Merge(m, src)
}
func (m *AdjustmentDTO) XXX_Size() int {
	return xxx_messageInfo_AdjustmentDTO.Size(m)
}
func (m *AdjustmentDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_AdjustmentDTO.DiscardUnknown(m)
}

var xxx_messageInfo_AdjustmentDTO proto.InternalMessageInfo

func (m *AdjustmentDTO) GetPredicate() string {
	if m != nil {
		return m.Predicate
	}
	return ""
}

func (m *AdjustmentDTO) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type SBCampaignDTO struct {
	ID                   int64           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PortfolioID          int64           `protobuf:"varint,2,opt,name=PortfolioID,proto3" json:"PortfolioID,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Budget               float32         `protobuf:"fixed32,4,opt,name=Budget,proto3" json:"Budget,omitempty"`
	BidOptimization      bool            `protobuf:"varint,5,opt,name=BidOptimization,proto3" json:"BidOptimization,omitempty"`
	AdFormat             string          `protobuf:"bytes,6,opt,name=AdFormat,proto3" json:"AdFormat,omitempty"`
	BudgetType           string          `protobuf:"bytes,7,opt,name=BudgetType,proto3" json:"BudgetType,omitempty"`
	StartDate            string          `protobuf:"bytes,8,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string          `protobuf:"bytes,9,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	State                string          `protobuf:"bytes,10,opt,name=State,proto3" json:"State,omitempty"`
	ServingStatus        string          `protobuf:"bytes,11,opt,name=ServingStatus,proto3" json:"ServingStatus,omitempty"`
	Creative             *CreativeDTO    `protobuf:"bytes,12,opt,name=Creative,proto3" json:"Creative,omitempty"`
	LandingPage          *LandingPageDTO `protobuf:"bytes,13,opt,name=LandingPage,proto3" json:"LandingPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SBCampaignDTO) Reset()         { *m = SBCampaignDTO{} }
func (m *SBCampaignDTO) String() string { return proto.CompactTextString(m) }
func (*SBCampaignDTO) ProtoMessage()    {}
func (*SBCampaignDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{3}
}

func (m *SBCampaignDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SBCampaignDTO.Unmarshal(m, b)
}
func (m *SBCampaignDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SBCampaignDTO.Marshal(b, m, deterministic)
}
func (m *SBCampaignDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SBCampaignDTO.Merge(m, src)
}
func (m *SBCampaignDTO) XXX_Size() int {
	return xxx_messageInfo_SBCampaignDTO.Size(m)
}
func (m *SBCampaignDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_SBCampaignDTO.DiscardUnknown(m)
}

var xxx_messageInfo_SBCampaignDTO proto.InternalMessageInfo

func (m *SBCampaignDTO) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SBCampaignDTO) GetPortfolioID() int64 {
	if m != nil {
		return m.PortfolioID
	}
	return 0
}

func (m *SBCampaignDTO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SBCampaignDTO) GetBudget() float32 {
	if m != nil {
		return m.Budget
	}
	return 0
}

func (m *SBCampaignDTO) GetBidOptimization() bool {
	if m != nil {
		return m.BidOptimization
	}
	return false
}

func (m *SBCampaignDTO) GetAdFormat() string {
	if m != nil {
		return m.AdFormat
	}
	return ""
}

func (m *SBCampaignDTO) GetBudgetType() string {
	if m != nil {
		return m.BudgetType
	}
	return ""
}

func (m *SBCampaignDTO) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *SBCampaignDTO) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *SBCampaignDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *SBCampaignDTO) GetServingStatus() string {
	if m != nil {
		return m.ServingStatus
	}
	return ""
}

func (m *SBCampaignDTO) GetCreative() *CreativeDTO {
	if m != nil {
		return m.Creative
	}
	return nil
}

func (m *SBCampaignDTO) GetLandingPage() *LandingPageDTO {
	if m != nil {
		return m.LandingPage
	}
	return nil
}

type CreativeDTO struct {
	BrandName            string   `protobuf:"bytes,1,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	BrandLogoAssetID     string   `protobuf:"bytes,2,opt,name=BrandLogoAssetID,proto3" json:"BrandLogoAssetID,omitempty"`
	Headline             string   `protobuf:"bytes,3,opt,name=Headline,proto3" json:"Headline,omitempty"`
	ShouldOptimizeAsins  bool     `protobuf:"varint,4,opt,name=ShouldOptimizeAsins,proto3" json:"ShouldOptimizeAsins,omitempty"`
	ASINs                []string `protobuf:"bytes,5,rep,name=ASINs,proto3" json:"ASINs,omitempty"`
	BrandLogoURL         string   `protobuf:"bytes,6,opt,name=BrandLogoURL,proto3" json:"BrandLogoURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreativeDTO) Reset()         { *m = CreativeDTO{} }
func (m *CreativeDTO) String() string { return proto.CompactTextString(m) }
func (*CreativeDTO) ProtoMessage()    {}
func (*CreativeDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{4}
}

func (m *CreativeDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreativeDTO.Unmarshal(m, b)
}
func (m *CreativeDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreativeDTO.Marshal(b, m, deterministic)
}
func (m *CreativeDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreativeDTO.Merge(m, src)
}
func (m *CreativeDTO) XXX_Size() int {
	return xxx_messageInfo_CreativeDTO.Size(m)
}
func (m *CreativeDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_CreativeDTO.DiscardUnknown(m)
}

var xxx_messageInfo_CreativeDTO proto.InternalMessageInfo

func (m *CreativeDTO) GetBrandName() string {
	if m != nil {
		return m.BrandName
	}
	return ""
}

func (m *CreativeDTO) GetBrandLogoAssetID() string {
	if m != nil {
		return m.BrandLogoAssetID
	}
	return ""
}

func (m *CreativeDTO) GetHeadline() string {
	if m != nil {
		return m.Headline
	}
	return ""
}

func (m *CreativeDTO) GetShouldOptimizeAsins() bool {
	if m != nil {
		return m.ShouldOptimizeAsins
	}
	return false
}

func (m *CreativeDTO) GetASINs() []string {
	if m != nil {
		return m.ASINs
	}
	return nil
}

func (m *CreativeDTO) GetBrandLogoURL() string {
	if m != nil {
		return m.BrandLogoURL
	}
	return ""
}

type LandingPageDTO struct {
	PageType             string   `protobuf:"bytes,1,opt,name=PageType,proto3" json:"PageType,omitempty"`
	URL                  string   `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LandingPageDTO) Reset()         { *m = LandingPageDTO{} }
func (m *LandingPageDTO) String() string { return proto.CompactTextString(m) }
func (*LandingPageDTO) ProtoMessage()    {}
func (*LandingPageDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{5}
}

func (m *LandingPageDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LandingPageDTO.Unmarshal(m, b)
}
func (m *LandingPageDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LandingPageDTO.Marshal(b, m, deterministic)
}
func (m *LandingPageDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LandingPageDTO.Merge(m, src)
}
func (m *LandingPageDTO) XXX_Size() int {
	return xxx_messageInfo_LandingPageDTO.Size(m)
}
func (m *LandingPageDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_LandingPageDTO.DiscardUnknown(m)
}

var xxx_messageInfo_LandingPageDTO proto.InternalMessageInfo

func (m *LandingPageDTO) GetPageType() string {
	if m != nil {
		return m.PageType
	}
	return ""
}

func (m *LandingPageDTO) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type BidRecommendationsDTO struct {
	AdGroupID            int64                 `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"AdGroupID,omitempty"`
	SuggestedBid         *SuggestedBidDTO      `protobuf:"bytes,2,opt,name=SuggestedBid,proto3" json:"SuggestedBid,omitempty"`
	Recommendations      []*RecommendationsDTO `protobuf:"bytes,3,rep,name=Recommendations,proto3" json:"Recommendations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BidRecommendationsDTO) Reset()         { *m = BidRecommendationsDTO{} }
func (m *BidRecommendationsDTO) String() string { return proto.CompactTextString(m) }
func (*BidRecommendationsDTO) ProtoMessage()    {}
func (*BidRecommendationsDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{6}
}

func (m *BidRecommendationsDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidRecommendationsDTO.Unmarshal(m, b)
}
func (m *BidRecommendationsDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidRecommendationsDTO.Marshal(b, m, deterministic)
}
func (m *BidRecommendationsDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidRecommendationsDTO.Merge(m, src)
}
func (m *BidRecommendationsDTO) XXX_Size() int {
	return xxx_messageInfo_BidRecommendationsDTO.Size(m)
}
func (m *BidRecommendationsDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_BidRecommendationsDTO.DiscardUnknown(m)
}

var xxx_messageInfo_BidRecommendationsDTO proto.InternalMessageInfo

func (m *BidRecommendationsDTO) GetAdGroupID() int64 {
	if m != nil {
		return m.AdGroupID
	}
	return 0
}

func (m *BidRecommendationsDTO) GetSuggestedBid() *SuggestedBidDTO {
	if m != nil {
		return m.SuggestedBid
	}
	return nil
}

func (m *BidRecommendationsDTO) GetRecommendations() []*RecommendationsDTO {
	if m != nil {
		return m.Recommendations
	}
	return nil
}

type RecommendationsDTO struct {
	Expression           []*ExpressionDTO `protobuf:"bytes,1,rep,name=Expression,proto3" json:"Expression,omitempty"`
	SuggestedBid         *SuggestedBidDTO `protobuf:"bytes,2,opt,name=SuggestedBid,proto3" json:"SuggestedBid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RecommendationsDTO) Reset()         { *m = RecommendationsDTO{} }
func (m *RecommendationsDTO) String() string { return proto.CompactTextString(m) }
func (*RecommendationsDTO) ProtoMessage()    {}
func (*RecommendationsDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{7}
}

func (m *RecommendationsDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationsDTO.Unmarshal(m, b)
}
func (m *RecommendationsDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationsDTO.Marshal(b, m, deterministic)
}
func (m *RecommendationsDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationsDTO.Merge(m, src)
}
func (m *RecommendationsDTO) XXX_Size() int {
	return xxx_messageInfo_RecommendationsDTO.Size(m)
}
func (m *RecommendationsDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationsDTO.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationsDTO proto.InternalMessageInfo

func (m *RecommendationsDTO) GetExpression() []*ExpressionDTO {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (m *RecommendationsDTO) GetSuggestedBid() *SuggestedBidDTO {
	if m != nil {
		return m.SuggestedBid
	}
	return nil
}

type ExpressionDTO struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpressionDTO) Reset()         { *m = ExpressionDTO{} }
func (m *ExpressionDTO) String() string { return proto.CompactTextString(m) }
func (*ExpressionDTO) ProtoMessage()    {}
func (*ExpressionDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{8}
}

func (m *ExpressionDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressionDTO.Unmarshal(m, b)
}
func (m *ExpressionDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressionDTO.Marshal(b, m, deterministic)
}
func (m *ExpressionDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressionDTO.Merge(m, src)
}
func (m *ExpressionDTO) XXX_Size() int {
	return xxx_messageInfo_ExpressionDTO.Size(m)
}
func (m *ExpressionDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressionDTO.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressionDTO proto.InternalMessageInfo

func (m *ExpressionDTO) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ExpressionDTO) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SuggestedBidDTO struct {
	End                  string   `protobuf:"bytes,1,opt,name=End,proto3" json:"End,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=Start,proto3" json:"Start,omitempty"`
	Suggested            string   `protobuf:"bytes,3,opt,name=Suggested,proto3" json:"Suggested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuggestedBidDTO) Reset()         { *m = SuggestedBidDTO{} }
func (m *SuggestedBidDTO) String() string { return proto.CompactTextString(m) }
func (*SuggestedBidDTO) ProtoMessage()    {}
func (*SuggestedBidDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{9}
}

func (m *SuggestedBidDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestedBidDTO.Unmarshal(m, b)
}
func (m *SuggestedBidDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestedBidDTO.Marshal(b, m, deterministic)
}
func (m *SuggestedBidDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestedBidDTO.Merge(m, src)
}
func (m *SuggestedBidDTO) XXX_Size() int {
	return xxx_messageInfo_SuggestedBidDTO.Size(m)
}
func (m *SuggestedBidDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestedBidDTO.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestedBidDTO proto.InternalMessageInfo

func (m *SuggestedBidDTO) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *SuggestedBidDTO) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *SuggestedBidDTO) GetSuggested() string {
	if m != nil {
		return m.Suggested
	}
	return ""
}

type AdGroupDTO struct {
	CampaignID           int64    `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"CampaignID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	DefaultBid           float32  `protobuf:"fixed32,4,opt,name=DefaultBid,proto3" json:"DefaultBid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupDTO) Reset()         { *m = AdGroupDTO{} }
func (m *AdGroupDTO) String() string { return proto.CompactTextString(m) }
func (*AdGroupDTO) ProtoMessage()    {}
func (*AdGroupDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{10}
}

func (m *AdGroupDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupDTO.Unmarshal(m, b)
}
func (m *AdGroupDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupDTO.Marshal(b, m, deterministic)
}
func (m *AdGroupDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupDTO.Merge(m, src)
}
func (m *AdGroupDTO) XXX_Size() int {
	return xxx_messageInfo_AdGroupDTO.Size(m)
}
func (m *AdGroupDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupDTO.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupDTO proto.InternalMessageInfo

func (m *AdGroupDTO) GetCampaignID() int64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *AdGroupDTO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdGroupDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AdGroupDTO) GetDefaultBid() float32 {
	if m != nil {
		return m.DefaultBid
	}
	return 0
}

type ProductAdDTO struct {
	AdGroupID            int64    `protobuf:"varint,1,opt,name=AdGroupID,proto3" json:"AdGroupID,omitempty"`
	CampaignID           int64    `protobuf:"varint,2,opt,name=CampaignID,proto3" json:"CampaignID,omitempty"`
	SKU                  string   `protobuf:"bytes,3,opt,name=SKU,proto3" json:"SKU,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductAdDTO) Reset()         { *m = ProductAdDTO{} }
func (m *ProductAdDTO) String() string { return proto.CompactTextString(m) }
func (*ProductAdDTO) ProtoMessage()    {}
func (*ProductAdDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{11}
}

func (m *ProductAdDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductAdDTO.Unmarshal(m, b)
}
func (m *ProductAdDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductAdDTO.Marshal(b, m, deterministic)
}
func (m *ProductAdDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductAdDTO.Merge(m, src)
}
func (m *ProductAdDTO) XXX_Size() int {
	return xxx_messageInfo_ProductAdDTO.Size(m)
}
func (m *ProductAdDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductAdDTO.DiscardUnknown(m)
}

var xxx_messageInfo_ProductAdDTO proto.InternalMessageInfo

func (m *ProductAdDTO) GetAdGroupID() int64 {
	if m != nil {
		return m.AdGroupID
	}
	return 0
}

func (m *ProductAdDTO) GetCampaignID() int64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *ProductAdDTO) GetSKU() string {
	if m != nil {
		return m.SKU
	}
	return ""
}

func (m *ProductAdDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type KeywordDTO struct {
	CampaignID           int64    `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"CampaignID,omitempty"`
	AdGroupID            int64    `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"AdGroupID,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	KeywordText          string   `protobuf:"bytes,4,opt,name=KeywordText,proto3" json:"KeywordText,omitempty"`
	MatchType            string   `protobuf:"bytes,5,opt,name=MatchType,proto3" json:"MatchType,omitempty"`
	Bid                  float32  `protobuf:"fixed32,6,opt,name=Bid,proto3" json:"Bid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordDTO) Reset()         { *m = KeywordDTO{} }
func (m *KeywordDTO) String() string { return proto.CompactTextString(m) }
func (*KeywordDTO) ProtoMessage()    {}
func (*KeywordDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{12}
}

func (m *KeywordDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordDTO.Unmarshal(m, b)
}
func (m *KeywordDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordDTO.Marshal(b, m, deterministic)
}
func (m *KeywordDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordDTO.Merge(m, src)
}
func (m *KeywordDTO) XXX_Size() int {
	return xxx_messageInfo_KeywordDTO.Size(m)
}
func (m *KeywordDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordDTO.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordDTO proto.InternalMessageInfo

func (m *KeywordDTO) GetCampaignID() int64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *KeywordDTO) GetAdGroupID() int64 {
	if m != nil {
		return m.AdGroupID
	}
	return 0
}

func (m *KeywordDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *KeywordDTO) GetKeywordText() string {
	if m != nil {
		return m.KeywordText
	}
	return ""
}

func (m *KeywordDTO) GetMatchType() string {
	if m != nil {
		return m.MatchType
	}
	return ""
}

func (m *KeywordDTO) GetBid() float32 {
	if m != nil {
		return m.Bid
	}
	return 0
}

type TargetDTO struct {
	CampaignID           int64            `protobuf:"varint,1,opt,name=CampaignID,proto3" json:"CampaignID,omitempty"`
	AdGroupID            int64            `protobuf:"varint,2,opt,name=AdGroupID,proto3" json:"AdGroupID,omitempty"`
	TargetID             int64            `protobuf:"varint,3,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	ExpressionType       string           `protobuf:"bytes,4,opt,name=ExpressionType,proto3" json:"ExpressionType,omitempty"`
	Expression           []*ExpressionDTO `protobuf:"bytes,5,rep,name=Expression,proto3" json:"Expression,omitempty"`
	Bid                  float32          `protobuf:"fixed32,6,opt,name=Bid,proto3" json:"Bid,omitempty"`
	State                string           `protobuf:"bytes,7,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TargetDTO) Reset()         { *m = TargetDTO{} }
func (m *TargetDTO) String() string { return proto.CompactTextString(m) }
func (*TargetDTO) ProtoMessage()    {}
func (*TargetDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe20d52a0da76480, []int{13}
}

func (m *TargetDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetDTO.Unmarshal(m, b)
}
func (m *TargetDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetDTO.Marshal(b, m, deterministic)
}
func (m *TargetDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetDTO.Merge(m, src)
}
func (m *TargetDTO) XXX_Size() int {
	return xxx_messageInfo_TargetDTO.Size(m)
}
func (m *TargetDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetDTO.DiscardUnknown(m)
}

var xxx_messageInfo_TargetDTO proto.InternalMessageInfo

func (m *TargetDTO) GetCampaignID() int64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *TargetDTO) GetAdGroupID() int64 {
	if m != nil {
		return m.AdGroupID
	}
	return 0
}

func (m *TargetDTO) GetTargetID() int64 {
	if m != nil {
		return m.TargetID
	}
	return 0
}

func (m *TargetDTO) GetExpressionType() string {
	if m != nil {
		return m.ExpressionType
	}
	return ""
}

func (m *TargetDTO) GetExpression() []*ExpressionDTO {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (m *TargetDTO) GetBid() float32 {
	if m != nil {
		return m.Bid
	}
	return 0
}

func (m *TargetDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*SPCampaignDTO)(nil), "campaignsmodel.SPCampaignDTO")
	proto.RegisterType((*BiddingDTO)(nil), "campaignsmodel.BiddingDTO")
	proto.RegisterType((*AdjustmentDTO)(nil), "campaignsmodel.AdjustmentDTO")
	proto.RegisterType((*SBCampaignDTO)(nil), "campaignsmodel.SBCampaignDTO")
	proto.RegisterType((*CreativeDTO)(nil), "campaignsmodel.CreativeDTO")
	proto.RegisterType((*LandingPageDTO)(nil), "campaignsmodel.LandingPageDTO")
	proto.RegisterType((*BidRecommendationsDTO)(nil), "campaignsmodel.BidRecommendationsDTO")
	proto.RegisterType((*RecommendationsDTO)(nil), "campaignsmodel.RecommendationsDTO")
	proto.RegisterType((*ExpressionDTO)(nil), "campaignsmodel.ExpressionDTO")
	proto.RegisterType((*SuggestedBidDTO)(nil), "campaignsmodel.SuggestedBidDTO")
	proto.RegisterType((*AdGroupDTO)(nil), "campaignsmodel.AdGroupDTO")
	proto.RegisterType((*ProductAdDTO)(nil), "campaignsmodel.ProductAdDTO")
	proto.RegisterType((*KeywordDTO)(nil), "campaignsmodel.KeywordDTO")
	proto.RegisterType((*TargetDTO)(nil), "campaignsmodel.TargetDTO")
}

func init() {
	proto.RegisterFile("campaignsmodel/campaignsmodel.proto", fileDescriptor_fe20d52a0da76480)
}

var fileDescriptor_fe20d52a0da76480 = []byte{
	// 956 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xed, 0xb6, 0x49, 0x4e, 0x9a, 0x6e, 0x35, 0x2c, 0xc8, 0x2a, 0x50, 0x2c, 0xb3, 0x42,
	0x11, 0x17, 0x2d, 0x0a, 0x2b, 0x21, 0x2e, 0xf8, 0x89, 0xeb, 0x02, 0xd1, 0x76, 0xb7, 0x91, 0x9d,
	0x05, 0x89, 0xbb, 0x59, 0xcf, 0xd4, 0x1d, 0x14, 0xdb, 0xd1, 0x78, 0x5c, 0x36, 0xfb, 0x24, 0x3c,
	0x07, 0x8f, 0xc0, 0x15, 0x6f, 0xc1, 0x4b, 0x70, 0xc3, 0x1d, 0x9a, 0xf1, 0xdf, 0xd8, 0x8d, 0x04,
	0x88, 0xde, 0xcd, 0xf9, 0xce, 0x9c, 0xdf, 0xef, 0xcc, 0xd1, 0xc0, 0x87, 0x11, 0x4e, 0x36, 0x98,
	0xc5, 0x69, 0x9e, 0x64, 0x84, 0xae, 0xcf, 0xbb, 0xe2, 0xd9, 0x86, 0x67, 0x22, 0x43, 0x47, 0x5d,
	0xd4, 0xfd, 0xcb, 0x84, 0x49, 0xb8, 0xbc, 0xa8, 0x40, 0x7f, 0x75, 0x8d, 0x8e, 0xc0, 0x5c, 0xf8,
	0xb6, 0xe1, 0x18, 0x53, 0x2b, 0x30, 0x17, 0x3e, 0x72, 0x60, 0xbc, 0xcc, 0xb8, 0xb8, 0xc9, 0xd6,
	0x2c, 0x5b, 0xf8, 0xb6, 0xa9, 0x14, 0x3a, 0x84, 0x10, 0xec, 0xbd, 0xc0, 0x09, 0xb5, 0x2d, 0xc7,
	0x98, 0x8e, 0x02, 0x75, 0x46, 0x2e, 0x1c, 0xd6, 0x4e, 0x57, 0xdb, 0x0d, 0xb5, 0xf7, 0x94, 0xae,
	0x83, 0xa1, 0x27, 0x30, 0x59, 0x61, 0x1e, 0x53, 0xc1, 0xd2, 0x58, 0x5d, 0xda, 0x57, 0x97, 0xba,
	0x20, 0x9a, 0xc1, 0xe3, 0x25, 0xa7, 0x09, 0x2b, 0x12, 0x8f, 0x91, 0x39, 0xf9, 0xa9, 0xc8, 0x45,
	0x42, 0x53, 0x61, 0x1f, 0x38, 0xc6, 0x74, 0x18, 0xec, 0xd4, 0xc9, 0x9c, 0x7d, 0xcc, 0xd6, 0x5b,
	0xaf, 0x20, 0x31, 0x15, 0xf6, 0xc0, 0x31, 0xa6, 0x66, 0xa0, 0x43, 0xe8, 0x3d, 0x18, 0x85, 0x02,
	0x73, 0xe1, 0x63, 0x41, 0xed, 0xa1, 0x8a, 0xdb, 0x02, 0xc8, 0x86, 0xc1, 0x65, 0x4a, 0x94, 0x6e,
	0xa4, 0x74, 0xb5, 0x88, 0x1e, 0xc3, 0x7e, 0x28, 0x24, 0x0e, 0x0a, 0x2f, 0x05, 0xf4, 0x14, 0x06,
	0x1e, 0x23, 0x84, 0xa5, 0xb1, 0x3d, 0x76, 0x8c, 0xe9, 0x78, 0x76, 0x72, 0xd6, 0xeb, 0x7e, 0xa5,
	0xf6, 0x57, 0xd7, 0x41, 0x7d, 0xd5, 0x65, 0x00, 0x2d, 0x8c, 0x4e, 0x60, 0x18, 0x0a, 0x8e, 0x05,
	0x8d, 0xb7, 0xaa, 0xfb, 0xa3, 0xa0, 0x91, 0xd1, 0x57, 0x30, 0x6e, 0xab, 0xcb, 0x6d, 0xd3, 0xb1,
	0xa6, 0xe3, 0xd9, 0xfb, 0xfd, 0x18, 0xed, 0x15, 0x19, 0x46, 0xb7, 0x70, 0x9f, 0xc3, 0xa4, 0xa3,
	0x95, 0xf5, 0x2f, 0x39, 0x25, 0x2c, 0x92, 0xb5, 0x94, 0xe1, 0x5a, 0x00, 0x9d, 0x02, 0x2c, 0x29,
	0x8f, 0x68, 0x2a, 0x70, 0x4c, 0x15, 0xe5, 0x66, 0xa0, 0x21, 0xee, 0x6f, 0x16, 0x4c, 0x42, 0xef,
	0xe1, 0xa7, 0xe6, 0x1d, 0x38, 0xa8, 0x28, 0xdb, 0x53, 0x31, 0x2b, 0x09, 0x4d, 0xe1, 0x91, 0xc7,
	0xc8, 0xf5, 0x46, 0xb0, 0x84, 0xbd, 0xc1, 0x82, 0x65, 0xa9, 0x9a, 0x95, 0x61, 0xd0, 0x87, 0x65,
	0x17, 0xe7, 0xe4, 0x9b, 0x8c, 0x27, 0xb8, 0x9c, 0x90, 0x51, 0xd0, 0xc8, 0xb2, 0xaa, 0xd2, 0x9f,
	0x1a, 0xb6, 0x81, 0xd2, 0x6a, 0xc8, 0x03, 0xcf, 0xc4, 0x13, 0x98, 0x84, 0x94, 0xdf, 0xb1, 0x34,
	0x96, 0x72, 0x91, 0xab, 0xc9, 0x18, 0x05, 0x5d, 0x10, 0x7d, 0x06, 0xc3, 0x0b, 0x4e, 0xb1, 0x60,
	0x77, 0xd4, 0x3e, 0x54, 0xa3, 0xf3, 0x6e, 0x9f, 0xd6, 0x5a, 0x2f, 0x49, 0x6d, 0x2e, 0xa3, 0xaf,
	0x61, 0x7c, 0x85, 0x53, 0x39, 0x3c, 0x4b, 0xc9, 0xd1, 0x44, 0xd9, 0x9e, 0xf6, 0x6d, 0xb5, 0x2b,
	0x6a, 0x26, 0x34, 0xd9, 0xfd, 0xc3, 0x80, 0xb1, 0xe6, 0x5b, 0x96, 0xef, 0x71, 0x9c, 0x12, 0xc5,
	0x4a, 0x35, 0x12, 0x0d, 0x80, 0x3e, 0x86, 0x63, 0x25, 0x5c, 0x65, 0x71, 0x36, 0xcf, 0x73, 0x2a,
	0x2a, 0x56, 0x47, 0xc1, 0x3d, 0x5c, 0x92, 0xf0, 0x1d, 0xc5, 0x64, 0xcd, 0xd2, 0x9a, 0xde, 0x46,
	0x46, 0x9f, 0xc0, 0x5b, 0xe1, 0x6d, 0x56, 0xac, 0x6b, 0xda, 0xe8, 0x3c, 0x67, 0x69, 0xae, 0xf8,
	0x1e, 0x06, 0xbb, 0x54, 0xb2, 0xbd, 0xf3, 0x70, 0xf1, 0x22, 0xb7, 0xf7, 0x1d, 0x4b, 0xb6, 0x57,
	0x09, 0x72, 0xc1, 0x34, 0x71, 0x5f, 0x06, 0x57, 0x15, 0xd9, 0x1d, 0xcc, 0xfd, 0x12, 0x8e, 0xba,
	0x0d, 0x90, 0x99, 0xc9, 0xa3, 0x1a, 0x80, 0xea, 0x91, 0xd5, 0x32, 0x3a, 0x06, 0x4b, 0x3a, 0x2a,
	0x8b, 0x92, 0x47, 0xf7, 0x77, 0x03, 0xde, 0xf6, 0x18, 0x09, 0x68, 0x94, 0x25, 0x09, 0x4d, 0x89,
	0x1a, 0xb1, 0xbc, 0xea, 0xd5, 0x9c, 0x7c, 0xcb, 0xb3, 0x62, 0xd3, 0x4c, 0x7d, 0x0b, 0xa0, 0x0b,
	0x38, 0x0c, 0x8b, 0x38, 0xa6, 0xb9, 0xa0, 0xc4, 0x63, 0x44, 0xb9, 0x1c, 0xcf, 0x3e, 0xe8, 0x93,
	0xa3, 0xdf, 0x91, 0xec, 0x74, 0x8c, 0xd0, 0x15, 0x3c, 0xea, 0x05, 0xb6, 0x2d, 0xf5, 0xee, 0xdd,
	0xbe, 0x9f, 0xfb, 0xf9, 0x05, 0x7d, 0x53, 0xf7, 0x17, 0x03, 0xd0, 0x8e, 0x3a, 0xbe, 0x00, 0xb8,
	0x7c, 0xbd, 0xe1, 0x34, 0xcf, 0xe5, 0x9b, 0x32, 0x76, 0xef, 0x95, 0xf6, 0x86, 0x74, 0xad, 0x19,
	0x3c, 0x48, 0xa1, 0xee, 0xe7, 0x30, 0xe9, 0x44, 0x90, 0x84, 0x7f, 0x8f, 0xd7, 0x45, 0xcd, 0x50,
	0x29, 0xc8, 0x7d, 0xa1, 0x68, 0x2b, 0xf9, 0x51, 0x67, 0xf7, 0x07, 0x78, 0xd4, 0xf3, 0x2d, 0x59,
	0xbc, 0x4c, 0x49, 0x65, 0x2a, 0x8f, 0xd5, 0xf3, 0xe4, 0xa2, 0xb2, 0x2c, 0x05, 0xf5, 0xd8, 0x6b,
	0xd3, 0x6a, 0x48, 0x5b, 0xc0, 0xbd, 0x03, 0xa8, 0xe8, 0x94, 0x3e, 0x4f, 0x01, 0xea, 0x5d, 0xd7,
	0xd0, 0xad, 0x21, 0xcd, 0x2a, 0x33, 0xb5, 0x55, 0xd6, 0x2c, 0x05, 0x4b, 0x5f, 0x0a, 0xa7, 0x00,
	0x3e, 0xbd, 0xc1, 0xc5, 0x5a, 0xc8, 0x76, 0x95, 0x4b, 0x4e, 0x43, 0x5c, 0x01, 0x87, 0x4b, 0x9e,
	0x91, 0x22, 0x12, 0x73, 0xf2, 0xcf, 0x73, 0xd6, 0xcd, 0xcb, 0xbc, 0x97, 0xd7, 0x31, 0x58, 0xe1,
	0xb3, 0x97, 0x55, 0x06, 0xf2, 0xd8, 0x66, 0xb5, 0xa7, 0x65, 0xe5, 0xfe, 0x6a, 0x00, 0x3c, 0xa3,
	0xdb, 0x9f, 0x33, 0x4e, 0xfe, 0x4d, 0xb9, 0x9d, 0xa4, 0xcc, 0x7e, 0x52, 0xbb, 0x0b, 0x77, 0x60,
	0x5c, 0x45, 0x58, 0xd1, 0xd7, 0xa2, 0x0a, 0xaf, 0x43, 0xd2, 0xeb, 0x73, 0x2c, 0xa2, 0x5b, 0xed,
	0x27, 0xd0, 0x02, 0xb2, 0x14, 0xd9, 0xb1, 0x03, 0xd5, 0x31, 0x79, 0x74, 0xff, 0x34, 0x60, 0x54,
	0xfe, 0x14, 0xfe, 0x7f, 0xce, 0x27, 0x30, 0x2c, 0x5d, 0x2d, 0x7c, 0x95, 0xb6, 0x15, 0x34, 0x32,
	0xfa, 0x08, 0x8e, 0xda, 0xf1, 0xd4, 0xfe, 0x32, 0x3d, 0xb4, 0xf7, 0x94, 0xf6, 0xff, 0xeb, 0x53,
	0xba, 0x57, 0x60, 0xdb, 0xc8, 0x81, 0xd6, 0x48, 0xef, 0xe9, 0x8f, 0xb3, 0x98, 0x89, 0xdb, 0xe2,
	0xd5, 0x59, 0x94, 0x25, 0xe7, 0xf9, 0x36, 0x8d, 0x6e, 0x0a, 0x51, 0x70, 0x7a, 0x8e, 0x93, 0x37,
	0x98, 0xdc, 0x9d, 0xab, 0x6f, 0x5e, 0xd4, 0xfb, 0xfc, 0xbd, 0x3a, 0x50, 0xf0, 0xa7, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x7a, 0x35, 0x66, 0x24, 0x0a, 0x00, 0x00,
}