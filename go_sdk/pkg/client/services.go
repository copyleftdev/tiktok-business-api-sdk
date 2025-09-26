package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// accountService implements the AccountService interface
type accountService struct {
	client *Client
}

// GetAdvertisers retrieves a list of advertisers
func (a *accountService) GetAdvertisers(ctx context.Context, req *GetAdvertisersRequest) (*GetAdvertisersResponse, error) {
	endpoint := "/open_api/v1.3/advertiser/info/"

	params := map[string]interface{}{
		"advertiser_ids": fmt.Sprintf("[%s]", strings.Join(req.AdvertiserIDs, ",")),
	}

	if len(req.Fields) > 0 {
		params["fields"] = strings.Join(req.Fields, ",")
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}

	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	url := a.client.BuildURL(endpoint, params)

	resp, err := a.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get advertisers: %w", err)
	}

	var response GetAdvertisersResponse
	if err := a.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAdvertiserInfo retrieves information about a specific advertiser
func (a *accountService) GetAdvertiserInfo(ctx context.Context, advertiserID string) (*AdvertiserInfo, error) {
	req := &GetAdvertisersRequest{
		AdvertiserIDs: []string{advertiserID},
		Fields:        []string{"advertiser_id", "advertiser_name", "status", "currency", "timezone"},
	}

	resp, err := a.GetAdvertisers(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("advertiser not found: %s", advertiserID)
	}

	return &resp.Data[0], nil
}

// UpdateAdvertiser updates advertiser information
func (a *accountService) UpdateAdvertiser(ctx context.Context, req *UpdateAdvertiserRequest) (*UpdateAdvertiserResponse, error) {
	endpoint := "/open_api/v1.3/advertiser/update/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := a.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update advertiser: %w", err)
	}

	var response UpdateAdvertiserResponse
	if err := a.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateAdvertiser creates a new advertiser account
func (a *accountService) CreateAdvertiser(ctx context.Context, req *CreateAdvertiserRequest) (*CreateAdvertiserResponse, error) {
	endpoint := "/open_api/v1.3/advertiser/create/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := a.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create advertiser: %w", err)
	}

	var response CreateAdvertiserResponse
	if err := a.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAdvertiserBalance retrieves advertiser account balance
func (a *accountService) GetAdvertiserBalance(ctx context.Context, req *GetAdvertiserBalanceRequest) (*GetAdvertiserBalanceResponse, error) {
	endpoint := "/open_api/v1.3/advertiser/balance/get/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	url := a.client.BuildURL(endpoint, params)

	resp, err := a.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get advertiser balance: %w", err)
	}

	var response GetAdvertiserBalanceResponse
	if err := a.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAdvertiserFund retrieves advertiser fund information
func (a *accountService) GetAdvertiserFund(ctx context.Context, req *GetAdvertiserFundRequest) (*GetAdvertiserFundResponse, error) {
	endpoint := "/open_api/v1.3/advertiser/fund/get/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if len(req.FundTypes) > 0 {
		params["fund_types"] = fmt.Sprintf("[%s]", strings.Join(req.FundTypes, ","))
	}

	url := a.client.BuildURL(endpoint, params)

	resp, err := a.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get advertiser fund: %w", err)
	}

	var response GetAdvertiserFundResponse
	if err := a.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// campaignService implements the CampaignService interface
type campaignService struct {
	client *Client
}

// Create creates a new campaign
func (c *campaignService) Create(ctx context.Context, req *CampaignCreateRequest) (*CampaignCreateResponse, error) {
	endpoint := "/open_api/v1.3/campaign/create/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	var response CampaignCreateResponse
	if err := c.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Get retrieves campaign information
func (c *campaignService) Get(ctx context.Context, req *CampaignGetRequest) (*CampaignGetResponse, error) {
	endpoint := "/open_api/v1.3/campaign/get/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if len(req.CampaignIDs) > 0 {
		params["campaign_ids"] = fmt.Sprintf("[%s]", strings.Join(req.CampaignIDs, ","))
	}

	if len(req.Fields) > 0 {
		params["fields"] = strings.Join(req.Fields, ",")
	}

	if req.Page > 0 {
		params["page"] = req.Page
	}

	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	url := c.client.BuildURL(endpoint, params)

	resp, err := c.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get campaigns: %w", err)
	}

	var response CampaignGetResponse
	if err := c.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Update updates a campaign
func (c *campaignService) Update(ctx context.Context, req *CampaignUpdateRequest) (*CampaignUpdateResponse, error) {
	endpoint := "/open_api/v1.3/campaign/update/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update campaign: %w", err)
	}

	var response CampaignUpdateResponse
	if err := c.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Delete deletes campaigns
func (c *campaignService) Delete(ctx context.Context, req *CampaignDeleteRequest) (*CampaignDeleteResponse, error) {
	endpoint := "/open_api/v1.3/campaign/delete/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete campaigns: %w", err)
	}

	var response CampaignDeleteResponse
	if err := c.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateStatus updates campaign status
func (c *campaignService) UpdateStatus(ctx context.Context, req *CampaignStatusUpdateRequest) (*CampaignStatusUpdateResponse, error) {
	endpoint := "/open_api/v1.3/campaign/status/update/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update campaign status: %w", err)
	}

	var response CampaignStatusUpdateResponse
	if err := c.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// toolService implements the ToolService interface
type toolService struct {
	client *Client
}

// GetLanguages retrieves supported languages
func (t *toolService) GetLanguages(ctx context.Context, advertiserID string) (*LanguagesResponse, error) {
	endpoint := "/open_api/v1.3/tool/language/"

	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get languages: %w", err)
	}

	var response LanguagesResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCurrencies retrieves supported currencies
func (t *toolService) GetCurrencies(ctx context.Context, advertiserID string) (*CurrenciesResponse, error) {
	endpoint := "/open_api/v1.3/tool/currency/"

	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get currencies: %w", err)
	}

	var response CurrenciesResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetRegions retrieves supported regions
func (t *toolService) GetRegions(ctx context.Context, advertiserID string) (*RegionsResponse, error) {
	endpoint := "/open_api/v1.3/tool/region/"

	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get regions: %w", err)
	}

	var response RegionsResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetInterestCategories retrieves interest categories for targeting
func (t *toolService) GetInterestCategories(ctx context.Context, req *InterestCategoriesRequest) (*InterestCategoriesResponse, error) {
	endpoint := "/open_api/v1.3/tool/interest_category/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.Version != 0 {
		params["version"] = req.Version
	}

	if req.Language != "" {
		params["language"] = req.Language
	}

	if len(req.SpecialIndustries) > 0 {
		params["special_industries"] = fmt.Sprintf("[%s]", strings.Join(req.SpecialIndustries, ","))
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get interest categories: %w", err)
	}

	var response InterestCategoriesResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCarriers retrieves mobile carriers for targeting
func (t *toolService) GetCarriers(ctx context.Context, req *CarriersRequest) (*CarriersResponse, error) {
	endpoint := "/open_api/v1.3/tool/carrier/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if len(req.LocationIDs) > 0 {
		params["location_ids"] = fmt.Sprintf("[%s]", strings.Join(req.LocationIDs, ","))
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get carriers: %w", err)
	}

	var response CarriersResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetDeviceModels retrieves device models for targeting
func (t *toolService) GetDeviceModels(ctx context.Context, req *DeviceModelsRequest) (*DeviceModelsResponse, error) {
	endpoint := "/open_api/v1.3/tool/device_model/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.OSType != "" {
		params["os_type"] = req.OSType
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get device models: %w", err)
	}

	var response DeviceModelsResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetTargetingInfo retrieves targeting information by ID
func (t *toolService) GetTargetingInfo(ctx context.Context, req *TargetingInfoRequest) (*TargetingInfoResponse, error) {
	endpoint := "/open_api/v1.3/tool/targeting/info/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := t.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get targeting info: %w", err)
	}

	var response TargetingInfoResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBidRecommendation retrieves bid recommendations
func (t *toolService) GetBidRecommendation(ctx context.Context, req *BidRecommendRequest) (*BidRecommendResponse, error) {
	endpoint := "/open_api/v1.3/tool/bid/recommend/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := t.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get bid recommendation: %w", err)
	}

	var response BidRecommendResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetTargetingList retrieves targeting list
func (t *toolService) GetTargetingList(ctx context.Context, req *TargetingListRequest) (*TargetingListResponse, error) {
	endpoint := "/open_api/v1.3/tool/targeting/list/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"type":          req.Type,
	}

	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.Keyword != "" {
		params["keyword"] = req.Keyword
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.Size > 0 {
		params["size"] = req.Size
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get targeting list: %w", err)
	}

	var response TargetingListResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// SearchTargeting searches targeting options
func (t *toolService) SearchTargeting(ctx context.Context, req *TargetingSearchRequest) (*TargetingSearchResponse, error) {
	endpoint := "/open_api/v1.3/tool/targeting/search/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"type":          req.Type,
		"keyword":       req.Keyword,
	}

	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.CountryCode != "" {
		params["country_code"] = req.CountryCode
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to search targeting: %w", err)
	}

	var response TargetingSearchResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetOSVersions retrieves OS versions for targeting
func (t *toolService) GetOSVersions(ctx context.Context, req *OSVersionRequest) (*OSVersionResponse, error) {
	endpoint := "/open_api/v1.3/tool/os_version/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"os_type":       req.OSType,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get OS versions: %w", err)
	}

	var response OSVersionResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetTimezones retrieves supported timezones
func (t *toolService) GetTimezones(ctx context.Context, advertiserID string) (*TimezoneResponse, error) {
	endpoint := "/open_api/v1.3/tool/timezone/"

	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get timezones: %w", err)
	}

	var response TimezoneResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ValidateURL validates a URL
func (t *toolService) ValidateURL(ctx context.Context, req *URLValidateRequest) (*URLValidateResponse, error) {
	endpoint := "/open_api/v1.3/tool/url/validate/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := t.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate URL: %w", err)
	}

	var response URLValidateResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetHashtagRecommendations retrieves hashtag recommendations
func (t *toolService) GetHashtagRecommendations(ctx context.Context, req *HashtagRecommendRequest) (*HashtagRecommendResponse, error) {
	endpoint := "/open_api/v1.3/tool/hashtag/recommend/"

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := t.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get hashtag recommendations: %w", err)
	}

	var response HashtagRecommendResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetInterestKeywords retrieves interest keywords
func (t *toolService) GetInterestKeywords(ctx context.Context, req *InterestKeywordRequest) (*InterestKeywordResponse, error) {
	endpoint := "/open_api/v1.3/tool/interest_keyword/get/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"keyword":       req.Keyword,
	}

	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.CountryCode != "" {
		params["country_code"] = req.CountryCode
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get interest keywords: %w", err)
	}

	var response InterestKeywordResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetActionCategories retrieves action categories
func (t *toolService) GetActionCategories(ctx context.Context, req *ActionCategoryRequest) (*ActionCategoryResponse, error) {
	endpoint := "/open_api/v1.3/tool/action_category/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.SpecialIndustries != nil && len(req.SpecialIndustries) > 0 {
		params["special_industries"] = req.SpecialIndustries
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get action categories: %w", err)
	}

	var response ActionCategoryResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetContextualTags retrieves contextual tags
func (t *toolService) GetContextualTags(ctx context.Context, req *ContextualTagRequest) (*ContextualTagResponse, error) {
	endpoint := "/open_api/v1.3/tool/contextual_tag/get/"

	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.CountryCode != "" {
		params["country_code"] = req.CountryCode
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextual tags: %w", err)
	}

	var response ContextualTagResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPhoneRegionCodes retrieves phone region codes
func (t *toolService) GetPhoneRegionCodes(ctx context.Context, advertiserID string) (*PhoneRegionCodeResponse, error) {
	endpoint := "/open_api/v1.3/tool/phone_region_code/"

	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	url := t.client.BuildURL(endpoint, params)

	resp, err := t.client.DoRequest(ctx, "GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get phone region codes: %w", err)
	}

	var response PhoneRegionCodeResponse
	if err := t.client.ParseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
