package models

import (
	"fmt"
	"net/http"
)

// APIError represents an error response from the TikTok Business API
type APIError struct {
	// Code is the error code returned by the API
	Code string `json:"code"`

	// Message is the human-readable error message
	Message string `json:"message"`

	// RequestID is the unique identifier for the request
	RequestID string `json:"request_id"`

	// Data contains additional error details
	Data interface{} `json:"data,omitempty"`

	// HTTPStatusCode is the HTTP status code of the response
	HTTPStatusCode int `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("TikTok API error [%s]: %s (request_id: %s)", e.Code, e.Message, e.RequestID)
	}
	return fmt.Sprintf("TikTok API error [%s]: %s", e.Code, e.Message)
}

// IsRetryable returns true if the error indicates a retryable condition
func (e *APIError) IsRetryable() bool {
	switch e.HTTPStatusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}

	// Check for specific error codes that indicate retryable conditions
	switch e.Code {
	case "RATE_LIMIT_EXCEEDED",
		"INTERNAL_ERROR",
		"SERVICE_UNAVAILABLE",
		"TIMEOUT":
		return true
	}

	return false
}

// IsAuthenticationError returns true if the error is related to authentication
func (e *APIError) IsAuthenticationError() bool {
	switch e.HTTPStatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return true
	}

	switch e.Code {
	case "UNAUTHORIZED",
		"FORBIDDEN",
		"INVALID_ACCESS_TOKEN",
		"ACCESS_TOKEN_EXPIRED",
		"INSUFFICIENT_PERMISSIONS":
		return true
	}

	return false
}

// IsValidationError returns true if the error is related to input validation
func (e *APIError) IsValidationError() bool {
	switch e.HTTPStatusCode {
	case http.StatusBadRequest:
		return true
	}

	switch e.Code {
	case "INVALID_PARAMETER",
		"MISSING_PARAMETER",
		"PARAMETER_VALUE_NOT_SUPPORTED",
		"VALIDATION_ERROR":
		return true
	}

	return false
}

// IsRateLimitError returns true if the error is related to rate limiting
func (e *APIError) IsRateLimitError() bool {
	return e.HTTPStatusCode == http.StatusTooManyRequests || e.Code == "RATE_LIMIT_EXCEEDED"
}

// ValidationError represents a client-side validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// NetworkError represents a network-related error
type NetworkError struct {
	Operation string
	Err       error
}

// Error implements the error interface
func (e NetworkError) Error() string {
	return fmt.Sprintf("network error during %s: %v", e.Operation, e.Err)
}

// Unwrap returns the underlying error
func (e NetworkError) Unwrap() error {
	return e.Err
}

// ConfigurationError represents a configuration-related error
type ConfigurationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e ConfigurationError) Error() string {
	return fmt.Sprintf("configuration error for field '%s': %s", e.Field, e.Message)
}

// Common error codes as constants
const (
	// Authentication errors
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeInvalidAccessToken = "INVALID_ACCESS_TOKEN"
	ErrCodeAccessTokenExpired = "ACCESS_TOKEN_EXPIRED"
	ErrCodeInsufficientPerms  = "INSUFFICIENT_PERMISSIONS"

	// Validation errors
	ErrCodeInvalidParameter      = "INVALID_PARAMETER"
	ErrCodeMissingParameter      = "MISSING_PARAMETER"
	ErrCodeParameterNotSupported = "PARAMETER_VALUE_NOT_SUPPORTED"
	ErrCodeValidationError       = "VALIDATION_ERROR"

	// Rate limiting errors
	ErrCodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"

	// Server errors
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	ErrCodeTimeout            = "TIMEOUT"

	// Resource errors
	ErrCodeResourceNotFound      = "RESOURCE_NOT_FOUND"
	ErrCodeResourceConflict      = "RESOURCE_CONFLICT"
	ErrCodeResourceLimitExceeded = "RESOURCE_LIMIT_EXCEEDED"

	// Business logic errors
	ErrCodeInsufficientBalance = "INSUFFICIENT_BALANCE"
	ErrCodeCampaignNotActive   = "CAMPAIGN_NOT_ACTIVE"
	ErrCodeAdGroupNotActive    = "ADGROUP_NOT_ACTIVE"
	ErrCodeCreativeNotApproved = "CREATIVE_NOT_APPROVED"
)

// NewAPIError creates a new APIError
func NewAPIError(code, message, requestID string, httpStatusCode int) *APIError {
	return &APIError{
		Code:           code,
		Message:        message,
		RequestID:      requestID,
		HTTPStatusCode: httpStatusCode,
	}
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// NewNetworkError creates a new NetworkError
func NewNetworkError(operation string, err error) NetworkError {
	return NetworkError{
		Operation: operation,
		Err:       err,
	}
}

// NewConfigurationError creates a new ConfigurationError
func NewConfigurationError(field, message string) ConfigurationError {
	return ConfigurationError{
		Field:   field,
		Message: message,
	}
}
