package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// AuthConfig holds authentication configuration
type AuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	BaseURL      string
}

// authService implements the AuthService interface
type authService struct {
	client *Client
	config *AuthConfig
}

// NewAuthService creates a new authentication service
func NewAuthService(config *AuthConfig) AuthService {
	if config.BaseURL == "" {
		config.BaseURL = "https://business-api.tiktok.com"
	}

	return &authService{
		config: config,
	}
}

// GetAuthorizationURL generates an OAuth authorization URL
func (a *authService) GetAuthorizationURL(scopes []string) string {
	baseURL := a.config.BaseURL + "/open_api/v1.3/oauth2/authorize/"

	params := url.Values{}
	params.Set("client_key", a.config.ClientID)
	params.Set("response_type", "code")
	params.Set("redirect_uri", a.config.RedirectURI)
	params.Set("state", "your_custom_params")

	if len(scopes) > 0 {
		params.Set("scope", strings.Join(scopes, ","))
	}

	return baseURL + "?" + params.Encode()
}

// GetAccessToken exchanges an authorization code for an access token
func (a *authService) GetAccessToken(ctx context.Context, code string) (*TokenResponse, error) {
	endpoint := "/open_api/v1.3/oauth2/access_token/"

	data := map[string]interface{}{
		"client_key":    a.config.ClientID,
		"client_secret": a.config.ClientSecret,
		"auth_code":     code,
		"grant_type":    "authorization_code",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := a.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var tokenResp struct {
		Code      int           `json:"code"`
		Message   string        `json:"message"`
		RequestID string        `json:"request_id"`
		Data      TokenResponse `json:"data"`
	}

	if err := a.client.ParseResponse(resp, &tokenResp); err != nil {
		return nil, err
	}

	if tokenResp.Code != 0 {
		return nil, fmt.Errorf("API error: %s", tokenResp.Message)
	}

	return &tokenResp.Data, nil
}

// RefreshToken refreshes an access token using a refresh token
func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	endpoint := "/open_api/v1.3/oauth2/refresh_token/"

	data := map[string]interface{}{
		"client_key":    a.config.ClientID,
		"client_secret": a.config.ClientSecret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := a.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var tokenResp struct {
		Code      int           `json:"code"`
		Message   string        `json:"message"`
		RequestID string        `json:"request_id"`
		Data      TokenResponse `json:"data"`
	}

	if err := a.client.ParseResponse(resp, &tokenResp); err != nil {
		return nil, err
	}

	if tokenResp.Code != 0 {
		return nil, fmt.Errorf("API error: %s", tokenResp.Message)
	}

	return &tokenResp.Data, nil
}

// ValidateToken validates an access token
func (a *authService) ValidateToken(ctx context.Context, token string) (*TokenValidationResponse, error) {
	endpoint := "/open_api/v1.3/oauth2/user_info/"

	headers := map[string]string{
		"Access-Token": token,
	}

	resp, err := a.client.DoRequest(ctx, "GET", endpoint, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var validationResp struct {
		Code      int                     `json:"code"`
		Message   string                  `json:"message"`
		RequestID string                  `json:"request_id"`
		Data      TokenValidationResponse `json:"data"`
	}

	if err := a.client.ParseResponse(resp, &validationResp); err != nil {
		return nil, err
	}

	if validationResp.Code != 0 {
		return &TokenValidationResponse{Valid: false}, nil
	}

	validationResp.Data.Valid = true
	return &validationResp.Data, nil
}

// RevokeToken revokes an access token
func (a *authService) RevokeToken(ctx context.Context, token string) error {
	endpoint := "/open_api/v1.3/oauth2/revoke/"

	data := map[string]interface{}{
		"client_key":    a.config.ClientID,
		"client_secret": a.config.ClientSecret,
		"access_token":  token,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := a.client.DoRequest(ctx, "POST", endpoint, strings.NewReader(string(body)), nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	var revokeResp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}

	if err := a.client.ParseResponse(resp, &revokeResp); err != nil {
		return err
	}

	if revokeResp.Code != 0 {
		return fmt.Errorf("API error: %s", revokeResp.Message)
	}

	return nil
}
