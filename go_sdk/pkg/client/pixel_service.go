package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// PixelService handles Pixel related operations
type PixelService struct {
	client *Client
}

// NewPixelService creates a new PixelService
func NewPixelService(client *Client) *PixelService {
	return &PixelService{client: client}
}

// PixelCreateRequest represents the request for creating a pixel
type PixelCreateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	PixelName    string `json:"pixel_name"`
	PixelMode    string `json:"pixel_mode"` // MANUAL_MODE, CONVERSIONS_API_MODE, BOTH_MODE
	Description  string `json:"description,omitempty"`
}

// PixelGetRequest represents the request for getting pixel information
type PixelGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	PixelID      string `json:"pixel_id,omitempty"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

// PixelUpdateRequest represents the request for updating a pixel
type PixelUpdateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	PixelID      string `json:"pixel_id"`
	PixelName    string `json:"pixel_name,omitempty"`
	PixelMode    string `json:"pixel_mode,omitempty"`
	Description  string `json:"description,omitempty"`
}

// PixelData represents pixel information
type PixelData struct {
	PixelID      string `json:"pixel_id"`
	PixelName    string `json:"pixel_name"`
	PixelMode    string `json:"pixel_mode"`
	PixelCode    string `json:"pixel_code"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	LastFireTime string `json:"last_fire_time"`
	EventCount   int64  `json:"event_count"`
}

// PixelResponse represents the response from pixel operations
type PixelResponse struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	RequestID string    `json:"request_id"`
	Data      PixelData `json:"data"`
}

// PixelListResponse represents the response for listing pixels
type PixelListResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      []PixelData `json:"data"`
}

// Create creates a new pixel
func (s *PixelService) Create(ctx context.Context, req *PixelCreateRequest) (*PixelResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.PixelName == "" {
		return nil, fmt.Errorf("pixel_name is required")
	}
	if req.PixelMode == "" {
		req.PixelMode = "MANUAL_MODE"
	}

	url := s.client.BuildURL("/pixel/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create pixel: %w", err)
	}

	var response PixelResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// List retrieves pixel information
func (s *PixelService) List(ctx context.Context, req *PixelGetRequest) (*PixelListResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.PixelID != "" {
		params["pixel_id"] = req.PixelID
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/pixel/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get pixels: %w", err)
	}

	var response PixelListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Update updates an existing pixel
func (s *PixelService) Update(ctx context.Context, req *PixelUpdateRequest) (*PixelResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}

	url := s.client.BuildURL("/pixel/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update pixel: %w", err)
	}

	var response PixelResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PixelEventCreateRequest represents the request for creating a pixel event
type PixelEventCreateRequest struct {
	AdvertiserID string                 `json:"advertiser_id"`
	PixelID      string                 `json:"pixel_id"`
	EventName    string                 `json:"event_name"`
	EventType    string                 `json:"event_type"` // STANDARD, CUSTOM
	Description  string                 `json:"description,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
}

// PixelEventGetRequest represents the request for getting pixel events
type PixelEventGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	PixelID      string `json:"pixel_id"`
	EventID      string `json:"event_id,omitempty"`
}

// PixelEventUpdateRequest represents the request for updating a pixel event
type PixelEventUpdateRequest struct {
	AdvertiserID string                 `json:"advertiser_id"`
	PixelID      string                 `json:"pixel_id"`
	EventID      string                 `json:"event_id"`
	EventName    string                 `json:"event_name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
}

// PixelEventDeleteRequest represents the request for deleting a pixel event
type PixelEventDeleteRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	PixelID      string `json:"pixel_id"`
	EventID      string `json:"event_id"`
}

// PixelEventData represents pixel event information
type PixelEventData struct {
	EventID     string                 `json:"event_id"`
	EventName   string                 `json:"event_name"`
	EventType   string                 `json:"event_type"`
	Description string                 `json:"description"`
	Status      string                 `json:"status"`
	CreateTime  string                 `json:"create_time"`
	UpdateTime  string                 `json:"update_time"`
	Parameters  map[string]interface{} `json:"parameters"`
	FireCount   int64                  `json:"fire_count"`
}

// PixelEventResponse represents the response from pixel event operations
type PixelEventResponse struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	RequestID string         `json:"request_id"`
	Data      PixelEventData `json:"data"`
}

// PixelEventListResponse represents the response for listing pixel events
type PixelEventListResponse struct {
	Code      int              `json:"code"`
	Message   string           `json:"message"`
	RequestID string           `json:"request_id"`
	Data      []PixelEventData `json:"data"`
}

// CreateEvent creates a new pixel event
func (s *PixelService) CreateEvent(ctx context.Context, req *PixelEventCreateRequest) (*PixelEventResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}
	if req.EventName == "" {
		return nil, fmt.Errorf("event_name is required")
	}
	if req.EventType == "" {
		req.EventType = "STANDARD"
	}

	url := s.client.BuildURL("/pixel/event/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create pixel event: %w", err)
	}

	var response PixelEventResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEvents retrieves pixel events
func (s *PixelService) GetEvents(ctx context.Context, req *PixelEventGetRequest) (*PixelEventListResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.PixelID == "" {
		return nil, fmt.Errorf("advertiser_id and pixel_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"pixel_id":      req.PixelID,
	}
	if req.EventID != "" {
		params["event_id"] = req.EventID
	}

	url := s.client.BuildURL("/pixel/event/stats/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get pixel events: %w", err)
	}

	var response PixelEventListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateEvent updates an existing pixel event
func (s *PixelService) UpdateEvent(ctx context.Context, req *PixelEventUpdateRequest) (*PixelEventResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}
	if req.EventID == "" {
		return nil, fmt.Errorf("event_id is required")
	}

	url := s.client.BuildURL("/pixel/event/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update pixel event: %w", err)
	}

	var response PixelEventResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteEvent deletes a pixel event
func (s *PixelService) DeleteEvent(ctx context.Context, req *PixelEventDeleteRequest) (*PixelEventResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}
	if req.EventID == "" {
		return nil, fmt.Errorf("event_id is required")
	}

	url := s.client.BuildURL("/pixel/event/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete pixel event: %w", err)
	}

	var response PixelEventResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
