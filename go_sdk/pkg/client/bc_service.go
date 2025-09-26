package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// BusinessCenterService handles Business Center related operations
type BusinessCenterService struct {
	client *Client
}

// NewBusinessCenterService creates a new BusinessCenterService
func NewBusinessCenterService(client *Client) *BusinessCenterService {
	return &BusinessCenterService{client: client}
}

// BCGetRequest represents the request for getting business center information
type BCGetRequest struct {
	BCID  string `json:"bc_id,omitempty"`
	Scene string `json:"scene,omitempty"` // SINGLE_ACCOUNT or TIERED_ACCOUNT
}

// BCCreateRequest represents the request for creating a business center
type BCCreateRequest struct {
	BCName      string      `json:"bc_name"`
	CompanyName string      `json:"company_name"`
	ContactInfo ContactInfo `json:"contact_info"`
	TimeZone    string      `json:"timezone,omitempty"`
}

// ContactInfo represents contact information for business center
type ContactInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone,omitempty"`
}

// BCResponse represents the response from business center operations
type BCResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      BCData `json:"data"`
}

// BCData represents business center data
type BCData struct {
	BCID        string      `json:"bc_id"`
	BCName      string      `json:"bc_name"`
	CompanyName string      `json:"company_name"`
	ContactInfo ContactInfo `json:"contact_info"`
	TimeZone    string      `json:"timezone"`
	Status      string      `json:"status"`
	CreateTime  string      `json:"create_time"`
	UpdateTime  string      `json:"update_time"`
}

// BCListResponse represents the response for listing business centers
type BCListResponse struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	RequestID string   `json:"request_id"`
	Data      []BCData `json:"data"`
}

// Get retrieves business center information
func (s *BusinessCenterService) Get(ctx context.Context, req *BCGetRequest) (*BCResponse, error) {
	if req == nil {
		req = &BCGetRequest{}
	}

	params := make(map[string]interface{})
	if req.BCID != "" {
		params["bc_id"] = req.BCID
	}
	if req.Scene != "" {
		params["scene"] = req.Scene
	} else {
		params["scene"] = "SINGLE_ACCOUNT"
	}

	url := s.client.BuildURL("/bc/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get business center: %w", err)
	}

	var response BCResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Create creates a new business center
func (s *BusinessCenterService) Create(ctx context.Context, req *BCCreateRequest) (*BCResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.BCName == "" {
		return nil, fmt.Errorf("bc_name is required")
	}
	if req.CompanyName == "" {
		return nil, fmt.Errorf("company_name is required")
	}
	if req.ContactInfo.Email == "" {
		return nil, fmt.Errorf("contact email is required")
	}
	if req.ContactInfo.Name == "" {
		return nil, fmt.Errorf("contact name is required")
	}

	url := s.client.BuildURL("/bc/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create business center: %w", err)
	}

	var response BCResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// BCMemberGetRequest represents the request for getting BC members
type BCMemberGetRequest struct {
	BCID string `json:"bc_id"`
	Page int    `json:"page,omitempty"`
	Size int    `json:"size,omitempty"`
}

// BCMemberData represents BC member information
type BCMemberData struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	JoinTime string `json:"join_time"`
}

// BCMemberResponse represents the response for BC member operations
type BCMemberResponse struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	RequestID string         `json:"request_id"`
	Data      []BCMemberData `json:"data"`
}

// GetMembers retrieves business center members
func (s *BusinessCenterService) GetMembers(ctx context.Context, req *BCMemberGetRequest) (*BCMemberResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/bc/member/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get BC members: %w", err)
	}

	var response BCMemberResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// BCAssetGetRequest represents the request for getting BC assets
type BCAssetGetRequest struct {
	BCID      string `json:"bc_id"`
	AssetType string `json:"asset_type,omitempty"` // ADVERTISER, PIXEL, CATALOG, etc.
	Page      int    `json:"page,omitempty"`
	Size      int    `json:"size,omitempty"`
}

// BCAssetData represents BC asset information
type BCAssetData struct {
	AssetID   string `json:"asset_id"`
	AssetType string `json:"asset_type"`
	AssetName string `json:"asset_name"`
	Status    string `json:"status"`
	CreateTime string `json:"create_time"`
}

// BCAssetResponse represents the response for BC asset operations
type BCAssetResponse struct {
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	RequestID string        `json:"request_id"`
	Data      []BCAssetData `json:"data"`
}

// GetAssets retrieves business center assets
func (s *BusinessCenterService) GetAssets(ctx context.Context, req *BCAssetGetRequest) (*BCAssetResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}
	if req.AssetType != "" {
		params["asset_type"] = req.AssetType
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/bc/asset/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get BC assets: %w", err)
	}

	var response BCAssetResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional BC request/response types
type BCTransferRequest struct {
	BCID          string  `json:"bc_id"`
	FromAccountID string  `json:"from_account_id"`
	ToAccountID   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency,omitempty"`
	Description   string  `json:"description,omitempty"`
}

type BCTransferResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      BCTransferData    `json:"data"`
}

type BCTransferData struct {
	TransferID    string  `json:"transfer_id"`
	Status        string  `json:"status"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	TransferTime  string  `json:"transfer_time"`
}

type BCBalanceGetRequest struct {
	BCID      string `json:"bc_id"`
	AccountID string `json:"account_id,omitempty"`
}

type BCBalanceResponse struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      BCBalanceData   `json:"data"`
}

type BCBalanceData struct {
	Balances []BalanceInfo `json:"balances"`
}

type BalanceInfo struct {
	AccountID string  `json:"account_id"`
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance"`
	Available float64 `json:"available"`
	Pending   float64 `json:"pending"`
}

type BCMemberInviteRequest struct {
	BCID        string   `json:"bc_id"`
	Email       string   `json:"email"`
	Role        string   `json:"role"` // ADMIN, MEMBER, VIEWER
	Permissions []string `json:"permissions,omitempty"`
	Message     string   `json:"message,omitempty"`
}

type BCMemberInviteResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCMemberInviteData   `json:"data"`
}

type BCMemberInviteData struct {
	InviteID   string `json:"invite_id"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	InviteTime string `json:"invite_time"`
}

type BCMemberUpdateRequest struct {
	BCID        string   `json:"bc_id"`
	MemberID    string   `json:"member_id"`
	Role        string   `json:"role,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Status      string   `json:"status,omitempty"` // ACTIVE, SUSPENDED
}

type BCMemberUpdateResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCMemberUpdateData   `json:"data"`
}

type BCMemberUpdateData struct {
	MemberID   string `json:"member_id"`
	Status     string `json:"status"`
	UpdateTime string `json:"update_time"`
}

type BCMemberDeleteRequest struct {
	BCID     string `json:"bc_id"`
	MemberID string `json:"member_id"`
}

type BCMemberDeleteResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCMemberDeleteData   `json:"data"`
}

type BCMemberDeleteData struct {
	MemberID   string `json:"member_id"`
	Status     string `json:"status"`
	DeleteTime string `json:"delete_time"`
}

type BCMemberAssignRequest struct {
	BCID     string   `json:"bc_id"`
	MemberID string   `json:"member_id"`
	AssetIDs []string `json:"asset_ids"`
	AssetType string  `json:"asset_type"` // ADVERTISER, PIXEL, PAGE
}

type BCMemberAssignResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      BCMemberAssignData    `json:"data"`
}

type BCMemberAssignData struct {
	MemberID     string              `json:"member_id"`
	Assignments  []AssetAssignment   `json:"assignments"`
}

type AssetAssignment struct {
	AssetID string `json:"asset_id"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type BCAssetAssignRequest struct {
	BCID      string `json:"bc_id"`
	AssetID   string `json:"asset_id"`
	AssetType string `json:"asset_type"` // ADVERTISER, PIXEL, PAGE
}

type BCAssetAssignResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      BCAssetAssignData   `json:"data"`
}

type BCAssetAssignData struct {
	AssetID    string `json:"asset_id"`
	Status     string `json:"status"`
	AssignTime string `json:"assign_time"`
}

type BCAssetUnassignRequest struct {
	BCID    string `json:"bc_id"`
	AssetID string `json:"asset_id"`
}

type BCAssetUnassignResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      BCAssetUnassignData   `json:"data"`
}

type BCAssetUnassignData struct {
	AssetID      string `json:"asset_id"`
	Status       string `json:"status"`
	UnassignTime string `json:"unassign_time"`
}

// Transfer transfers funds between business center accounts
func (s *BusinessCenterService) Transfer(ctx context.Context, req *BCTransferRequest) (*BCTransferResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.FromAccountID == "" {
		return nil, fmt.Errorf("from_account_id is required")
	}
	if req.ToAccountID == "" {
		return nil, fmt.Errorf("to_account_id is required")
	}
	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be positive")
	}

	url := s.client.BuildURL("/bc/transfer/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer funds: %w", err)
	}

	var response BCTransferResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalance retrieves business center balance information
func (s *BusinessCenterService) GetBalance(ctx context.Context, req *BCBalanceGetRequest) (*BCBalanceResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.AccountID != "" {
		params["account_id"] = req.AccountID
	}

	url := s.client.BuildURL("/bc/balance/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	var response BCBalanceResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// InviteMember invites a new member to the business center
func (s *BusinessCenterService) InviteMember(ctx context.Context, req *BCMemberInviteRequest) (*BCMemberInviteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if req.Role == "" {
		return nil, fmt.Errorf("role is required")
	}

	url := s.client.BuildURL("/bc/member/invite/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to invite member: %w", err)
	}

	var response BCMemberInviteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateMember updates a business center member
func (s *BusinessCenterService) UpdateMember(ctx context.Context, req *BCMemberUpdateRequest) (*BCMemberUpdateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.MemberID == "" {
		return nil, fmt.Errorf("member_id is required")
	}

	url := s.client.BuildURL("/bc/member/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update member: %w", err)
	}

	var response BCMemberUpdateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteMember removes a member from the business center
func (s *BusinessCenterService) DeleteMember(ctx context.Context, req *BCMemberDeleteRequest) (*BCMemberDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.MemberID == "" {
		return nil, fmt.Errorf("member_id is required")
	}

	url := s.client.BuildURL("/bc/member/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete member: %w", err)
	}

	var response BCMemberDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AssignMember assigns a member to specific assets
func (s *BusinessCenterService) AssignMember(ctx context.Context, req *BCMemberAssignRequest) (*BCMemberAssignResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.MemberID == "" {
		return nil, fmt.Errorf("member_id is required")
	}
	if len(req.AssetIDs) == 0 {
		return nil, fmt.Errorf("asset_ids is required")
	}

	url := s.client.BuildURL("/bc/member/assign/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to assign member: %w", err)
	}

	var response BCMemberAssignResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AssignAsset assigns an asset to the business center
func (s *BusinessCenterService) AssignAsset(ctx context.Context, req *BCAssetAssignRequest) (*BCAssetAssignResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.AssetID == "" {
		return nil, fmt.Errorf("asset_id is required")
	}
	if req.AssetType == "" {
		return nil, fmt.Errorf("asset_type is required")
	}

	url := s.client.BuildURL("/bc/asset/assign/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to assign asset: %w", err)
	}

	var response BCAssetAssignResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UnassignAsset removes an asset from the business center
func (s *BusinessCenterService) UnassignAsset(ctx context.Context, req *BCAssetUnassignRequest) (*BCAssetUnassignResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.AssetID == "" {
		return nil, fmt.Errorf("asset_id is required")
	}

	url := s.client.BuildURL("/bc/asset/unassign/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to unassign asset: %w", err)
	}

	var response BCAssetUnassignResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional BC types for new methods
type BCTransactionGetRequest struct {
	BCID      string `json:"bc_id"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Page      int    `json:"page,omitempty"`
	Size      int    `json:"size,omitempty"`
}

type BCTransactionResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      BCTransactionData   `json:"data"`
}

type BCTransactionData struct {
	Transactions []TransactionInfo `json:"transactions"`
	PageInfo     struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type TransactionInfo struct {
	TransactionID string  `json:"transaction_id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	CreateTime    string  `json:"create_time"`
}

type BCAssetGroupCreateRequest struct {
	BCID        string `json:"bc_id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description,omitempty"`
}

type BCAssetGroupResponse struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	RequestID string             `json:"request_id"`
	Data      BCAssetGroupData   `json:"data"`
}

type BCAssetGroupData struct {
	GroupID     string `json:"group_id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type BCAssetGroupGetRequest struct {
	BCID    string `json:"bc_id"`
	GroupID string `json:"group_id,omitempty"`
}

type BCAssetGroupListResponse struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id"`
	Data      BCAssetGroupListData   `json:"data"`
}

type BCAssetGroupListData struct {
	Groups []BCAssetGroupData `json:"groups"`
}

type BCAssetGroupUpdateRequest struct {
	BCID        string `json:"bc_id"`
	GroupID     string `json:"group_id"`
	GroupName   string `json:"group_name,omitempty"`
	Description string `json:"description,omitempty"`
}

type BCAssetGroupDeleteRequest struct {
	BCID    string `json:"bc_id"`
	GroupID string `json:"group_id"`
}

type BCImageUploadRequest struct {
	BCID      string `json:"bc_id"`
	ImageData []byte `json:"image_data"`
	ImageName string `json:"image_name"`
	ImageType string `json:"image_type"` // JPG, PNG, GIF
}

type BCImageUploadResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      BCImageUploadData   `json:"data"`
}

type BCImageUploadData struct {
	ImageID   string `json:"image_id"`
	ImageURL  string `json:"image_url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Size      int64  `json:"size"`
}

// GetTransactions retrieves business center transaction history
func (s *BusinessCenterService) GetTransactions(ctx context.Context, req *BCTransactionGetRequest) (*BCTransactionResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
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

	url := s.client.BuildURL("/bc/transaction/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	var response BCTransactionResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateAssetGroup creates a new asset group
func (s *BusinessCenterService) CreateAssetGroup(ctx context.Context, req *BCAssetGroupCreateRequest) (*BCAssetGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.GroupName == "" {
		return nil, fmt.Errorf("group_name is required")
	}

	url := s.client.BuildURL("/bc/asset_group/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create asset group: %w", err)
	}

	var response BCAssetGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAssetGroups retrieves asset groups
func (s *BusinessCenterService) GetAssetGroups(ctx context.Context, req *BCAssetGroupGetRequest) (*BCAssetGroupListResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.GroupID != "" {
		params["group_id"] = req.GroupID
	}

	url := s.client.BuildURL("/bc/asset_group/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset groups: %w", err)
	}

	var response BCAssetGroupListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateAssetGroup updates an asset group
func (s *BusinessCenterService) UpdateAssetGroup(ctx context.Context, req *BCAssetGroupUpdateRequest) (*BCAssetGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.GroupID == "" {
		return nil, fmt.Errorf("group_id is required")
	}

	url := s.client.BuildURL("/bc/asset_group/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update asset group: %w", err)
	}

	var response BCAssetGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteAssetGroup deletes an asset group
func (s *BusinessCenterService) DeleteAssetGroup(ctx context.Context, req *BCAssetGroupDeleteRequest) (*BCAssetGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.GroupID == "" {
		return nil, fmt.Errorf("group_id is required")
	}

	url := s.client.BuildURL("/bc/asset_group/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete asset group: %w", err)
	}

	var response BCAssetGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UploadImage uploads an image to business center
func (s *BusinessCenterService) UploadImage(ctx context.Context, req *BCImageUploadRequest) (*BCImageUploadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if len(req.ImageData) == 0 {
		return nil, fmt.Errorf("image_data is required")
	}

	url := s.client.BuildURL("/bc/image/upload/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	var response BCImageUploadResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListAssetGroups retrieves a list of asset groups
func (s *BusinessCenterService) ListAssetGroups(ctx context.Context, req *BCAssetGroupListRequest) (*BCAssetGroupListResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/bc/asset_group/list/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list asset groups: %w", err)
	}

	var response BCAssetGroupListResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAssetMembers retrieves asset member information
func (s *BusinessCenterService) GetAssetMembers(ctx context.Context, req *BCAssetMemberGetRequest) (*BCAssetMemberResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.AssetID != "" {
		params["asset_id"] = req.AssetID
	}
	if req.AssetType != "" {
		params["asset_type"] = req.AssetType
	}

	url := s.client.BuildURL("/bc/asset_member/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset members: %w", err)
	}

	var response BCAssetMemberResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAssetPartners retrieves asset partner information
func (s *BusinessCenterService) GetAssetPartners(ctx context.Context, req *BCAssetPartnerGetRequest) (*BCAssetPartnerResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.AssetID != "" {
		params["asset_id"] = req.AssetID
	}

	url := s.client.BuildURL("/bc/asset_partner/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset partners: %w", err)
	}

	var response BCAssetPartnerResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAssetAdmins retrieves asset admin information
func (s *BusinessCenterService) GetAssetAdmins(ctx context.Context, req *BCAssetAdminGetRequest) (*BCAssetAdminResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.AssetID != "" {
		params["asset_id"] = req.AssetID
	}

	url := s.client.BuildURL("/bc/asset_admin/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset admins: %w", err)
	}

	var response BCAssetAdminResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteAssetAdmin removes an asset admin
func (s *BusinessCenterService) DeleteAssetAdmin(ctx context.Context, req *BCAssetAdminDeleteRequest) (*BCAssetAdminDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.AssetID == "" {
		return nil, fmt.Errorf("asset_id is required")
	}
	if req.UserID == "" {
		return nil, fmt.Errorf("user_id is required")
	}

	url := s.client.BuildURL("/bc/asset_admin/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete asset admin: %w", err)
	}

	var response BCAssetAdminDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional BC types for new methods
type BCAssetGroupListRequest struct {
	BCID string `json:"bc_id"`
	Page int    `json:"page,omitempty"`
	Size int    `json:"size,omitempty"`
}

type BCAssetMemberGetRequest struct {
	BCID      string `json:"bc_id"`
	AssetID   string `json:"asset_id,omitempty"`
	AssetType string `json:"asset_type,omitempty"`
}

type BCAssetMemberResponse struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	RequestID string              `json:"request_id"`
	Data      BCAssetMemberData   `json:"data"`
}

type BCAssetMemberData struct {
	Members []BCAssetMember `json:"members"`
}

type BCAssetMember struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	AssignTime  string `json:"assign_time"`
}

type BCAssetPartnerGetRequest struct {
	BCID    string `json:"bc_id"`
	AssetID string `json:"asset_id,omitempty"`
}

type BCAssetPartnerResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCAssetPartnerData   `json:"data"`
}

type BCAssetPartnerData struct {
	Partners []BCAssetPartner `json:"partners"`
}

type BCAssetPartner struct {
	PartnerID   string `json:"partner_id"`
	PartnerName string `json:"partner_name"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	AssignTime  string `json:"assign_time"`
}

type BCAssetAdminGetRequest struct {
	BCID    string `json:"bc_id"`
	AssetID string `json:"asset_id,omitempty"`
}

type BCAssetAdminResponse struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	RequestID string             `json:"request_id"`
	Data      BCAssetAdminData   `json:"data"`
}

type BCAssetAdminData struct {
	Admins []BCAssetAdmin `json:"admins"`
}

type BCAssetAdmin struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	AssignTime string `json:"assign_time"`
}

type BCAssetAdminDeleteRequest struct {
	BCID    string `json:"bc_id"`
	AssetID string `json:"asset_id"`
	UserID  string `json:"user_id"`
}

type BCAssetAdminDeleteResponse struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	RequestID string                   `json:"request_id"`
	Data      BCAssetAdminDeleteData   `json:"data"`
}

type BCAssetAdminDeleteData struct {
	UserID     string `json:"user_id"`
	Status     string `json:"status"`
	DeleteTime string `json:"delete_time"`
}

// GetAccountTransactions retrieves account transaction history
func (s *BusinessCenterService) GetAccountTransactions(ctx context.Context, req *BCAccountTransactionGetRequest) (*BCAccountTransactionResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.AccountID != "" {
		params["account_id"] = req.AccountID
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

	url := s.client.BuildURL("/bc/account_transaction/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account transactions: %w", err)
	}

	var response BCAccountTransactionResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateBillingGroup creates a new billing group
func (s *BusinessCenterService) CreateBillingGroup(ctx context.Context, req *BCBillingGroupCreateRequest) (*BCBillingGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.GroupName == "" {
		return nil, fmt.Errorf("group_name is required")
	}

	url := s.client.BuildURL("/bc/billing_group/create/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create billing group: %w", err)
	}

	var response BCBillingGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBillingGroup retrieves billing group information
func (s *BusinessCenterService) GetBillingGroup(ctx context.Context, req *BCBillingGroupGetRequest) (*BCBillingGroupResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.GroupID != "" {
		params["group_id"] = req.GroupID
	}

	url := s.client.BuildURL("/bc/billing_group/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get billing group: %w", err)
	}

	var response BCBillingGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateBillingGroup updates a billing group
func (s *BusinessCenterService) UpdateBillingGroup(ctx context.Context, req *BCBillingGroupUpdateRequest) (*BCBillingGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.GroupID == "" {
		return nil, fmt.Errorf("group_id is required")
	}

	url := s.client.BuildURL("/bc/billing_group/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update billing group: %w", err)
	}

	var response BCBillingGroupResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUnpaidInvoices retrieves unpaid invoice information
func (s *BusinessCenterService) GetUnpaidInvoices(ctx context.Context, req *BCInvoiceUnpaidGetRequest) (*BCInvoiceUnpaidResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := s.client.BuildURL("/bc/invoice_unpaid/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get unpaid invoices: %w", err)
	}

	var response BCInvoiceUnpaidResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// AddPartner adds a partner to the business center
func (s *BusinessCenterService) AddPartner(ctx context.Context, req *BCPartnerAddRequest) (*BCPartnerAddResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.PartnerID == "" {
		return nil, fmt.Errorf("partner_id is required")
	}

	url := s.client.BuildURL("/bc/partner/add/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to add partner: %w", err)
	}

	var response BCPartnerAddResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPartner retrieves partner information
func (s *BusinessCenterService) GetPartner(ctx context.Context, req *BCPartnerGetRequest) (*BCPartnerResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.PartnerID != "" {
		params["partner_id"] = req.PartnerID
	}

	url := s.client.BuildURL("/bc/partner/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get partner: %w", err)
	}

	var response BCPartnerResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeletePartner removes a partner from the business center
func (s *BusinessCenterService) DeletePartner(ctx context.Context, req *BCPartnerDeleteRequest) (*BCPartnerDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.PartnerID == "" {
		return nil, fmt.Errorf("partner_id is required")
	}

	url := s.client.BuildURL("/bc/partner/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete partner: %w", err)
	}

	var response BCPartnerDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPartnerAssets retrieves partner asset information
func (s *BusinessCenterService) GetPartnerAssets(ctx context.Context, req *BCPartnerAssetGetRequest) (*BCPartnerAssetResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.PartnerID != "" {
		params["partner_id"] = req.PartnerID
	}
	if req.AssetType != "" {
		params["asset_type"] = req.AssetType
	}

	url := s.client.BuildURL("/bc/partner_asset/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get partner assets: %w", err)
	}

	var response BCPartnerAssetResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeletePartnerAsset removes a partner asset
func (s *BusinessCenterService) DeletePartnerAsset(ctx context.Context, req *BCPartnerAssetDeleteRequest) (*BCPartnerAssetDeleteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.PartnerID == "" {
		return nil, fmt.Errorf("partner_id is required")
	}
	if req.AssetID == "" {
		return nil, fmt.Errorf("asset_id is required")
	}

	url := s.client.BuildURL("/bc/partner_asset/delete/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete partner asset: %w", err)
	}

	var response BCPartnerAssetDeleteResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPixelLinks retrieves pixel link information
func (s *BusinessCenterService) GetPixelLinks(ctx context.Context, req *BCPixelLinkGetRequest) (*BCPixelLinkResponse, error) {
	if req == nil || req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}

	params := map[string]interface{}{
		"bc_id": req.BCID,
	}

	if req.PixelID != "" {
		params["pixel_id"] = req.PixelID
	}

	url := s.client.BuildURL("/bc/pixel_link/get/", params)
	
	resp, err := s.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get pixel links: %w", err)
	}

	var response BCPixelLinkResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdatePixelLink updates pixel link configuration
func (s *BusinessCenterService) UpdatePixelLink(ctx context.Context, req *BCPixelLinkUpdateRequest) (*BCPixelLinkUpdateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}

	url := s.client.BuildURL("/bc/pixel_link/update/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update pixel link: %w", err)
	}

	var response BCPixelLinkUpdateResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// TransferPixel transfers pixel ownership
func (s *BusinessCenterService) TransferPixel(ctx context.Context, req *BCPixelTransferRequest) (*BCPixelTransferResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.BCID == "" {
		return nil, fmt.Errorf("bc_id is required")
	}
	if req.PixelID == "" {
		return nil, fmt.Errorf("pixel_id is required")
	}
	if req.TargetBCID == "" {
		return nil, fmt.Errorf("target_bc_id is required")
	}

	url := s.client.BuildURL("/bc/pixel/transfer/", nil)
	
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.DoRequest(ctx, "POST", url, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer pixel: %w", err)
	}

	var response BCPixelTransferResponse
	if err := s.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Additional BC types for remaining methods
type BCAccountTransactionGetRequest struct {
	BCID      string `json:"bc_id"`
	AccountID string `json:"account_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Page      int    `json:"page,omitempty"`
	Size      int    `json:"size,omitempty"`
}

type BCAccountTransactionResponse struct {
	Code      int                         `json:"code"`
	Message   string                      `json:"message"`
	RequestID string                      `json:"request_id"`
	Data      BCAccountTransactionData    `json:"data"`
}

type BCAccountTransactionData struct {
	Transactions []BCAccountTransaction `json:"transactions"`
	PageInfo     struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type BCAccountTransaction struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Status          string  `json:"status"`
	CreateTime      string  `json:"create_time"`
	UpdateTime      string  `json:"update_time"`
	Description     string  `json:"description,omitempty"`
}

type BCBillingGroupCreateRequest struct {
	BCID        string   `json:"bc_id"`
	GroupName   string   `json:"group_name"`
	Description string   `json:"description,omitempty"`
	AccountIDs  []string `json:"account_ids,omitempty"`
}

type BCBillingGroupGetRequest struct {
	BCID    string `json:"bc_id"`
	GroupID string `json:"group_id,omitempty"`
}

type BCBillingGroupUpdateRequest struct {
	BCID        string   `json:"bc_id"`
	GroupID     string   `json:"group_id"`
	GroupName   string   `json:"group_name,omitempty"`
	Description string   `json:"description,omitempty"`
	AccountIDs  []string `json:"account_ids,omitempty"`
}

type BCBillingGroupResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCBillingGroupData   `json:"data"`
}

type BCBillingGroupData struct {
	Groups []BCBillingGroup `json:"groups"`
}

type BCBillingGroup struct {
	GroupID     string   `json:"group_id"`
	GroupName   string   `json:"group_name"`
	Description string   `json:"description"`
	AccountIDs  []string `json:"account_ids"`
	CreateTime  string   `json:"create_time"`
	UpdateTime  string   `json:"update_time"`
}

type BCInvoiceUnpaidGetRequest struct {
	BCID string `json:"bc_id"`
	Page int    `json:"page,omitempty"`
	Size int    `json:"size,omitempty"`
}

type BCInvoiceUnpaidResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      BCInvoiceUnpaidData   `json:"data"`
}

type BCInvoiceUnpaidData struct {
	Invoices []BCInvoiceUnpaid `json:"invoices"`
	PageInfo struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalCount int `json:"total_count"`
	} `json:"page_info"`
}

type BCInvoiceUnpaid struct {
	InvoiceID    string  `json:"invoice_id"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	DueDate      string  `json:"due_date"`
	Status       string  `json:"status"`
	CreateTime   string  `json:"create_time"`
	Description  string  `json:"description,omitempty"`
}

type BCPartnerAddRequest struct {
	BCID      string `json:"bc_id"`
	PartnerID string `json:"partner_id"`
	Role      string `json:"role,omitempty"`
}

type BCPartnerGetRequest struct {
	BCID      string `json:"bc_id"`
	PartnerID string `json:"partner_id,omitempty"`
}

type BCPartnerDeleteRequest struct {
	BCID      string `json:"bc_id"`
	PartnerID string `json:"partner_id"`
}

type BCPartnerAddResponse struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	RequestID string             `json:"request_id"`
	Data      BCPartnerAddData   `json:"data"`
}

type BCPartnerAddData struct {
	PartnerID  string `json:"partner_id"`
	Status     string `json:"status"`
	AssignTime string `json:"assign_time"`
}

type BCPartnerDeleteResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      BCPartnerDeleteData   `json:"data"`
}

type BCPartnerDeleteData struct {
	PartnerID  string `json:"partner_id"`
	Status     string `json:"status"`
	DeleteTime string `json:"delete_time"`
}

type BCPartnerAssetGetRequest struct {
	BCID      string `json:"bc_id"`
	PartnerID string `json:"partner_id,omitempty"`
	AssetType string `json:"asset_type,omitempty"`
}

type BCPartnerAssetDeleteRequest struct {
	BCID      string `json:"bc_id"`
	PartnerID string `json:"partner_id"`
	AssetID   string `json:"asset_id"`
}

type BCPartnerAssetDeleteResponse struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"request_id"`
	Data      BCPartnerAssetDeleteData   `json:"data"`
}

type BCPartnerAssetDeleteData struct {
	AssetID    string `json:"asset_id"`
	Status     string `json:"status"`
	DeleteTime string `json:"delete_time"`
}

type BCPixelLinkGetRequest struct {
	BCID    string `json:"bc_id"`
	PixelID string `json:"pixel_id,omitempty"`
}

type BCPixelLinkResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Data      BCPixelLinkData   `json:"data"`
}

type BCPixelLinkData struct {
	Links []BCPixelLink `json:"links"`
}

type BCPixelLink struct {
	PixelID      string `json:"pixel_id"`
	PixelName    string `json:"pixel_name"`
	LinkStatus   string `json:"link_status"`
	LinkTime     string `json:"link_time"`
	AdvertiserID string `json:"advertiser_id"`
}

type BCPixelLinkUpdateRequest struct {
	BCID         string `json:"bc_id"`
	PixelID      string `json:"pixel_id"`
	AdvertiserID string `json:"advertiser_id,omitempty"`
	LinkStatus   string `json:"link_status,omitempty"`
}

type BCPixelLinkUpdateResponse struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	RequestID string                  `json:"request_id"`
	Data      BCPixelLinkUpdateData   `json:"data"`
}

type BCPixelLinkUpdateData struct {
	PixelID    string `json:"pixel_id"`
	Status     string `json:"status"`
	UpdateTime string `json:"update_time"`
}

type BCPixelTransferRequest struct {
	BCID       string `json:"bc_id"`
	PixelID    string `json:"pixel_id"`
	TargetBCID string `json:"target_bc_id"`
}

type BCPixelTransferResponse struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	RequestID string                `json:"request_id"`
	Data      BCPixelTransferData   `json:"data"`
}

type BCPixelTransferData struct {
	PixelID      string `json:"pixel_id"`
	Status       string `json:"status"`
	TransferTime string `json:"transfer_time"`
}

// Missing BC types
type BCPartnerResponse struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      BCPartnerData   `json:"data"`
}

type BCPartnerData struct {
	Partners []BCPartner `json:"partners"`
}

type BCPartner struct {
	PartnerID   string `json:"partner_id"`
	PartnerName string `json:"partner_name"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	AssignTime  string `json:"assign_time"`
}

type BCPartnerAssetResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	RequestID string               `json:"request_id"`
	Data      BCPartnerAssetData   `json:"data"`
}

type BCPartnerAssetData struct {
	Assets []BCPartnerAsset `json:"assets"`
}

type BCPartnerAsset struct {
	AssetID    string `json:"asset_id"`
	AssetName  string `json:"asset_name"`
	AssetType  string `json:"asset_type"`
	Status     string `json:"status"`
	AssignTime string `json:"assign_time"`
}
