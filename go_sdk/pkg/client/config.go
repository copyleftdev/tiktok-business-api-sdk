package client

import (
	"time"
)

// Config holds the configuration for the TikTok Business API client
type Config struct {
	// BaseURL is the base URL for the TikTok Business API
	BaseURL string

	// AccessToken is the OAuth 2.0 access token for authentication
	AccessToken string

	// ClientID is the OAuth 2.0 client ID
	ClientID string

	// ClientSecret is the OAuth 2.0 client secret
	ClientSecret string

	// Timeout is the HTTP request timeout
	Timeout time.Duration

	// RetryConfig configures retry behavior
	RetryConfig *RetryConfig

	// RateLimit configures rate limiting
	RateLimit *RateLimitConfig

	// UserAgent is the User-Agent header to send with requests
	UserAgent string

	// Debug enables debug logging
	Debug bool
}

// RetryConfig configures retry behavior for failed requests
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts
	MaxRetries int

	// BackoffStrategy defines the backoff strategy
	BackoffStrategy BackoffStrategy

	// InitialDelay is the initial delay before the first retry
	InitialDelay time.Duration

	// MaxDelay is the maximum delay between retries
	MaxDelay time.Duration

	// Multiplier is the backoff multiplier for exponential backoff
	Multiplier float64

	// RetryableStatusCodes defines which HTTP status codes should trigger a retry
	RetryableStatusCodes []int
}

// RateLimitConfig configures rate limiting
type RateLimitConfig struct {
	// RequestsPerSecond is the maximum number of requests per second
	RequestsPerSecond float64

	// BurstSize is the maximum number of requests that can be made in a burst
	BurstSize int
}

// BackoffStrategy defines the backoff strategy for retries
type BackoffStrategy int

const (
	// LinearBackoff increases delay linearly
	LinearBackoff BackoffStrategy = iota

	// ExponentialBackoff increases delay exponentially
	ExponentialBackoff

	// FixedBackoff uses a fixed delay
	FixedBackoff
)

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL:   "https://business-api.tiktok.com",
		Timeout:   30 * time.Second,
		UserAgent: "tiktok-business-api-go-sdk/1.0.0",
		RetryConfig: &RetryConfig{
			MaxRetries:      3,
			BackoffStrategy: ExponentialBackoff,
			InitialDelay:    1 * time.Second,
			MaxDelay:        30 * time.Second,
			Multiplier:      2.0,
			RetryableStatusCodes: []int{
				429, // Too Many Requests
				500, // Internal Server Error
				502, // Bad Gateway
				503, // Service Unavailable
				504, // Gateway Timeout
			},
		},
		RateLimit: &RateLimitConfig{
			RequestsPerSecond: 10,
			BurstSize:         20,
		},
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return ErrInvalidConfig{Field: "BaseURL", Message: "base URL is required"}
	}

	if c.AccessToken == "" && (c.ClientID == "" || c.ClientSecret == "") {
		return ErrInvalidConfig{
			Field:   "Authentication",
			Message: "either access token or client credentials are required",
		}
	}

	if c.Timeout <= 0 {
		return ErrInvalidConfig{Field: "Timeout", Message: "timeout must be positive"}
	}

	if c.RetryConfig != nil {
		if c.RetryConfig.MaxRetries < 0 {
			return ErrInvalidConfig{Field: "RetryConfig.MaxRetries", Message: "max retries cannot be negative"}
		}

		if c.RetryConfig.InitialDelay <= 0 {
			return ErrInvalidConfig{Field: "RetryConfig.InitialDelay", Message: "initial delay must be positive"}
		}

		if c.RetryConfig.MaxDelay <= 0 {
			return ErrInvalidConfig{Field: "RetryConfig.MaxDelay", Message: "max delay must be positive"}
		}

		if c.RetryConfig.Multiplier <= 0 {
			return ErrInvalidConfig{Field: "RetryConfig.Multiplier", Message: "multiplier must be positive"}
		}
	}

	if c.RateLimit != nil {
		if c.RateLimit.RequestsPerSecond <= 0 {
			return ErrInvalidConfig{Field: "RateLimit.RequestsPerSecond", Message: "requests per second must be positive"}
		}

		if c.RateLimit.BurstSize <= 0 {
			return ErrInvalidConfig{Field: "RateLimit.BurstSize", Message: "burst size must be positive"}
		}
	}

	return nil
}

// ErrInvalidConfig represents a configuration validation error
type ErrInvalidConfig struct {
	Field   string
	Message string
}

func (e ErrInvalidConfig) Error() string {
	return "invalid config field '" + e.Field + "': " + e.Message
}
