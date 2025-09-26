package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// reportService handles Report related operations
type reportService struct {
	client *Client
}

// NewReportService creates a new ReportService
func NewReportService(client *Client) ReportService {
	return &reportService{client: client}
}

// GetIntegratedReport retrieves integrated reporting data
func (s *reportService) GetIntegratedReport(ctx context.Context, req *ReportIntegratedGetRequest) (*ReportIntegratedResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.StartDate != "" {
		params["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		params["end_date"] = req.EndDate
	}
	if len(req.Dimensions) > 0 {
		params["dimensions"] = req.Dimensions
	}
	if len(req.Metrics) > 0 {
		params["metrics"] = req.Metrics
	}
	if req.ReportType != "" {
		params["report_type"] = req.ReportType
	}
	if req.DataLevel != "" {
		params["data_level"] = req.DataLevel
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/report/integrated/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get integrated report: %w", err)
	}

	var response ReportIntegratedResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateReportTask creates a new report generation task
func (s *reportService) CreateReportTask(ctx context.Context, req *ReportTaskCreateRequest) (*ReportTaskResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.ReportType == "" {
		return nil, fmt.Errorf("report_type is required")
	}
	if req.StartDate == "" || req.EndDate == "" {
		return nil, fmt.Errorf("start_date and end_date are required")
	}

	url := s.client.BuildURL("/report/task/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create report task: %w", err)
	}

	var response ReportTaskResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CheckReportTask checks the status of a report generation task
func (s *reportService) CheckReportTask(ctx context.Context, req *ReportTaskCheckRequest) (*ReportTaskCheckResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.TaskID == "" {
		return nil, fmt.Errorf("advertiser_id and task_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"task_id":       req.TaskID,
	}

	url := s.client.BuildURL("/report/task/check/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check report task: %w", err)
	}

	var response ReportTaskCheckResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CancelReportTask cancels a report generation task
func (s *reportService) CancelReportTask(ctx context.Context, req *ReportTaskCancelRequest) (*ReportTaskCancelResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.TaskID == "" {
		return nil, fmt.Errorf("task_id is required")
	}

	url := s.client.BuildURL("/report/task/cancel/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel report task: %w", err)
	}

	var response ReportTaskCancelResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Report types
type ReportIntegratedGetRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	StartDate    string   `json:"start_date,omitempty"`
	EndDate      string   `json:"end_date,omitempty"`
	Dimensions   []string `json:"dimensions,omitempty"`
	Metrics      []string `json:"metrics,omitempty"`
	ReportType   string   `json:"report_type,omitempty"` // BASIC, AUDIENCE, CONVERSION
	DataLevel    string   `json:"data_level,omitempty"`  // ADVERTISER, CAMPAIGN, ADGROUP, AD
	Page         int      `json:"page,omitempty"`
	Size         int      `json:"size,omitempty"`
	Filters      []ReportFilter `json:"filters,omitempty"`
}

type ReportFilter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"` // EQUALS, IN, GREATER_THAN, LESS_THAN
	Value    interface{} `json:"value"`
}

type ReportIntegratedResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      ReportIntegratedData   `json:"data"`
}

type ReportIntegratedData struct {
	List     []ReportDataRow `json:"list"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type ReportDataRow struct {
	Dimensions map[string]string `json:"dimensions"`
	Metrics    map[string]interface{} `json:"metrics"`
}

type ReportTaskCreateRequest struct {
	AdvertiserID string                 `json:"advertiser_id"`
	ReportType   string                 `json:"report_type"` // BASIC, AUDIENCE, CONVERSION, CUSTOM
	StartDate    string                 `json:"start_date"`
	EndDate      string                 `json:"end_date"`
	DataLevel    string                 `json:"data_level,omitempty"`
	Dimensions   []string               `json:"dimensions,omitempty"`
	Metrics      []string               `json:"metrics,omitempty"`
	Filters      []ReportFilter         `json:"filters,omitempty"`
	Format       string                 `json:"format,omitempty"` // CSV, JSON
	TimeZone     string                 `json:"time_zone,omitempty"`
	Settings     map[string]interface{} `json:"settings,omitempty"`
}

type ReportTaskResponse struct {
	Code      int              `json:"code"`
	Message   string           `json:"message"`
	RequestID string           `json:"request_id"`
	Data      ReportTaskData   `json:"data"`
}

type ReportTaskData struct {
	TaskID     string `json:"task_id"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
}

type ReportTaskCheckRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	TaskID       string `json:"task_id"`
}

type ReportTaskCheckResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      ReportTaskCheckData   `json:"data"`
}

type ReportTaskCheckData struct {
	TaskID       string `json:"task_id"`
	Status       string `json:"status"` // PENDING, PROCESSING, COMPLETED, FAILED
	Progress     int    `json:"progress"` // 0-100
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	CompleteTime string `json:"complete_time,omitempty"`
	DownloadURL  string `json:"download_url,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type ReportTaskCancelRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	TaskID       string `json:"task_id"`
}

type ReportTaskCancelResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      ReportTaskCancelData   `json:"data"`
}

type ReportTaskCancelData struct {
	TaskID     string `json:"task_id"`
	Status     string `json:"status"`
	CancelTime string `json:"cancel_time"`
}
