package client

import (
	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
)

// Account-related types
type GetAdvertisersRequest struct {
	Fields        []string `json:"fields,omitempty"`
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	Page          int      `json:"page,omitempty"`
	PageSize      int      `json:"page_size,omitempty"`
}

type GetAdvertisersResponse struct {
	models.ListResponse
	Data []AdvertiserInfo `json:"data"`
}

type AdvertiserInfo struct {
	AdvertiserID            string  `json:"advertiser_id"`
	AdvertiserName          string  `json:"advertiser_name"`
	Status                  string  `json:"status"`
	Currency                string  `json:"currency"`
	Timezone                string  `json:"timezone"`
	CompanyName             string  `json:"company_name,omitempty"`
	Industry                string  `json:"industry,omitempty"`
	Language                string  `json:"language,omitempty"`
	ContactName             string  `json:"contact_name,omitempty"`
	ContactEmail            string  `json:"contact_email,omitempty"`
	ContactPhone            string  `json:"contact_phone,omitempty"`
	Address                 string  `json:"address,omitempty"`
	LicenseNo               string  `json:"license_no,omitempty"`
	LicenseURL              string  `json:"license_url,omitempty"`
	PromotionCenterCity     string  `json:"promotion_center_city,omitempty"`
	PromotionCenterProvince string  `json:"promotion_center_province,omitempty"`
	Balance                 float64 `json:"balance,omitempty"`
	CreateTime              string  `json:"create_time,omitempty"`
	Role                    string  `json:"role,omitempty"`
}

type UpdateAdvertiserRequest struct {
	AdvertiserID   string `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name,omitempty"`
	CompanyName    string `json:"company_name,omitempty"`
	ContactName    string `json:"contact_name,omitempty"`
	ContactEmail   string `json:"contact_email,omitempty"`
	ContactPhone   string `json:"contact_phone,omitempty"`
	Address        string `json:"address,omitempty"`
}

type UpdateAdvertiserResponse struct {
	models.BaseResponse
	Data struct {
		AdvertiserID string `json:"advertiser_id"`
	} `json:"data"`
}

type GetAdvertiserBalanceRequest struct {
	AdvertiserID string `json:"advertiser_id"`
}

type GetAdvertiserBalanceResponse struct {
	models.BaseResponse
	Data struct {
		AdvertiserID string  `json:"advertiser_id"`
		Balance      float64 `json:"balance"`
		Currency     string  `json:"currency"`
	} `json:"data"`
}

type GetAdvertiserFundRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	FundTypes    []string `json:"fund_types,omitempty"`
}

type GetAdvertiserFundResponse struct {
	models.BaseResponse
	Data []FundInfo `json:"data"`
}

type FundInfo struct {
	FundType   string  `json:"fund_type"`
	Balance    float64 `json:"balance"`
	Currency   string  `json:"currency"`
	ValidStart string  `json:"valid_start,omitempty"`
	ValidEnd   string  `json:"valid_end,omitempty"`
}

type CreateAdvertiserRequest struct {
	AdvertiserName string `json:"advertiser_name"`
	CompanyName    string `json:"company_name"`
	Industry       string `json:"industry"`
	Currency       string `json:"currency"`
	Timezone       string `json:"timezone"`
	ContactName    string `json:"contact_name"`
	ContactEmail   string `json:"contact_email"`
	ContactPhone   string `json:"contact_phone"`
	Address        string `json:"address"`
	LicenseNo      string `json:"license_no,omitempty"`
	LicenseURL     string `json:"license_url,omitempty"`
}

type CreateAdvertiserResponse struct {
	models.BaseResponse
	Data struct {
		AdvertiserID string `json:"advertiser_id"`
	} `json:"data"`
}

// Campaign-related types
type CampaignCreateRequest struct {
	AdvertiserID      string               `json:"advertiser_id"`
	CampaignName      string               `json:"campaign_name"`
	ObjectiveType     models.ObjectiveType `json:"objective_type"`
	Budget            float64              `json:"budget"`
	BudgetMode        models.BudgetMode    `json:"budget_mode"`
	AppPromotionType  string               `json:"app_promotion_type,omitempty"`
	DeepBidType       string               `json:"deep_bid_type,omitempty"`
	CampaignType      string               `json:"campaign_type,omitempty"`
	SpecialIndustries []string             `json:"special_industries,omitempty"`
}

type CampaignCreateResponse struct {
	models.BaseResponse
	Data struct {
		CampaignID string `json:"campaign_id"`
	} `json:"data"`
}

type CampaignGetRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	CampaignIDs  []string `json:"campaign_ids,omitempty"`
	Fields       []string `json:"fields,omitempty"`
	Page         int      `json:"page,omitempty"`
	PageSize     int      `json:"page_size,omitempty"`
}

type CampaignGetResponse struct {
	models.ListResponse
	Data []CampaignInfo `json:"data"`
}

type CampaignInfo struct {
	CampaignID        string   `json:"campaign_id"`
	CampaignName      string   `json:"campaign_name"`
	AdvertiserID      string   `json:"advertiser_id"`
	Status            string   `json:"status"`
	ObjectiveType     string   `json:"objective_type"`
	Budget            float64  `json:"budget"`
	BudgetMode        string   `json:"budget_mode"`
	AppPromotionType  string   `json:"app_promotion_type,omitempty"`
	DeepBidType       string   `json:"deep_bid_type,omitempty"`
	CampaignType      string   `json:"campaign_type,omitempty"`
	SpecialIndustries []string `json:"special_industries,omitempty"`
	CreateTime        string   `json:"create_time,omitempty"`
	ModifyTime        string   `json:"modify_time,omitempty"`
}

type CampaignUpdateRequest struct {
	AdvertiserID string  `json:"advertiser_id"`
	CampaignID   string  `json:"campaign_id"`
	CampaignName string  `json:"campaign_name,omitempty"`
	Budget       float64 `json:"budget,omitempty"`
	BudgetMode   string  `json:"budget_mode,omitempty"`
}

type CampaignUpdateResponse struct {
	models.BaseResponse
	Data struct {
		CampaignID string `json:"campaign_id"`
	} `json:"data"`
}

type CampaignDeleteRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	CampaignIDs  []string `json:"campaign_ids"`
}

type CampaignDeleteResponse struct {
	models.BaseResponse
	Data struct {
		CampaignIDs []string `json:"campaign_ids"`
	} `json:"data"`
}

type CampaignStatusUpdateRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	CampaignIDs  []string `json:"campaign_ids"`
	Operation    string   `json:"operation"` // ENABLE, DISABLE, DELETE
}

type CampaignStatusUpdateResponse struct {
	models.BaseResponse
	Data struct {
		CampaignIDs []string `json:"campaign_ids"`
	} `json:"data"`
}

// Authentication types
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type TokenValidationResponse struct {
	Valid     bool   `json:"valid"`
	ExpiresAt int64  `json:"expires_at"`
	Scope     string `json:"scope"`
}

// Type definitions for services not yet fully implemented

type AdCreateRequest struct{}
type AdCreateResponse struct{}
type AdGetRequest struct{}
type AdGetResponse struct{}
type AdUpdateRequest struct{}
type AdUpdateResponse struct{}
type AdDeleteRequest struct{}
type AdDeleteResponse struct{}
type AdStatusUpdateRequest struct{}
type AdStatusUpdateResponse struct{}

type AdGroupCreateRequest struct{}
type AdGroupCreateResponse struct{}
type AdGroupGetRequest struct{}
type AdGroupGetResponse struct{}
type AdGroupUpdateRequest struct{}
type AdGroupUpdateResponse struct{}
type AdGroupDeleteRequest struct{}
type AdGroupDeleteResponse struct{}
type AdGroupStatusUpdateRequest struct{}
type AdGroupStatusUpdateResponse struct{}

type CustomAudienceCreateRequest struct{}
type CustomAudienceCreateResponse struct{}
type CustomAudienceGetRequest struct{}
type CustomAudienceGetResponse struct{}
type CustomAudienceUpdateRequest struct{}
type CustomAudienceUpdateResponse struct{}
type CustomAudienceDeleteRequest struct{}
type CustomAudienceDeleteResponse struct{}
type LookalikeAudienceCreateRequest struct{}
type LookalikeAudienceCreateResponse struct{}

type ImageUploadRequest struct{}
type ImageUploadResponse struct{}
type VideoUploadRequest struct{}
type VideoUploadResponse struct{}
type CreativeGetRequest struct{}
type CreativeGetResponse struct{}
type CreativeUpdateRequest struct{}
type CreativeUpdateResponse struct{}

type ReportingRequest struct{}
type ReportingResponse struct{}
type AudienceReportingRequest struct{}
type AudienceReportingResponse struct{}
type AsyncReportRequest struct{}
type AsyncReportResponse struct{}
type AsyncReportStatusResponse struct{}

// Tool API types
type LanguagesResponse struct {
	models.BaseResponse
	Data []LanguageInfo `json:"data"`
}

type LanguageInfo struct {
	LanguageCode string `json:"language_code"`
	LanguageName string `json:"language_name"`
}

type CurrenciesResponse struct {
	models.BaseResponse
	Data []CurrencyInfo `json:"data"`
}

type CurrencyInfo struct {
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
}

type RegionsResponse struct {
	models.BaseResponse
	Data []RegionInfo `json:"data"`
}

type RegionInfo struct {
	RegionCode string `json:"region_code"`
	RegionName string `json:"region_name"`
}

type InterestCategoriesRequest struct {
	AdvertiserID      string   `json:"advertiser_id"`
	Version           int      `json:"version,omitempty"`
	Language          string   `json:"language,omitempty"`
	SpecialIndustries []string `json:"special_industries,omitempty"`
}

type InterestCategoriesResponse struct {
	models.BaseResponse
	Data []InterestCategory `json:"data"`
}

type InterestCategory struct {
	InterestCategoryID   string             `json:"interest_category_id"`
	InterestCategoryName string             `json:"interest_category_name"`
	ParentCategoryID     string             `json:"parent_category_id,omitempty"`
	Level                int                `json:"level"`
	Children             []InterestCategory `json:"children,omitempty"`
}

type CarriersRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	LocationIDs  []string `json:"location_ids,omitempty"`
}

type CarriersResponse struct {
	models.BaseResponse
	Data []CarrierInfo `json:"data"`
}

type CarrierInfo struct {
	CarrierID   string `json:"carrier_id"`
	CarrierName string `json:"carrier_name"`
	LocationID  string `json:"location_id"`
}

type DeviceModelsRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	OSType       string `json:"os_type,omitempty"`
}

type DeviceModelsResponse struct {
	models.BaseResponse
	Data []DeviceModelInfo `json:"data"`
}

type DeviceModelInfo struct {
	DeviceModelID   string `json:"device_model_id"`
	DeviceModelName string `json:"device_model_name"`
	OSType          string `json:"os_type"`
}

type BusinessCentersResponse struct{}
type BusinessCenterInfo struct{}
type BCAdvertisersResponse struct{}
type TransferAdvertiserRequest struct{}
type TransferAdvertiserResponse struct{}
