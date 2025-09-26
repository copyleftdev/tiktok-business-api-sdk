package client

import (
	"context"
	"io"
)

// AccountService defines the interface for account-related operations
type AccountService interface {
	// GetAdvertisers retrieves a list of advertisers
	GetAdvertisers(ctx context.Context, req *GetAdvertisersRequest) (*GetAdvertisersResponse, error)

	// GetAdvertiserInfo retrieves information about a specific advertiser
	GetAdvertiserInfo(ctx context.Context, advertiserID string) (*AdvertiserInfo, error)

	// UpdateAdvertiser updates advertiser information
	UpdateAdvertiser(ctx context.Context, req *UpdateAdvertiserRequest) (*UpdateAdvertiserResponse, error)

	// CreateAdvertiser creates a new advertiser account
	CreateAdvertiser(ctx context.Context, req *CreateAdvertiserRequest) (*CreateAdvertiserResponse, error)

	// GetAdvertiserBalance retrieves advertiser account balance
	GetAdvertiserBalance(ctx context.Context, req *GetAdvertiserBalanceRequest) (*GetAdvertiserBalanceResponse, error)

	// GetAdvertiserFund retrieves advertiser fund information
	GetAdvertiserFund(ctx context.Context, req *GetAdvertiserFundRequest) (*GetAdvertiserFundResponse, error)
}

// CampaignService defines the interface for campaign-related operations
type CampaignService interface {
	// Create creates a new campaign
	Create(ctx context.Context, req *CampaignCreateRequest) (*CampaignCreateResponse, error)

	// Get retrieves campaign information
	Get(ctx context.Context, req *CampaignGetRequest) (*CampaignGetResponse, error)

	// Update updates a campaign
	Update(ctx context.Context, req *CampaignUpdateRequest) (*CampaignUpdateResponse, error)

	// Delete deletes campaigns
	Delete(ctx context.Context, req *CampaignDeleteRequest) (*CampaignDeleteResponse, error)

	// UpdateStatus updates campaign status
	UpdateStatus(ctx context.Context, req *CampaignStatusUpdateRequest) (*CampaignStatusUpdateResponse, error)
}

// AdService defines the interface for ad-related operations
type AdService interface {
	// Create creates new ads
	Create(ctx context.Context, req *AdCreateRequest) (*AdCreateResponse, error)

	// Get retrieves ad information
	Get(ctx context.Context, req *AdGetRequest) (*AdGetResponse, error)

	// Update updates ads
	Update(ctx context.Context, req *AdUpdateRequest) (*AdUpdateResponse, error)

	// Delete deletes ads
	Delete(ctx context.Context, req *AdDeleteRequest) (*AdDeleteResponse, error)

	// UpdateStatus updates ad status
	UpdateStatus(ctx context.Context, req *AdStatusUpdateRequest) (*AdStatusUpdateResponse, error)
}

// AdGroupService defines the interface for ad group-related operations
type AdGroupService interface {
	// Create creates new ad groups
	Create(ctx context.Context, req *AdGroupCreateRequest) (*AdGroupCreateResponse, error)

	// Get retrieves ad group information
	Get(ctx context.Context, req *AdGroupGetRequest) (*AdGroupGetResponse, error)

	// Update updates ad groups
	Update(ctx context.Context, req *AdGroupUpdateRequest) (*AdGroupUpdateResponse, error)

	// Delete deletes ad groups
	Delete(ctx context.Context, req *AdGroupDeleteRequest) (*AdGroupDeleteResponse, error)

	// UpdateStatus updates ad group status
	UpdateStatus(ctx context.Context, req *AdGroupStatusUpdateRequest) (*AdGroupStatusUpdateResponse, error)
}

// AudienceService defines the interface for audience-related operations
type AudienceService interface {
	// CreateCustomAudience creates a custom audience
	CreateCustomAudience(ctx context.Context, req *CustomAudienceCreateRequest) (*CustomAudienceCreateResponse, error)

	// GetCustomAudiences retrieves custom audiences
	GetCustomAudiences(ctx context.Context, req *CustomAudienceGetRequest) (*CustomAudienceGetResponse, error)

	// UpdateCustomAudience updates a custom audience
	UpdateCustomAudience(ctx context.Context, req *CustomAudienceUpdateRequest) (*CustomAudienceUpdateResponse, error)

	// DeleteCustomAudience deletes a custom audience
	DeleteCustomAudience(ctx context.Context, req *CustomAudienceDeleteRequest) (*CustomAudienceDeleteResponse, error)

	// CreateLookalikeAudience creates a lookalike audience
	CreateLookalikeAudience(ctx context.Context, req *LookalikeAudienceCreateRequest) (*LookalikeAudienceCreateResponse, error)
}

// CreativeService defines the interface for creative-related operations
type CreativeService interface {
	// UploadImage uploads an image creative
	UploadImage(ctx context.Context, req *ImageUploadRequest) (*ImageUploadResponse, error)

	// UploadVideo uploads a video creative
	UploadVideo(ctx context.Context, req *VideoUploadRequest) (*VideoUploadResponse, error)

	// GetCreatives retrieves creative assets
	GetCreatives(ctx context.Context, req *CreativeGetRequest) (*CreativeGetResponse, error)

	// UpdateCreative updates creative information
	UpdateCreative(ctx context.Context, req *CreativeUpdateRequest) (*CreativeUpdateResponse, error)
}

// ReportingService defines the interface for reporting operations
type ReportingService interface {
	// GetBasicReports retrieves basic performance reports
	GetBasicReports(ctx context.Context, req *ReportingRequest) (*ReportingResponse, error)

	// GetAudienceReports retrieves audience reports
	GetAudienceReports(ctx context.Context, req *AudienceReportingRequest) (*AudienceReportingResponse, error)

	// CreateAsyncReport creates an asynchronous report
	CreateAsyncReport(ctx context.Context, req *AsyncReportRequest) (*AsyncReportResponse, error)

	// GetAsyncReportStatus checks the status of an asynchronous report
	GetAsyncReportStatus(ctx context.Context, taskID string) (*AsyncReportStatusResponse, error)

	// DownloadAsyncReport downloads the results of an asynchronous report
	DownloadAsyncReport(ctx context.Context, taskID string) (io.ReadCloser, error)
}

// ToolService defines the interface for tool-related operations
type ToolService interface {
	// GetLanguages retrieves supported languages
	GetLanguages(ctx context.Context, advertiserID string) (*LanguagesResponse, error)

	// GetCurrencies retrieves supported currencies
	GetCurrencies(ctx context.Context, advertiserID string) (*CurrenciesResponse, error)

	// GetRegions retrieves supported regions
	GetRegions(ctx context.Context, advertiserID string) (*RegionsResponse, error)

	// GetInterestCategories retrieves interest categories for targeting
	GetInterestCategories(ctx context.Context, req *InterestCategoriesRequest) (*InterestCategoriesResponse, error)

	// GetCarriers retrieves mobile carriers for targeting
	GetCarriers(ctx context.Context, req *CarriersRequest) (*CarriersResponse, error)

	// GetDeviceModels retrieves device models for targeting
	GetDeviceModels(ctx context.Context, req *DeviceModelsRequest) (*DeviceModelsResponse, error)
}

// BCService defines the interface for Business Center operations
type BCService interface {
	// GetBusinessCenters retrieves business centers
	GetBusinessCenters(ctx context.Context) (*BusinessCentersResponse, error)

	// GetBusinessCenterInfo retrieves business center information
	GetBusinessCenterInfo(ctx context.Context, bcID string) (*BusinessCenterInfo, error)

	// GetAdvertisersInBC retrieves advertisers in a business center
	GetAdvertisersInBC(ctx context.Context, bcID string) (*BCAdvertisersResponse, error)

	// TransferAdvertiser transfers an advertiser to a business center
	TransferAdvertiser(ctx context.Context, req *TransferAdvertiserRequest) (*TransferAdvertiserResponse, error)
}

// AuthService defines the interface for authentication operations
type AuthService interface {
	// GetAuthorizationURL generates an OAuth authorization URL
	GetAuthorizationURL(scopes []string) string

	// GetAccessToken exchanges an authorization code for an access token
	GetAccessToken(ctx context.Context, code string) (*TokenResponse, error)

	// RefreshToken refreshes an access token using a refresh token
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error)

	// ValidateToken validates an access token
	ValidateToken(ctx context.Context, token string) (*TokenValidationResponse, error)

	// RevokeToken revokes an access token
	RevokeToken(ctx context.Context, token string) error
}

// OptimizerService defines the interface for optimizer-related operations
type OptimizerService interface {
	// CreateRule creates a new automated rule
	CreateRule(ctx context.Context, req *OptimizerRuleCreateRequest) (*OptimizerRuleResponse, error)

	// GetRule retrieves optimizer rule information
	GetRule(ctx context.Context, req *OptimizerRuleGetRequest) (*OptimizerRuleResponse, error)

	// ListRules retrieves a list of optimizer rules
	ListRules(ctx context.Context, req *OptimizerRuleListRequest) (*OptimizerRuleListResponse, error)

	// UpdateRule updates an existing optimizer rule
	UpdateRule(ctx context.Context, req *OptimizerRuleUpdateRequest) (*OptimizerRuleResponse, error)

	// BatchBindRule binds rules to campaigns/ad groups in batch
	BatchBindRule(ctx context.Context, req *OptimizerRuleBatchBindRequest) (*OptimizerRuleBatchBindResponse, error)

	// GetRuleResult retrieves optimizer rule execution results
	GetRuleResult(ctx context.Context, req *OptimizerRuleResultGetRequest) (*OptimizerRuleResultResponse, error)

	// ListRuleResults retrieves a list of optimizer rule execution results
	ListRuleResults(ctx context.Context, req *OptimizerRuleResultListRequest) (*OptimizerRuleResultListResponse, error)
}

// CommentService defines the interface for comment-related operations
type CommentService interface {
	// ListComments retrieves a list of comments
	ListComments(ctx context.Context, req *CommentListRequest) (*CommentListResponse, error)

	// PostComment posts a new comment
	PostComment(ctx context.Context, req *CommentPostRequest) (*CommentPostResponse, error)

	// DeleteComment deletes a comment
	DeleteComment(ctx context.Context, req *CommentDeleteRequest) (*CommentDeleteResponse, error)

	// UpdateCommentStatus updates comment status
	UpdateCommentStatus(ctx context.Context, req *CommentStatusUpdateRequest) (*CommentStatusUpdateResponse, error)

	// GetCommentReference retrieves comment reference information
	GetCommentReference(ctx context.Context, req *CommentReferenceRequest) (*CommentReferenceResponse, error)

	// CreateCommentTask creates a comment management task
	CreateCommentTask(ctx context.Context, req *CommentTaskCreateRequest) (*CommentTaskResponse, error)

	// CheckCommentTask checks comment task status
	CheckCommentTask(ctx context.Context, req *CommentTaskCheckRequest) (*CommentTaskCheckResponse, error)
}

// ReportService defines the interface for report-related operations
type ReportService interface {
	// GetIntegratedReport retrieves integrated reporting data
	GetIntegratedReport(ctx context.Context, req *ReportIntegratedGetRequest) (*ReportIntegratedResponse, error)

	// CreateReportTask creates a new report generation task
	CreateReportTask(ctx context.Context, req *ReportTaskCreateRequest) (*ReportTaskResponse, error)

	// CheckReportTask checks the status of a report generation task
	CheckReportTask(ctx context.Context, req *ReportTaskCheckRequest) (*ReportTaskCheckResponse, error)

	// CancelReportTask cancels a report generation task
	CancelReportTask(ctx context.Context, req *ReportTaskCancelRequest) (*ReportTaskCancelResponse, error)
}
