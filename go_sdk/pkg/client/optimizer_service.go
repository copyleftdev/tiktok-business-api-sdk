package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// optimizerService handles Optimizer related operations
type optimizerService struct {
	client *Client
}

// NewOptimizerService creates a new OptimizerService
func NewOptimizerService(client *Client) OptimizerService {
	return &optimizerService{client: client}
}

// CreateRule creates a new automated rule
func (s *optimizerService) CreateRule(ctx context.Context, req *OptimizerRuleCreateRequest) (*OptimizerRuleResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.RuleName == "" {
		return nil, fmt.Errorf("rule_name is required")
	}
	if len(req.Conditions) == 0 {
		return nil, fmt.Errorf("conditions is required")
	}
	if len(req.Actions) == 0 {
		return nil, fmt.Errorf("actions is required")
	}

	url := s.client.BuildURL("/optimizer/rule/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create optimizer rule: %w", err)
	}

	var response OptimizerRuleResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRule retrieves optimizer rule information
func (s *optimizerService) GetRule(ctx context.Context, req *OptimizerRuleGetRequest) (*OptimizerRuleResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.RuleID == "" {
		return nil, fmt.Errorf("advertiser_id and rule_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"rule_id":       req.RuleID,
	}

	url := s.client.BuildURL("/optimizer/rule/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get optimizer rule: %w", err)
	}

	var response OptimizerRuleResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListRules retrieves a list of optimizer rules
func (s *optimizerService) ListRules(ctx context.Context, req *OptimizerRuleListRequest) (*OptimizerRuleListResponse, error) {
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

	url := s.client.BuildURL("/optimizer/rule/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list optimizer rules: %w", err)
	}

	var response OptimizerRuleListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateRule updates an existing optimizer rule
func (s *optimizerService) UpdateRule(ctx context.Context, req *OptimizerRuleUpdateRequest) (*OptimizerRuleResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.RuleID == "" {
		return nil, fmt.Errorf("rule_id is required")
	}

	url := s.client.BuildURL("/optimizer/rule/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update optimizer rule: %w", err)
	}

	var response OptimizerRuleResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// BatchBindRule binds rules to campaigns/ad groups in batch
func (s *optimizerService) BatchBindRule(ctx context.Context, req *OptimizerRuleBatchBindRequest) (*OptimizerRuleBatchBindResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}
	if req.RuleID == "" {
		return nil, fmt.Errorf("rule_id is required")
	}
	if len(req.ObjectIDs) == 0 {
		return nil, fmt.Errorf("object_ids is required")
	}

	url := s.client.BuildURL("/optimizer/rule/batch/bind/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to batch bind optimizer rule: %w", err)
	}

	var response OptimizerRuleBatchBindResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRuleResult retrieves optimizer rule execution results
func (s *optimizerService) GetRuleResult(ctx context.Context, req *OptimizerRuleResultGetRequest) (*OptimizerRuleResultResponse, error) {
	if req == nil || req.AdvertiserID == "" || req.RuleID == "" {
		return nil, fmt.Errorf("advertiser_id and rule_id are required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"rule_id":       req.RuleID,
	}

	if req.StartDate != "" {
		params["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		params["end_date"] = req.EndDate
	}

	url := s.client.BuildURL("/optimizer/rule/result/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get optimizer rule result: %w", err)
	}

	var response OptimizerRuleResultResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListRuleResults retrieves a list of optimizer rule execution results
func (s *optimizerService) ListRuleResults(ctx context.Context, req *OptimizerRuleResultListRequest) (*OptimizerRuleResultListResponse, error) {
	if req == nil || req.AdvertiserID == "" {
		return nil, fmt.Errorf("advertiser_id is required")
	}

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.RuleID != "" {
		params["rule_id"] = req.RuleID
	}
	if req.StartDate != "" {
		params["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		params["end_date"] = req.EndDate
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/optimizer/rule/result/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list optimizer rule results: %w", err)
	}

	var response OptimizerRuleResultListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Optimizer types
type OptimizerRuleCreateRequest struct {
	AdvertiserID string                   `json:"advertiser_id"`
	RuleName     string                   `json:"rule_name"`
	Description  string                   `json:"description,omitempty"`
	Status       string                   `json:"status,omitempty"` // ACTIVE, PAUSED
	Conditions   []OptimizerRuleCondition `json:"conditions"`
	Actions      []OptimizerRuleAction    `json:"actions"`
	ObjectType   string                   `json:"object_type"` // CAMPAIGN, ADGROUP
	ObjectIDs    []string                 `json:"object_ids,omitempty"`
}

type OptimizerRuleCondition struct {
	Metric    string  `json:"metric"`    // CPC, CPM, CTR, CVR, etc.
	Operator  string  `json:"operator"`  // GREATER_THAN, LESS_THAN, EQUAL_TO
	Value     float64 `json:"value"`
	TimeRange string  `json:"time_range"` // LAST_7_DAYS, LAST_14_DAYS, etc.
}

type OptimizerRuleAction struct {
	ActionType string      `json:"action_type"` // PAUSE, ADJUST_BID, ADJUST_BUDGET
	Value      interface{} `json:"value,omitempty"`
}

type OptimizerRuleResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      OptimizerRuleData   `json:"data"`
}

type OptimizerRuleData struct {
	RuleID      string                   `json:"rule_id"`
	RuleName    string                   `json:"rule_name"`
	Description string                   `json:"description"`
	Status      string                   `json:"status"`
	Conditions  []OptimizerRuleCondition `json:"conditions"`
	Actions     []OptimizerRuleAction    `json:"actions"`
	ObjectType  string                   `json:"object_type"`
	CreateTime  string                   `json:"create_time"`
	UpdateTime  string                   `json:"update_time"`
}

type OptimizerRuleGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	RuleID       string `json:"rule_id"`
}

type OptimizerRuleListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
	Status       string `json:"status,omitempty"`
}

type OptimizerRuleListResponse struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	RequestID string                  `json:"request_id"`
	Data      OptimizerRuleListData   `json:"data"`
}

type OptimizerRuleListData struct {
	Rules    []OptimizerRuleData `json:"rules"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type OptimizerRuleUpdateRequest struct {
	AdvertiserID string                   `json:"advertiser_id"`
	RuleID       string                   `json:"rule_id"`
	RuleName     string                   `json:"rule_name,omitempty"`
	Description  string                   `json:"description,omitempty"`
	Status       string                   `json:"status,omitempty"`
	Conditions   []OptimizerRuleCondition `json:"conditions,omitempty"`
	Actions      []OptimizerRuleAction    `json:"actions,omitempty"`
}

type OptimizerRuleBatchBindRequest struct {
	AdvertiserID string   `json:"advertiser_id"`
	RuleID       string   `json:"rule_id"`
	ObjectIDs    []string `json:"object_ids"`
	ObjectType   string   `json:"object_type"` // CAMPAIGN, ADGROUP
}

type OptimizerRuleBatchBindResponse struct {
	Code      int                          `json:"code"`
	Message   string                       `json:"message"`
	RequestID string                       `json:"request_id"`
	Data      OptimizerRuleBatchBindData   `json:"data"`
}

type OptimizerRuleBatchBindData struct {
	Results []OptimizerRuleBindResult `json:"results"`
}

type OptimizerRuleBindResult struct {
	ObjectID string `json:"object_id"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

type OptimizerRuleResultGetRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	RuleID       string `json:"rule_id"`
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
}

type OptimizerRuleResultResponse struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	RequestID string                    `json:"request_id"`
	Data      OptimizerRuleResultData   `json:"data"`
}

type OptimizerRuleResultData struct {
	RuleID      string                      `json:"rule_id"`
	Executions  []OptimizerRuleExecution    `json:"executions"`
}

type OptimizerRuleExecution struct {
	ExecutionID   string `json:"execution_id"`
	ObjectID      string `json:"object_id"`
	ObjectType    string `json:"object_type"`
	ActionType    string `json:"action_type"`
	Status        string `json:"status"`
	ExecutionTime string `json:"execution_time"`
	Message       string `json:"message,omitempty"`
}

type OptimizerRuleResultListRequest struct {
	AdvertiserID string `json:"advertiser_id"`
	RuleID       string `json:"rule_id,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
	Page         int    `json:"page,omitempty"`
	Size         int    `json:"size,omitempty"`
}

type OptimizerRuleResultListResponse struct {
	Code      int                           `json:"code"`
	Message   string                        `json:"message"`
	RequestID string                        `json:"request_id"`
	Data      OptimizerRuleResultListData   `json:"data"`
}

type OptimizerRuleResultListData struct {
	Results  []OptimizerRuleResultData `json:"results"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}
