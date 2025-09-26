package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// DMPService handles Data Management Platform operations (Custom Audiences)
type DMPService struct {
	client *Client
}

// NewDMPService creates a new DMPService
func NewDMPService(client *Client) *DMPService {
	return &DMPService{client: client}
}

// CustomAudienceCreateRequest represents the request for creating a custom audience
type CustomAudienceCreateRequest struct {
	AdvertiserID     string                 `json:"advertiser_id"`
	AudienceName     string                 `json:"audience_name"`
	AudienceType     string                 `json:"audience_type"` // CUSTOMER_FILE, WEBSITE_TRAFFIC, APP_ACTIVITY, etc.
	Description      string                 `json:"description,omitempty"`
	RetentionDays    int                    `json:"retention_days,omitempty"`
	ShareToBC        bool                   `json:"share_to_bc,omitempty"`
	Rules            []AudienceRule         `json:"rules,omitempty"`
	FileUploadConfig *FileUploadConfig      `json:"file_upload_config,omitempty"`
	PixelConfig      *PixelConfig           `json:"pixel_config,omitempty"`
	AppConfig        *AppConfig             `json:"app_config,omitempty"`
	CustomData       map[string]interface{} `json:"custom_data,omitempty"`
}

// AudienceRule represents rules for custom audience creation
type AudienceRule struct {
	RuleType   string                 `json:"rule_type"`
	Conditions []AudienceCondition    `json:"conditions"`
	Operator   string                 `json:"operator"` // AND, OR
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

// AudienceCondition represents conditions within audience rules
type AudienceCondition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"` // EQUALS, CONTAINS, GREATER_THAN, etc.
	Value    interface{} `json:"value"`
}

// FileUploadConfig represents configuration for file-based audiences
type FileUploadConfig struct {
	FileType     string   `json:"file_type"` // CSV, TXT
	FileURL      string   `json:"file_url,omitempty"`
	FileContent  string   `json:"file_content,omitempty"`
	IdentifierType string `json:"identifier_type"` // EMAIL, PHONE, IDFA, etc.
	HasHeader    bool     `json:"has_header"`
	Delimiter    string   `json:"delimiter,omitempty"`
	Columns      []string `json:"columns,omitempty"`
}

// PixelConfig represents configuration for pixel-based audiences
type PixelConfig struct {
	PixelID     string   `json:"pixel_id"`
	EventNames  []string `json:"event_names,omitempty"`
	URLContains string   `json:"url_contains,omitempty"`
	URLEquals   string   `json:"url_equals,omitempty"`
}

// AppConfig represents configuration for app-based audiences
type AppConfig struct {
	AppID       string   `json:"app_id"`
	EventNames  []string `json:"event_names,omitempty"`
	EventParams []string `json:"event_params,omitempty"`
}

// CustomAudienceGetRequest represents the request for getting custom audience information
type CustomAudienceGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	AudienceID   string `json:"audience_id,omitempty"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

// CustomAudienceUpdateRequest represents the request for updating a custom audience
type CustomAudienceUpdateRequest struct {
	AdvertiserID  string `json:"advertiser_id"`
	AudienceID    string `json:"audience_id"`
	AudienceName  string `json:"audience_name,omitempty"`
	Description   string `json:"description,omitempty"`
	RetentionDays int    `json:"retention_days,omitempty"`
}

// CustomAudienceDeleteRequest represents the request for deleting a custom audience
type CustomAudienceDeleteRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	AudienceID   string `json:"audience_id"`
}

// CustomAudienceData represents custom audience information
type CustomAudienceData struct {
	AudienceID      string                 `json:"audience_id"`
	AudienceName    string                 `json:"audience_name"`
	AudienceType    string                 `json:"audience_type"`
	Description     string                 `json:"description"`
	Size            int64                  `json:"size"`
	Status          string                 `json:"status"`
	RetentionDays   int                    `json:"retention_days"`
	ShareToBC       bool                   `json:"share_to_bc"`
	CreateTime      string                 `json:"create_time"`
	UpdateTime      string                 `json:"update_time"`
	LastSyncTime    string                 `json:"last_sync_time"`
	Rules           []AudienceRule         `json:"rules"`
	CustomData      map[string]interface{} `json:"custom_data"`
}

// CustomAudienceResponse represents the response from custom audience operations
type CustomAudienceResponse struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	RequestID string             `json:"request_id"`
	Data      CustomAudienceData `json:"data"`
}

// CustomAudienceListResponse represents the response for listing custom audiences
type CustomAudienceListResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      []CustomAudienceData `json:"data"`
}

// Aliases for interface compatibility
type CustomAudienceCreateResponse = CustomAudienceResponse
type CustomAudienceGetResponse = CustomAudienceListResponse
type CustomAudienceUpdateResponse = CustomAudienceResponse
type CustomAudienceDeleteResponse = CustomAudienceResponse
type LookalikeAudienceCreateResponse = CustomAudienceResponse

// CreateCustomAudience creates a new custom audience
func (s *DMPService) CreateCustomAudience(ctx context.Context, req *CustomAudienceCreateRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceName == "" {
		return nil, fmt.Errorf("audience_name is required")
	}
	if req.AudienceType == "" {
		return nil, fmt.Errorf("audience_type is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom audience: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomAudience retrieves custom audience information
func (s *DMPService) GetCustomAudience(ctx context.Context, req *CustomAudienceGetRequest) (*CustomAudienceListResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.AudienceID != "" {
		params["audience_id"] = req.AudienceID
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/dmp/custom_audience/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom audience: %w", err)
	}

	var response CustomAudienceListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCustomAudience updates an existing custom audience
func (s *DMPService) UpdateCustomAudience(ctx context.Context, req *CustomAudienceUpdateRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceID == "" {
		return nil, fmt.Errorf("audience_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update custom audience: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCustomAudience deletes a custom audience
func (s *DMPService) DeleteCustomAudience(ctx context.Context, req *CustomAudienceDeleteRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceID == "" {
		return nil, fmt.Errorf("audience_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete custom audience: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// LookalikeAudienceCreateRequest represents the request for creating a lookalike audience
type LookalikeAudienceCreateRequest struct {
	AdvertiserID     string  `json:"advertiser_id"`
	AudienceName     string  `json:"audience_name"`
	SourceAudienceID string  `json:"source_audience_id"`
	SimilarityLevel  float64 `json:"similarity_level"` // 1-10, higher = more similar
	CountryCode      string  `json:"country_code"`
	Description      string  `json:"description,omitempty"`
	ShareToBC        bool    `json:"share_to_bc,omitempty"`
}

// CreateLookalikeAudience creates a new lookalike audience
func (s *DMPService) CreateLookalikeAudience(ctx context.Context, req *LookalikeAudienceCreateRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceName == "" {
		return nil, fmt.Errorf("audience_name is required")
	}
	if req.SourceAudienceID == "" {
		return nil, fmt.Errorf("source_audience_id is required")
	}
	if req.CountryCode == "" {
		return nil, fmt.Errorf("country_code is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/lookalike/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create lookalike audience: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListCustomAudiences retrieves a list of custom audiences
func (s *DMPService) ListCustomAudiences(ctx context.Context, req *CustomAudienceListRequest) (*CustomAudienceListResponse, error) {
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
	if req.AudienceType != "" {
		params["audience_type"] = req.AudienceType
	}

	url := s.client.BuildURL("/dmp/custom_audience/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list custom audiences: %w", err)
	}

	var response CustomAudienceListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UploadCustomAudienceFile uploads a file for custom audience creation
func (s *DMPService) UploadCustomAudienceFile(ctx context.Context, req *CustomAudienceFileUploadRequest) (*CustomAudienceFileUploadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if len(req.FileData) == 0 {
		return nil, fmt.Errorf("file_data is required")
	}
	if req.FileType == "" {
		return nil, fmt.Errorf("file_type is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/file/upload/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upload custom audience file: %w", err)
	}

	var response CustomAudienceFileUploadResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ApplyCustomAudience applies custom audience data
func (s *DMPService) ApplyCustomAudience(ctx context.Context, req *CustomAudienceApplyRequest) (*CustomAudienceApplyResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CustomAudienceID == "" {
		return nil, fmt.Errorf("custom_audience_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/apply/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to apply custom audience: %w", err)
	}

	var response CustomAudienceApplyResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ShareCustomAudience shares a custom audience with another advertiser
func (s *DMPService) ShareCustomAudience(ctx context.Context, req *CustomAudienceShareRequest) (*CustomAudienceShareResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CustomAudienceID == "" {
		return nil, fmt.Errorf("custom_audience_id is required")
	}
	if req.TargetAdvertiserID == "" {
		return nil, fmt.Errorf("target_advertiser_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/share/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to share custom audience: %w", err)
	}

	var response CustomAudienceShareResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateSavedAudience creates a new saved audience
func (s *DMPService) CreateSavedAudience(ctx context.Context, req *SavedAudienceCreateRequest) (*SavedAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceName == "" {
		return nil, fmt.Errorf("audience_name is required")
	}

	url := s.client.BuildURL("/dmp/saved_audience/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create saved audience: %w", err)
	}

	var response SavedAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListSavedAudiences retrieves a list of saved audiences
func (s *DMPService) ListSavedAudiences(ctx context.Context, req *SavedAudienceListRequest) (*SavedAudienceListResponse, error) {
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

	url := s.client.BuildURL("/dmp/saved_audience/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list saved audiences: %w", err)
	}

	var response SavedAudienceListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteSavedAudience deletes a saved audience
func (s *DMPService) DeleteSavedAudience(ctx context.Context, req *SavedAudienceDeleteRequest) (*SavedAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.SavedAudienceID == "" {
		return nil, fmt.Errorf("saved_audience_id is required")
	}

	url := s.client.BuildURL("/dmp/saved_audience/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete saved audience: %w", err)
	}

	var response SavedAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomAudienceApplyLog retrieves custom audience apply logs
func (s *DMPService) GetCustomAudienceApplyLog(ctx context.Context, req *CustomAudienceApplyLogRequest) (*CustomAudienceApplyLogResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.CustomAudienceID == "" {
		return nil, fmt.Errorf("advertiser_id and custom_audience_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id":       req.AdvertiserID,
		"custom_audience_id":  req.CustomAudienceID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/dmp/custom_audience/apply/log/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom audience apply log: %w", err)
	}

	var response CustomAudienceApplyLogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateLookalikeAudience updates a lookalike audience
func (s *DMPService) UpdateLookalikeAudience(ctx context.Context, req *LookalikeAudienceUpdateRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CustomAudienceID == "" {
		return nil, fmt.Errorf("custom_audience_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/lookalike/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update lookalike audience: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateCustomAudienceRule creates a rule-based custom audience
func (s *DMPService) CreateCustomAudienceRule(ctx context.Context, req *CustomAudienceRuleCreateRequest) (*CustomAudienceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.AudienceName == "" {
		return nil, fmt.Errorf("audience_name is required")
	}
	if len(req.Rules) == 0 {
		return nil, fmt.Errorf("rules is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/rule/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom audience rule: %w", err)
	}

	var response CustomAudienceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CancelCustomAudienceShare cancels a custom audience share
func (s *DMPService) CancelCustomAudienceShare(ctx context.Context, req *CustomAudienceShareCancelRequest) (*CustomAudienceShareCancelResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CustomAudienceID == "" {
		return nil, fmt.Errorf("custom_audience_id is required")
	}
	if req.TargetAdvertiserID == "" {
		return nil, fmt.Errorf("target_advertiser_id is required")
	}

	url := s.client.BuildURL("/dmp/custom_audience/share/cancel/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel custom audience share: %w", err)
	}

	var response CustomAudienceShareCancelResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomAudienceShareLog retrieves custom audience share logs
func (s *DMPService) GetCustomAudienceShareLog(ctx context.Context, req *CustomAudienceShareLogRequest) (*CustomAudienceShareLogResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.CustomAudienceID != "" {
		params["custom_audience_id"] = req.CustomAudienceID
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/dmp/custom_audience/share/log/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom audience share log: %w", err)
	}

	var response CustomAudienceShareLogResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional DMP types for new methods
type CustomAudienceListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
	AudienceType string `json:"audience_type,omitempty"` // CUSTOM, LOOKALIKE
}

type CustomAudienceFileUploadRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	FileData     []byte `json:"file_data"`
	FileType     string `json:"file_type"` // CSV, TXT
	FileName     string `json:"file_name"`
}

type CustomAudienceFileUploadResponse struct {
	Code      int                             `json:"code"`
	Message   string                          `json:"message"`
	RequestID string                          `json:"request_id"`
	Data      CustomAudienceFileUploadData    `json:"data"`
}

type CustomAudienceFileUploadData struct {
	FileID     string `json:"file_id"`
	Status     string `json:"status"`
	UploadTime string `json:"upload_time"`
}

type CustomAudienceApplyRequest struct {
	AdvertiserID       string   `json:"advertiser_id"`
	CustomAudienceID   string   `json:"custom_audience_id"`
	FileID             string   `json:"file_id,omitempty"`
	UserData           []string `json:"user_data,omitempty"`
	Operation          string   `json:"operation"` // ADD, REMOVE, REPLACE
}

type CustomAudienceApplyResponse struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"request_id"`
	Data      CustomAudienceApplyData    `json:"data"`
}

type CustomAudienceApplyData struct {
	CustomAudienceID string `json:"custom_audience_id"`
	Status           string `json:"status"`
	ApplyTime        string `json:"apply_time"`
}

type CustomAudienceShareRequest struct {
	AdvertiserID       string `json:"advertiser_id"`
	CustomAudienceID   string `json:"custom_audience_id"`
	TargetAdvertiserID string `json:"target_advertiser_id"`
	ShareType          string `json:"share_type,omitempty"` // READ_ONLY, FULL_ACCESS
}

type CustomAudienceShareResponse struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"request_id"`
	Data      CustomAudienceShareData    `json:"data"`
}

type CustomAudienceShareData struct {
	ShareID   string `json:"share_id"`
	Status    string `json:"status"`
	ShareTime string `json:"share_time"`
}

type SavedAudienceCreateRequest struct {
	AdvertiserID   string                 `json:"advertiser_id"`
	AudienceName   string                 `json:"audience_name"`
	Description    string                 `json:"description,omitempty"`
	TargetingSpec  map[string]interface{} `json:"targeting_spec"`
}

type SavedAudienceResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      SavedAudienceData   `json:"data"`
}

type SavedAudienceData struct {
	SavedAudienceID string `json:"saved_audience_id"`
	AudienceName    string `json:"audience_name"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	CreateTime      string `json:"create_time"`
	UpdateTime      string `json:"update_time"`
}

type SavedAudienceListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

type SavedAudienceListResponse struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	RequestID string                  `json:"request_id"`
	Data      SavedAudienceListData   `json:"data"`
}

type SavedAudienceListData struct {
	Audiences []SavedAudienceData `json:"audiences"`
	PageInfo  struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type SavedAudienceDeleteRequest struct {
	AdvertiserID    string `json:"advertiser_id"`
	SavedAudienceID string `json:"saved_audience_id"`
}

// Additional missing types for new DMP methods
type CustomAudienceApplyLogRequest struct {
	AdvertiserID       string `json:"advertiser_id"`
	CustomAudienceID   string `json:"custom_audience_id"`
	Page               int    `json:"page,omitempty"`
	Size               int    `json:"size,omitempty"`
}

type CustomAudienceApplyLogResponse struct {
	Code      int                           `json:"code"`
	Message   string                        `json:"message"`
	RequestID string                        `json:"request_id"`
	Data      CustomAudienceApplyLogData    `json:"data"`
}

type CustomAudienceApplyLogData struct {
	Logs     []CustomAudienceApplyLogEntry `json:"logs"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type CustomAudienceApplyLogEntry struct {
	LogID            string `json:"log_id"`
	Operation        string `json:"operation"`
	Status           string `json:"status"`
	ProcessedCount   int    `json:"processed_count"`
	SuccessCount     int    `json:"success_count"`
	FailedCount      int    `json:"failed_count"`
	ApplyTime        string `json:"apply_time"`
	CompleteTime     string `json:"complete_time,omitempty"`
	ErrorMessage     string `json:"error_message,omitempty"`
}

type LookalikeAudienceUpdateRequest struct {
	AdvertiserID       string                 `json:"advertiser_id"`
	CustomAudienceID   string                 `json:"custom_audience_id"`
	AudienceName       string                 `json:"audience_name,omitempty"`
	Description        string                 `json:"description,omitempty"`
	LookalikeSpec      map[string]interface{} `json:"lookalike_spec,omitempty"`
	TargetCountries    []string               `json:"target_countries,omitempty"`
	AudienceSize       string                 `json:"audience_size,omitempty"` // NARROW, BALANCED, BROAD
}

type CustomAudienceRuleCreateRequest struct {
	AdvertiserID   string                      `json:"advertiser_id"`
	AudienceName   string                      `json:"audience_name"`
	Description    string                      `json:"description,omitempty"`
	Rules          []CustomAudienceRule        `json:"rules"`
	RetentionDays  int                         `json:"retention_days,omitempty"`
}

type CustomAudienceRule struct {
	RuleType    string                 `json:"rule_type"` // URL, EVENT, APP_EVENT
	Conditions  []RuleCondition        `json:"conditions"`
	Operator    string                 `json:"operator,omitempty"` // AND, OR
}

type RuleCondition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"` // EQUALS, CONTAINS, STARTS_WITH, etc.
	Value    interface{} `json:"value"`
}

type CustomAudienceShareCancelRequest struct {
	AdvertiserID       string `json:"advertiser_id"`
	CustomAudienceID   string `json:"custom_audience_id"`
	TargetAdvertiserID string `json:"target_advertiser_id"`
}

type CustomAudienceShareCancelResponse struct {
	Code      int                              `json:"code"`
	Message   string                           `json:"message"`
	RequestID string                           `json:"request_id"`
	Data      CustomAudienceShareCancelData    `json:"data"`
}

type CustomAudienceShareCancelData struct {
	ShareID    string `json:"share_id"`
	Status     string `json:"status"`
	CancelTime string `json:"cancel_time"`
}

type CustomAudienceShareLogRequest struct {
	AdvertiserID       string `json:"advertiser_id"`
	CustomAudienceID   string `json:"custom_audience_id,omitempty"`
	Page               int    `json:"page,omitempty"`
	Size               int    `json:"size,omitempty"`
}

type CustomAudienceShareLogResponse struct {
	Code      int                           `json:"code"`
	Message   string                        `json:"message"`
	RequestID string                        `json:"request_id"`
	Data      CustomAudienceShareLogData    `json:"data"`
}

type CustomAudienceShareLogData struct {
	Logs     []CustomAudienceShareLogEntry `json:"logs"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type CustomAudienceShareLogEntry struct {
	LogID              string `json:"log_id"`
	ShareID            string `json:"share_id"`
	CustomAudienceID   string `json:"custom_audience_id"`
	TargetAdvertiserID string `json:"target_advertiser_id"`
	ShareType          string `json:"share_type"`
	Status             string `json:"status"`
	ShareTime          string `json:"share_time"`
	UpdateTime         string `json:"update_time"`
	Message            string `json:"message,omitempty"`
}
