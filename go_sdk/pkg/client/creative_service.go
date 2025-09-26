package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// creativeService handles Creative related operations
type creativeService struct {
	client *Client
}

// NewCreativeService creates a new CreativeService
func NewCreativeService(client *Client) CreativeService {
	return &creativeService{client: client}
}

// CreativePortfolioCreateRequest represents the request for creating a creative portfolio
type CreativePortfolioCreateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
}

// CreativePortfolioResponse represents the response for creative portfolio operations
type CreativePortfolioResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      CreativePortfolioData `json:"data"`
}

type CreativePortfolioData struct {
	PortfolioID string `json:"portfolio_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// CreativePortfolioGetRequest represents the request for getting a creative portfolio
type CreativePortfolioGetRequest struct {
	AdvertiserID        string `json:"advertiser_id"`
	CreativePortfolioID string `json:"creative_portfolio_id"`
}

// CreativePortfolioListRequest represents the request for listing creative portfolios
type CreativePortfolioListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

// CreativePortfolioListResponse represents the response for listing creative portfolios
type CreativePortfolioListResponse struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"request_id"`
	Data      CreativePortfolioListData  `json:"data"`
}

type CreativePortfolioListData struct {
	Portfolios []CreativePortfolioData `json:"portfolios"`
	PageInfo   struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
		TotalPage  int `json:"total_page"`
	} `json:"page_info"`
}

// CreativeAssetShareRequest represents the request for sharing creative assets
type CreativeAssetShareRequest struct {
	AdvertiserID       string   `json:"advertiser_id"`
	AssetIDs           []string `json:"asset_ids"`
	TargetAdvertiserID string   `json:"target_advertiser_id"`
	AssetType          string   `json:"asset_type"` // IMAGE, VIDEO, AUDIO
}

// CreativeAssetShareResponse represents the response for sharing creative assets
type CreativeAssetShareResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      CreativeAssetShareData `json:"data"`
}

type CreativeAssetShareData struct {
	SharedAssets []SharedAssetInfo `json:"shared_assets"`
}

type SharedAssetInfo struct {
	AssetID        string `json:"asset_id"`
	SharedAssetID  string `json:"shared_asset_id"`
	Status         string `json:"status"`
	FailureReason  string `json:"failure_reason,omitempty"`
}

// CreativeAssetDeleteRequest represents the request for deleting creative assets
type CreativeAssetDeleteRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	AssetIDs     []string `json:"asset_ids"`
	AssetType    string   `json:"asset_type"`
}

// CreativeAssetDeleteResponse represents the response for deleting creative assets
type CreativeAssetDeleteResponse struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	RequestID string                  `json:"request_id"`
	Data      CreativeAssetDeleteData `json:"data"`
}

type CreativeAssetDeleteData struct {
	DeletedAssets []DeletedAssetInfo `json:"deleted_assets"`
}

type DeletedAssetInfo struct {
	AssetID       string `json:"asset_id"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason,omitempty"`
}

// CreativeImageEditRequest represents the request for editing creative images
type CreativeImageEditRequest struct {
	AdvertiserID string                 `json:"advertiser_id"`
	ImageID      string                 `json:"image_id"`
	Operations   []ImageEditOperation   `json:"operations"`
}

type ImageEditOperation struct {
	Type       string                 `json:"type"` // CROP, RESIZE, FILTER
	Parameters map[string]interface{} `json:"parameters"`
}

// CreativeImageEditResponse represents the response for editing creative images
type CreativeImageEditResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      CreativeImageEditData `json:"data"`
}

type CreativeImageEditData struct {
	EditedImageID string `json:"edited_image_id"`
	ImageURL      string `json:"image_url"`
	Status        string `json:"status"`
}

// CreativeSmartTextGenerateRequest represents the request for generating smart text
type CreativeSmartTextGenerateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Prompt       string `json:"prompt"`
	Language     string `json:"language,omitempty"`
	MaxLength    int    `json:"max_length,omitempty"`
	Style        string `json:"style,omitempty"` // CASUAL, FORMAL, CREATIVE
}

// CreativeSmartTextGenerateResponse represents the response for generating smart text
type CreativeSmartTextGenerateResponse struct {
	Code      int                           `json:"code"`
	Message   string                        `json:"message"`
	RequestID string                        `json:"request_id"`
	Data      CreativeSmartTextGenerateData `json:"data"`
}

type CreativeSmartTextGenerateData struct {
	GeneratedTexts []GeneratedTextInfo `json:"generated_texts"`
}

type GeneratedTextInfo struct {
	Text      string  `json:"text"`
	Score     float64 `json:"score"`
	Language  string  `json:"language"`
}

// CreativeShareableLinkCreateRequest represents the request for creating shareable links
type CreativeShareableLinkCreateRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	AssetIDs     []string `json:"asset_ids"`
	AssetType    string   `json:"asset_type"`
	ExpiryDays   int      `json:"expiry_days,omitempty"`
}

// CreativeShareableLinkCreateResponse represents the response for creating shareable links
type CreativeShareableLinkCreateResponse struct {
	Code      int                             `json:"code"`
	Message   string                          `json:"message"`
	RequestID string                          `json:"request_id"`
	Data      CreativeShareableLinkCreateData `json:"data"`
}

type CreativeShareableLinkCreateData struct {
	ShareableLinks []ShareableLinkInfo `json:"shareable_links"`
}

type ShareableLinkInfo struct {
	AssetID     string `json:"asset_id"`
	ShareURL    string `json:"share_url"`
	ExpiryTime  string `json:"expiry_time"`
}

// UploadImage uploads an image creative (interface method)
func (s *creativeService) UploadImage(ctx context.Context, req *ImageUploadRequest) (*ImageUploadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	url := s.client.BuildURL("/creative/image/upload/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	var response ImageUploadResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UploadVideo uploads a video creative (interface method)
func (s *creativeService) UploadVideo(ctx context.Context, req *VideoUploadRequest) (*VideoUploadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	url := s.client.BuildURL("/creative/video/upload/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upload video: %w", err)
	}

	var response VideoUploadResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCreatives retrieves creative assets (interface method)
func (s *creativeService) GetCreatives(ctx context.Context, req *CreativeGetRequest) (*CreativeGetResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.CreativeType != "" {
		params["creative_type"] = req.CreativeType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/creative/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get creatives: %w", err)
	}

	var response CreativeGetResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCreative updates creative information (interface method)
func (s *creativeService) UpdateCreative(ctx context.Context, req *CreativeUpdateRequest) (*CreativeUpdateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CreativeID == "" {
		return nil, fmt.Errorf("creative_id is required")
	}

	url := s.client.BuildURL("/creative/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update creative: %w", err)
	}

	var response CreativeUpdateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreatePortfolio creates a new creative portfolio
func (s *creativeService) CreatePortfolio(ctx context.Context, req *CreativePortfolioCreateRequest) (*CreativePortfolioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	url := s.client.BuildURL("/creative/portfolio/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create creative portfolio: %w", err)
	}

	var response CreativePortfolioResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPortfolio retrieves a creative portfolio by ID
func (s *creativeService) GetPortfolio(ctx context.Context, req *CreativePortfolioGetRequest) (*CreativePortfolioResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CreativePortfolioID == "" {
		return nil, fmt.Errorf("advertiser_id and creative_portfolio_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id":         req.AdvertiserID,
		"creative_portfolio_id": req.CreativePortfolioID,
	}

	url := s.client.BuildURL("/creative/portfolio/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get creative portfolio: %w", err)
	}

	var response CreativePortfolioResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListPortfolios retrieves a list of creative portfolios
func (s *creativeService) ListPortfolios(ctx context.Context, req *CreativePortfolioListRequest) (*CreativePortfolioListResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/creative/portfolio/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list creative portfolios: %w", err)
	}

	var response CreativePortfolioListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ShareAssets shares creative assets with another advertiser
func (s *creativeService) ShareAssets(ctx context.Context, req *CreativeAssetShareRequest) (*CreativeAssetShareResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if len(req.AssetIDs) == 0 {
		return nil, fmt.Errorf("asset_ids is required")
	}
	if req.TargetAdvertiserID == "" {
		return nil, fmt.Errorf("target_advertiser_id is required")
	}

	url := s.client.BuildURL("/creative/asset/share/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to share creative assets: %w", err)
	}

	var response CreativeAssetShareResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteAssets deletes creative assets
func (s *creativeService) DeleteAssets(ctx context.Context, req *CreativeAssetDeleteRequest) (*CreativeAssetDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if len(req.AssetIDs) == 0 {
		return nil, fmt.Errorf("asset_ids is required")
	}

	url := s.client.BuildURL("/creative/asset/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete creative assets: %w", err)
	}

	var response CreativeAssetDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// EditImage edits a creative image
func (s *creativeService) EditImage(ctx context.Context, req *CreativeImageEditRequest) (*CreativeImageEditResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.ImageID == "" {
		return nil, fmt.Errorf("image_id is required")
	}

	url := s.client.BuildURL("/creative/image/edit/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to edit creative image: %w", err)
	}

	var response CreativeImageEditResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GenerateSmartText generates smart text using AI
func (s *creativeService) GenerateSmartText(ctx context.Context, req *CreativeSmartTextGenerateRequest) (*CreativeSmartTextGenerateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.Prompt == "" {
		return nil, fmt.Errorf("prompt is required")
	}

	url := s.client.BuildURL("/creative/smart_text/generate/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate smart text: %w", err)
	}

	var response CreativeSmartTextGenerateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateShareableLink creates shareable links for creative assets
func (s *creativeService) CreateShareableLink(ctx context.Context, req *CreativeShareableLinkCreateRequest) (*CreativeShareableLinkCreateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if len(req.AssetIDs) == 0 {
		return nil, fmt.Errorf("asset_ids is required")
	}

	url := s.client.BuildURL("/creative/shareable_link/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create shareable link: %w", err)
	}

	var response CreativeShareableLinkCreateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
