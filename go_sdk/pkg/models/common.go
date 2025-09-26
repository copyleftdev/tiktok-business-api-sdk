package models

import (
	"time"
)

// BaseResponse represents the common structure of TikTok API responses
type BaseResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data,omitempty"`
}

// PaginationInfo contains pagination information for list responses
type PaginationInfo struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
	TotalPage  int  `json:"total_page"`
	HasMore    bool `json:"has_more"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	BaseResponse
	PageInfo PaginationInfo `json:"page_info"`
}

// Status represents the status of various entities
type Status string

const (
	StatusActive   Status = "ENABLE"
	StatusInactive Status = "DISABLE"
	StatusDeleted  Status = "DELETE"
	StatusPaused   Status = "PAUSE"
)

// ObjectiveType represents campaign objective types
type ObjectiveType string

const (
	ObjectiveReach          ObjectiveType = "REACH"
	ObjectiveTraffic        ObjectiveType = "TRAFFIC"
	ObjectiveVideoViews     ObjectiveType = "VIDEO_VIEWS"
	ObjectiveLeadGeneration ObjectiveType = "LEAD_GENERATION"
	ObjectiveAppPromotion   ObjectiveType = "APP_PROMOTION"
	ObjectiveConversions    ObjectiveType = "CONVERSIONS"
	ObjectiveProductSales   ObjectiveType = "PRODUCT_SALES"
	ObjectiveEngagement     ObjectiveType = "ENGAGEMENT"
)

// BudgetMode represents budget mode types
type BudgetMode string

const (
	BudgetModeTotal BudgetMode = "BUDGET_MODE_TOTAL"
	BudgetModeDaily BudgetMode = "BUDGET_MODE_DAY"
)

// OptimizationGoal represents optimization goal types
type OptimizationGoal string

const (
	OptimizationGoalClick      OptimizationGoal = "CLICK"
	OptimizationGoalReach      OptimizationGoal = "REACH"
	OptimizationGoalImpression OptimizationGoal = "IMPRESSION"
	OptimizationGoalInstall    OptimizationGoal = "INSTALL"
	OptimizationGoalConversion OptimizationGoal = "CONVERSION"
	OptimizationGoalValue      OptimizationGoal = "VALUE"
)

// BidType represents bidding strategy types
type BidType string

const (
	BidTypeNoBid   BidType = "NO_BID"
	BidTypeBidCap  BidType = "BID_TYPE_NO_BID"
	BidTypeMaxBid  BidType = "BID_TYPE_CUSTOM"
	BidTypeAutoBid BidType = "BID_TYPE_AUTO"
)

// PlacementType represents placement types
type PlacementType string

const (
	PlacementTypeAutomatic PlacementType = "PLACEMENT_TYPE_AUTOMATIC"
	PlacementTypeNormal    PlacementType = "PLACEMENT_TYPE_NORMAL"
)

// Placement represents available placements
type Placement string

const (
	PlacementTikTok          Placement = "PLACEMENT_TIKTOK"
	PlacementPangle          Placement = "PLACEMENT_PANGLE"
	PlacementGlobalAppBundle Placement = "PLACEMENT_GLOBAL_APP_BUNDLE"
)

// Gender represents gender targeting options
type Gender string

const (
	GenderMale      Gender = "GENDER_MALE"
	GenderFemale    Gender = "GENDER_FEMALE"
	GenderUnlimited Gender = "GENDER_UNLIMITED"
)

// AgeGroup represents age group targeting
type AgeGroup string

const (
	Age13To17 AgeGroup = "AGE_13_17"
	Age18To24 AgeGroup = "AGE_18_24"
	Age25To34 AgeGroup = "AGE_25_34"
	Age35To44 AgeGroup = "AGE_35_44"
	Age45To54 AgeGroup = "AGE_45_54"
	Age55Plus AgeGroup = "AGE_55_PLUS"
)

// CreativeType represents creative asset types
type CreativeType string

const (
	CreativeTypeImage CreativeType = "IMAGE"
	CreativeTypeVideo CreativeType = "VIDEO"
)

// FileType represents file types for uploads
type FileType string

const (
	FileTypeImage FileType = "IMAGE"
	FileTypeVideo FileType = "VIDEO"
	FileTypeAudio FileType = "AUDIO"
)

// ReportType represents report types
type ReportType string

const (
	ReportTypeBasic    ReportType = "BASIC"
	ReportTypeAudience ReportType = "AUDIENCE"
	ReportTypeVideo    ReportType = "VIDEO"
)

// DataLevel represents data aggregation levels for reports
type DataLevel string

const (
	DataLevelCampaign DataLevel = "AUCTION_CAMPAIGN"
	DataLevelAdGroup  DataLevel = "AUCTION_ADGROUP"
	DataLevelAd       DataLevel = "AUCTION_AD"
	DataLevelCreative DataLevel = "AUCTION_CREATIVE"
)

// Dimension represents report dimensions
type Dimension string

const (
	DimensionCampaignID   Dimension = "campaign_id"
	DimensionAdGroupID    Dimension = "adgroup_id"
	DimensionAdID         Dimension = "ad_id"
	DimensionStatTimeDay  Dimension = "stat_time_day"
	DimensionStatTimeHour Dimension = "stat_time_hour"
)

// Metric represents report metrics
type Metric string

const (
	MetricImpressions Metric = "impressions"
	MetricClicks      Metric = "clicks"
	MetricCost        Metric = "cost"
	MetricCTR         Metric = "ctr"
	MetricCPC         Metric = "cpc"
	MetricCPM         Metric = "cpm"
	MetricConversions Metric = "conversions"
	MetricCPA         Metric = "cpa"
	MetricROAS        Metric = "roas"
)

// TimestampField represents a timestamp field that can be marshaled/unmarshaled
type TimestampField struct {
	time.Time
}

// MarshalJSON implements json.Marshaler
func (t TimestampField) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format("2006-01-02 15:04:05") + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (t *TimestampField) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	parsed, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

// DateField represents a date field that can be marshaled/unmarshaled
type DateField struct {
	time.Time
}

// MarshalJSON implements json.Marshaler
func (d DateField) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (d *DateField) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	parsed, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}

	d.Time = parsed
	return nil
}

// Money represents a monetary value with currency
type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// TargetingInfo represents common targeting information
type TargetingInfo struct {
	Genders            []Gender   `json:"genders,omitempty"`
	AgeGroups          []AgeGroup `json:"age_groups,omitempty"`
	Locations          []Location `json:"locations,omitempty"`
	Languages          []string   `json:"languages,omitempty"`
	InterestCategories []string   `json:"interest_categories,omitempty"`
	Behaviors          []string   `json:"behaviors,omitempty"`
	CustomAudiences    []string   `json:"custom_audiences,omitempty"`
}

// Location represents a geographic location for targeting
type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// CreativeInfo represents creative asset information
type CreativeInfo struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         CreativeType `json:"type"`
	URL          string       `json:"url"`
	ThumbnailURL string       `json:"thumbnail_url,omitempty"`
	Duration     int          `json:"duration,omitempty"`
	Width        int          `json:"width,omitempty"`
	Height       int          `json:"height,omitempty"`
	FileSize     int64        `json:"file_size,omitempty"`
}

// ValidationResult represents the result of a validation operation
type ValidationResult struct {
	Valid  bool     `json:"valid"`
	Errors []string `json:"errors,omitempty"`
}
