package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// commentService handles Comment related operations
type commentService struct {
	client *Client
}

// NewCommentService creates a new CommentService
func NewCommentService(client *Client) CommentService {
	return &commentService{client: client}
}

// ListComments retrieves a list of comments
func (s *commentService) ListComments(ctx context.Context, req *CommentListRequest) (*CommentListResponse, error) {
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
	if req.Status != "" {
		params["status"] = req.Status
	}
	if req.StartDate != "" {
		params["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		params["end_date"] = req.EndDate
	}

	url := s.client.BuildURL("/comment/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list comments: %w", err)
	}

	var response CommentListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PostComment posts a new comment
func (s *commentService) PostComment(ctx context.Context, req *CommentPostRequest) (*CommentPostResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CommentText == "" {
		return nil, fmt.Errorf("comment_text is required")
	}
	if req.VideoID == "" {
		return nil, fmt.Errorf("video_id is required")
	}

	url := s.client.BuildURL("/comment/post/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to post comment: %w", err)
	}

	var response CommentPostResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteComment deletes a comment
func (s *commentService) DeleteComment(ctx context.Context, req *CommentDeleteRequest) (*CommentDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CommentID == "" {
		return nil, fmt.Errorf("comment_id is required")
	}

	url := s.client.BuildURL("/comment/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete comment: %w", err)
	}

	var response CommentDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCommentStatus updates comment status
func (s *commentService) UpdateCommentStatus(ctx context.Context, req *CommentStatusUpdateRequest) (*CommentStatusUpdateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.CommentID == "" {
		return nil, fmt.Errorf("comment_id is required")
	}
	if req.Status == "" {
		return nil, fmt.Errorf("status is required")
	}

	url := s.client.BuildURL("/comment/status/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment status: %w", err)
	}

	var response CommentStatusUpdateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCommentReference retrieves comment reference information
func (s *commentService) GetCommentReference(ctx context.Context, req *CommentReferenceRequest) (*CommentReferenceResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.VideoID != "" {
		params["video_id"] = req.VideoID
	}

	url := s.client.BuildURL("/comment/reference/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment reference: %w", err)
	}

	var response CommentReferenceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateCommentTask creates a comment management task
func (s *commentService) CreateCommentTask(ctx context.Context, req *CommentTaskCreateRequest) (*CommentTaskResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.TaskType == "" {
		return nil, fmt.Errorf("task_type is required")
	}

	url := s.client.BuildURL("/comment/task/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment task: %w", err)
	}

	var response CommentTaskResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CheckCommentTask checks comment task status
func (s *commentService) CheckCommentTask(ctx context.Context, req *CommentTaskCheckRequest) (*CommentTaskCheckResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.TaskID == "" {
		return nil, fmt.Errorf("advertiser_id and task_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"task_id":       req.TaskID,
	}

	url := s.client.BuildURL("/comment/task/check/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check comment task: %w", err)
	}

	var response CommentTaskCheckResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Comment types
type CommentListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
	Status       string `json:"status,omitempty"` // ACTIVE, HIDDEN, DELETED
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
	VideoID      string `json:"video_id,omitempty"`
}

type CommentListResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      CommentListData   `json:"data"`
}

type CommentListData struct {
	Comments []CommentInfo `json:"comments"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type CommentInfo struct {
	CommentID   string `json:"comment_id"`
	VideoID     string `json:"video_id"`
	CommentText string `json:"comment_text"`
	Status      string `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	UserName    string `json:"user_name,omitempty"`
	LikeCount   int    `json:"like_count,omitempty"`
	ReplyCount  int    `json:"reply_count,omitempty"`
}

type CommentPostRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	VideoID      string `json:"video_id"`
	CommentText  string `json:"comment_text"`
	ParentID     string `json:"parent_id,omitempty"` // For replies
}

type CommentPostResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      CommentPostData   `json:"data"`
}

type CommentPostData struct {
	CommentID  string `json:"comment_id"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
}

type CommentDeleteRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CommentID    string `json:"comment_id"`
}

type CommentDeleteResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      CommentDeleteData   `json:"data"`
}

type CommentDeleteData struct {
	CommentID  string `json:"comment_id"`
	Status     string `json:"status"`
	DeleteTime string `json:"delete_time"`
}

type CommentStatusUpdateRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	CommentID    string `json:"comment_id"`
	Status       string `json:"status"` // ACTIVE, HIDDEN
}

type CommentStatusUpdateResponse struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	RequestID string                    `json:"request_id"`
	Data      CommentStatusUpdateData   `json:"data"`
}

type CommentStatusUpdateData struct {
	CommentID  string `json:"comment_id"`
	Status     string `json:"status"`
	UpdateTime string `json:"update_time"`
}

type CommentReferenceRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	VideoID      string `json:"video_id,omitempty"`
}

type CommentReferenceResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      CommentReferenceData   `json:"data"`
}

type CommentReferenceData struct {
	VideoInfo   VideoInfo     `json:"video_info"`
	Guidelines  []string      `json:"guidelines"`
	Policies    []PolicyInfo  `json:"policies"`
}

type VideoInfo struct {
	VideoID     string `json:"video_id"`
	VideoTitle  string `json:"video_title"`
	VideoURL    string `json:"video_url"`
	Duration    int    `json:"duration"`
	ViewCount   int64  `json:"view_count"`
	LikeCount   int64  `json:"like_count"`
	ShareCount  int64  `json:"share_count"`
}

type PolicyInfo struct {
	PolicyID   string `json:"policy_id"`
	PolicyName string `json:"policy_name"`
	PolicyURL  string `json:"policy_url"`
}

type CommentTaskCreateRequest struct {
	AdvertiserID string                 `json:"advertiser_id"`
	TaskType     string                 `json:"task_type"` // BULK_DELETE, BULK_HIDE, BULK_APPROVE
	Filters      map[string]interface{} `json:"filters"`
	VideoIDs     []string               `json:"video_ids,omitempty"`
}

type CommentTaskResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      CommentTaskData   `json:"data"`
}

type CommentTaskData struct {
	TaskID     string `json:"task_id"`
	TaskType   string `json:"task_type"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
}

type CommentTaskCheckRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	TaskID       string `json:"task_id"`
}

type CommentTaskCheckResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      CommentTaskCheckData   `json:"data"`
}

type CommentTaskCheckData struct {
	TaskID       string `json:"task_id"`
	TaskType     string `json:"task_type"`
	Status       string `json:"status"` // PENDING, PROCESSING, COMPLETED, FAILED
	Progress     int    `json:"progress"` // 0-100
	ProcessedCount int  `json:"processed_count"`
	TotalCount     int  `json:"total_count"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
	CompleteTime   string `json:"complete_time,omitempty"`
	ErrorMessage   string `json:"error_message,omitempty"`
}
