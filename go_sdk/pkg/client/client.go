package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
	"golang.org/x/time/rate"
)

// Client is the main TikTok Business API client
type Client struct {
	config      *Config
	httpClient  *http.Client
	rateLimiter *rate.Limiter
	baseURL     *url.URL

	// API services
	account        AccountService
	campaign       CampaignService
	ad             AdService
	adGroup        AdGroupService
	audience       AudienceService
	creative       CreativeService
	reporting      ReportingService
	tool           ToolService
	bc             BCService
	auth           AuthService
	businessCenter *BusinessCenterService
	catalog        *CatalogService
	dmp            *DMPService
	pixel          *PixelService
	optimizer      OptimizerService
	comment        CommentService
	report         ReportService
}

// NewClient creates a new TikTok Business API client
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	// Create HTTP client with timeout and secure transport configuration
	httpClient := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			// Explicit TLS configuration for security
			TLSHandshakeTimeout: 10 * time.Second,
			// Force TLS 1.2+ for security compliance
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				// Verify server certificates
				InsecureSkipVerify: false,
			},
		},
	}

	var rateLimiter *rate.Limiter
	if config.RateLimit != nil {
		rateLimiter = rate.NewLimiter(
			rate.Limit(config.RateLimit.RequestsPerSecond),
			config.RateLimit.BurstSize,
		)
	}

	client := &Client{
		config:      config,
		httpClient:  httpClient,
		rateLimiter: rateLimiter,
		baseURL:     baseURL,
	}

	// Initialize API services
	client.initServices()
	return client, nil
}

// initServices initializes all API services
func (c *Client) initServices() {
	// Initialize services with actual implementations
	c.account = &accountService{client: c}
	c.campaign = &campaignService{client: c}
	c.tool = &toolService{client: c}
	c.auth = &authService{client: c}

	// New expanded services
	c.businessCenter = NewBusinessCenterService(c)
	c.catalog = NewCatalogService(c)
	c.dmp = NewDMPService(c)
	c.pixel = NewPixelService(c)
	c.creative = NewCreativeService(c)
	c.optimizer = NewOptimizerService(c)
	c.comment = NewCommentService(c)
	c.report = NewReportService(c)

	// Services not yet implemented - return clear error messages
	c.ad = &notImplementedAdService{}
	c.adGroup = &notImplementedAdGroupService{}
	c.audience = &notImplementedAudienceService{}
	c.reporting = &notImplementedReportingService{}
	c.bc = &notImplementedBCService{}
}

// DoRequest performs an HTTP request with rate limiting and retry logic
func (c *Client) DoRequest(ctx context.Context, method, endpoint string, body io.Reader, headers map[string]string) (*http.Response, error) {
	// Apply rate limiting
	if c.rateLimiter != nil {
		if err := c.rateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limit error: %w", err)
		}
	}

	// Build full URL
	fullURL := c.baseURL.ResolveReference(&url.URL{Path: endpoint})

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header.Set("User-Agent", c.config.UserAgent)
	if c.config.AccessToken != "" {
		req.Header.Set("Access-Token", c.config.AccessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform request with retry logic
	maxRetries := 3
	if c.config.RetryConfig != nil {
		maxRetries = c.config.RetryConfig.MaxRetries
	}

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		// Check if we should retry based on status code
		if c.shouldRetry(resp.StatusCode) && attempt < maxRetries {
			_ = resp.Body.Close()
			lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("request failed after %d attempts: %w", maxRetries+1, lastErr)
}

// ParseResponse parses an HTTP response into the given interface
func (c *Client) ParseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for API errors
	if resp.StatusCode >= 400 {
		var apiErr models.APIError
		if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Code != "" {
			return &apiErr
		}
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// Parse successful response
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

// BuildURL builds a URL with query parameters
func (c *Client) BuildURL(endpoint string, params map[string]interface{}) string {
	u := c.baseURL.ResolveReference(&url.URL{Path: endpoint})

	if len(params) > 0 {
		q := u.Query()
		for key, value := range params {
			q.Set(key, fmt.Sprintf("%v", value))
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}

// shouldRetry determines if a request should be retried based on status code
func (c *Client) shouldRetry(statusCode int) bool {
	retryableCodes := []int{429, 500, 502, 503, 504}
	if c.config.RetryConfig != nil && len(c.config.RetryConfig.RetryableStatusCodes) > 0 {
		retryableCodes = c.config.RetryConfig.RetryableStatusCodes
	}

	for _, code := range retryableCodes {
		if statusCode == code {
			return true
		}
	}
	return false
}

// Account returns the Account API service
func (c *Client) Account() AccountService {
	return c.account
}

// Campaign returns the Campaign API service
func (c *Client) Campaign() CampaignService {
	return c.campaign
}

// Ad returns the Ad API service
func (c *Client) Ad() AdService {
	return c.ad
}

// AdGroup returns the AdGroup API service
func (c *Client) AdGroup() AdGroupService {
	return c.adGroup
}

// Audience returns the Audience API service
func (c *Client) Audience() AudienceService {
	return c.audience
}

// Creative returns the Creative API service
func (c *Client) Creative() CreativeService {
	return c.creative
}

// Reporting returns the Reporting API service
func (c *Client) Reporting() ReportingService {
	return c.reporting
}

// Tool returns the Tool API service
func (c *Client) Tool() ToolService {
	return c.tool
}

// BusinessCenter returns the Business Center API service
func (c *Client) BusinessCenter() *BusinessCenterService {
	return c.businessCenter
}

// Catalog returns the Catalog API service
func (c *Client) Catalog() *CatalogService {
	return c.catalog
}

// DMP returns the DMP (Data Management Platform) API service
func (c *Client) DMP() *DMPService {
	return c.dmp
}

// Pixel returns the Pixel API service
func (c *Client) Pixel() *PixelService {
	return c.pixel
}

// Optimizer returns the Optimizer API service
func (c *Client) Optimizer() OptimizerService {
	return c.optimizer
}

// Comment returns the Comment API service
func (c *Client) Comment() CommentService {
	return c.comment
}

// Report returns the Report API service
func (c *Client) Report() ReportService {
	return c.report
}

// BC returns the Business Center API service
func (c *Client) BC() BCService {
	return c.bc
}

// Auth returns the Authentication API service
func (c *Client) Auth() AuthService {
	return c.auth
}

// SetAccessToken sets the access token for authentication
func (c *Client) SetAccessToken(token string) {
	c.config.AccessToken = token
}

// SetTimeout sets the HTTP request timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.config.Timeout = timeout
	c.httpClient.Timeout = timeout
}

// BuildQueryParams builds query parameters from a map
func (c *Client) BuildQueryParams(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	values := url.Values{}
	for key, value := range params {
		if value != nil {
			values.Set(key, fmt.Sprintf("%v", value))
		}
	}

	return values.Encode()
}
