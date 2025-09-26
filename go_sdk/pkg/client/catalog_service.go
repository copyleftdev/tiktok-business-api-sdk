package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// CatalogService handles Catalog related operations
type CatalogService struct {
	client *Client
}

// NewCatalogService creates a new CatalogService
func NewCatalogService(client *Client) *CatalogService {
	return &CatalogService{client: client}
}

// CatalogCreateRequest represents the request for creating a catalog
type CatalogCreateRequest struct {
	AdvertiserID   string `json:"advertiser_id"`
	CatalogName    string `json:"catalog_name"`
	CatalogType    string `json:"catalog_type"` // PRODUCT, HOTEL, FLIGHT, etc.
	Description    string `json:"description,omitempty"`
	DefaultLocale  string `json:"default_locale,omitempty"`
	DefaultCurrency string `json:"default_currency,omitempty"`
}

// CatalogGetRequest represents the request for getting catalog information
type CatalogGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id,omitempty"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

// CatalogUpdateRequest represents the request for updating a catalog
type CatalogUpdateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	CatalogName  string `json:"catalog_name,omitempty"`
	Description  string `json:"description,omitempty"`
}

// CatalogDeleteRequest represents the request for deleting a catalog
type CatalogDeleteRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
}

// CatalogData represents catalog information
type CatalogData struct {
	CatalogID       string `json:"catalog_id"`
	CatalogName     string `json:"catalog_name"`
	CatalogType     string `json:"catalog_type"`
	Description     string `json:"description"`
	DefaultLocale   string `json:"default_locale"`
	DefaultCurrency string `json:"default_currency"`
	Status          string `json:"status"`
	ProductCount    int    `json:"product_count"`
	CreateTime      string `json:"create_time"`
	UpdateTime      string `json:"update_time"`
}

// CatalogResponse represents the response from catalog operations
type CatalogResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      CatalogData `json:"data"`
}

// CatalogListResponse represents the response for listing catalogs
type CatalogListResponse struct {
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	RequestID string        `json:"request_id"`
	Data      []CatalogData `json:"data"`
}

// Create creates a new catalog
func (s *CatalogService) Create(ctx context.Context, req *CatalogCreateRequest) (*CatalogResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogName == "" {
		return nil, fmt.Errorf("catalog_name is required")
	}
	if req.CatalogType == "" {
		return nil, fmt.Errorf("catalog_type is required")
	}

	url := s.client.BuildURL("/catalog/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create catalog: %w", err)
	}

	var response CatalogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Get retrieves catalog information
func (s *CatalogService) Get(ctx context.Context, req *CatalogGetRequest) (*CatalogListResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.CatalogID != "" {
		params["catalog_id"] = req.CatalogID
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/catalog/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog: %w", err)
	}

	var response CatalogListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Update updates an existing catalog
func (s *CatalogService) Update(ctx context.Context, req *CatalogUpdateRequest) (*CatalogResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}

	url := s.client.BuildURL("/catalog/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update catalog: %w", err)
	}

	var response CatalogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Delete deletes a catalog
func (s *CatalogService) Delete(ctx context.Context, req *CatalogDeleteRequest) (*CatalogResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}

	url := s.client.BuildURL("/catalog/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete catalog: %w", err)
	}

	var response CatalogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CatalogOverviewRequest represents the request for catalog overview
type CatalogOverviewRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
}

// CatalogOverviewData represents catalog overview information
type CatalogOverviewData struct {
	CatalogID       string                 `json:"catalog_id"`
	CatalogName     string                 `json:"catalog_name"`
	ProductCount    int                    `json:"product_count"`
	ActiveProducts  int                    `json:"active_products"`
	PendingProducts int                    `json:"pending_products"`
	RejectedProducts int                   `json:"rejected_products"`
	LastSyncTime    string                 `json:"last_sync_time"`
	Statistics      map[string]interface{} `json:"statistics"`
}

// CatalogOverviewResponse represents the response for catalog overview
type CatalogOverviewResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      CatalogOverviewData `json:"data"`
}

// GetOverview retrieves catalog overview information
func (s *CatalogService) GetOverview(ctx context.Context, req *CatalogOverviewRequest) (*CatalogOverviewResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CatalogID == "" {
		return nil, fmt.Errorf("advertiser_id and catalog_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"catalog_id":    req.CatalogID,
	}

	url := s.client.BuildURL("/catalog/overview/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog overview: %w", err)
	}

	var response CatalogOverviewResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateFeed creates a new catalog feed
func (s *CatalogService) CreateFeed(ctx context.Context, req *CatalogFeedCreateRequest) (*CatalogFeedResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}
	if req.FeedURL == "" {
		return nil, fmt.Errorf("feed_url is required")
	}

	url := s.client.BuildURL("/catalog/feed/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create catalog feed: %w", err)
	}

	var response CatalogFeedResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetFeed retrieves catalog feed information
func (s *CatalogService) GetFeed(ctx context.Context, req *CatalogFeedGetRequest) (*CatalogFeedListResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CatalogID == "" {
		return nil, fmt.Errorf("advertiser_id and catalog_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"catalog_id":    req.CatalogID,
	}

	if req.FeedID != "" {
		params["feed_id"] = req.FeedID
	}

	url := s.client.BuildURL("/catalog/feed/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog feed: %w", err)
	}

	var response CatalogFeedListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateFeed updates a catalog feed
func (s *CatalogService) UpdateFeed(ctx context.Context, req *CatalogFeedUpdateRequest) (*CatalogFeedResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}
	if req.FeedID == "" {
		return nil, fmt.Errorf("feed_id is required")
	}

	url := s.client.BuildURL("/catalog/feed/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update catalog feed: %w", err)
	}

	var response CatalogFeedResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteFeed deletes a catalog feed
func (s *CatalogService) DeleteFeed(ctx context.Context, req *CatalogFeedDeleteRequest) (*CatalogFeedResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}
	if req.FeedID == "" {
		return nil, fmt.Errorf("feed_id is required")
	}

	url := s.client.BuildURL("/catalog/feed/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete catalog feed: %w", err)
	}

	var response CatalogFeedResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetFeedLog retrieves catalog feed processing logs
func (s *CatalogService) GetFeedLog(ctx context.Context, req *CatalogFeedLogRequest) (*CatalogFeedLogResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CatalogID == "" || req.FeedID == "" {
		return nil, fmt.Errorf("advertiser_id, catalog_id, and feed_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"catalog_id":    req.CatalogID,
		"feed_id":       req.FeedID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/catalog/feed/log/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog feed log: %w", err)
	}

	var response CatalogFeedLogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteProduct deletes products from catalog
func (s *CatalogService) DeleteProduct(ctx context.Context, req *CatalogProductDeleteRequest) (*CatalogProductDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CatalogID == "" {
		return nil, fmt.Errorf("catalog_id is required")
	}
	if len(req.ProductIDs) == 0 {
		return nil, fmt.Errorf("product_ids is required")
	}

	url := s.client.BuildURL("/catalog/product/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete catalog products: %w", err)
	}

	var response CatalogProductDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetProductFile retrieves product file information
func (s *CatalogService) GetProductFile(ctx context.Context, req *CatalogProductFileRequest) (*CatalogProductFileResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CatalogID == "" {
		return nil, fmt.Errorf("advertiser_id and catalog_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"catalog_id":    req.CatalogID,
	}

	if req.FileType != "" {
		params["file_type"] = req.FileType
	}

	url := s.client.BuildURL("/catalog/product/file/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog product file: %w", err)
	}

	var response CatalogProductFileResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetProductLog retrieves product processing logs
func (s *CatalogService) GetProductLog(ctx context.Context, req *CatalogProductLogRequest) (*CatalogProductLogResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CatalogID == "" {
		return nil, fmt.Errorf("advertiser_id and catalog_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"catalog_id":    req.CatalogID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/catalog/product/log/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog product log: %w", err)
	}

	var response CatalogProductLogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional Catalog types for new methods
type CatalogFeedCreateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FeedURL      string `json:"feed_url"`
	FeedName     string `json:"feed_name,omitempty"`
	Schedule     string `json:"schedule,omitempty"` // DAILY, WEEKLY, MONTHLY
}

type CatalogFeedResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      CatalogFeedData   `json:"data"`
}

type CatalogFeedData struct {
	FeedID     string `json:"feed_id"`
	FeedName   string `json:"feed_name"`
	FeedURL    string `json:"feed_url"`
	Status     string `json:"status"`
	Schedule   string `json:"schedule"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type CatalogFeedGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FeedID       string `json:"feed_id,omitempty"`
}

type CatalogFeedListResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      CatalogFeedListData   `json:"data"`
}

type CatalogFeedListData struct {
	Feeds []CatalogFeedData `json:"feeds"`
}

type CatalogFeedUpdateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FeedID       string `json:"feed_id"`
	FeedName     string `json:"feed_name,omitempty"`
	FeedURL      string `json:"feed_url,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}

type CatalogFeedDeleteRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FeedID       string `json:"feed_id"`
}

type CatalogFeedLogRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FeedID       string `json:"feed_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

type CatalogFeedLogResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      CatalogFeedLogData   `json:"data"`
}

type CatalogFeedLogData struct {
	Logs []FeedLogEntry `json:"logs"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type FeedLogEntry struct {
	LogID       string `json:"log_id"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ProcessTime string `json:"process_time"`
	ProductsProcessed int `json:"products_processed"`
	ProductsSuccess   int `json:"products_success"`
	ProductsError     int `json:"products_error"`
}

type CatalogProductDeleteRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	CatalogID    string   `json:"catalog_id"`
	ProductIDs   []string `json:"product_ids"`
}

type CatalogProductDeleteResponse struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"request_id"`
	Data      CatalogProductDeleteData   `json:"data"`
}

type CatalogProductDeleteData struct {
	Results []ProductDeleteResult `json:"results"`
}

type ProductDeleteResult struct {
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
	Message   string `json:"message,omitempty"`
}

type CatalogProductFileRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	FileType     string `json:"file_type,omitempty"` // CSV, TSV, XML
}

type CatalogProductFileResponse struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	RequestID string                   `json:"request_id"`
	Data      CatalogProductFileData   `json:"data"`
}

type CatalogProductFileData struct {
	FileURL    string `json:"file_url"`
	FileType   string `json:"file_type"`
	FileSize   int64  `json:"file_size"`
	CreateTime string `json:"create_time"`
	ExpiryTime string `json:"expiry_time"`
}

type CatalogProductLogRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CatalogID    string `json:"catalog_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

type CatalogProductLogResponse struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	RequestID string                  `json:"request_id"`
	Data      CatalogProductLogData   `json:"data"`
}

type CatalogProductLogData struct {
	Logs []ProductLogEntry `json:"logs"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type ProductLogEntry struct {
	LogID       string `json:"log_id"`
	ProductID   string `json:"product_id"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ProcessTime string `json:"process_time"`
}
