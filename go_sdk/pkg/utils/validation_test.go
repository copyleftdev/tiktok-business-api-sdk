package utils

import (
	"testing"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
)

func TestValidateAdvertiserID(t *testing.T) {
	tests := []struct {
		name         string
		advertiserID string
		expectError  bool
	}{
		{
			name:         "valid numeric ID",
			advertiserID: "123456789",
			expectError:  false,
		},
		{
			name:         "empty ID",
			advertiserID: "",
			expectError:  true,
		},
		{
			name:         "non-numeric ID",
			advertiserID: "abc123",
			expectError:  true,
		},
		{
			name:         "ID with spaces",
			advertiserID: "123 456",
			expectError:  true,
		},
		{
			name:         "very long numeric ID",
			advertiserID: "12345678901234567890",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAdvertiserID(tt.advertiserID)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateAdvertiserID() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateCampaignName(t *testing.T) {
	tests := []struct {
		name         string
		campaignName string
		expectError  bool
	}{
		{
			name:         "valid campaign name",
			campaignName: "Test Campaign",
			expectError:  false,
		},
		{
			name:         "empty campaign name",
			campaignName: "",
			expectError:  true,
		},
		{
			name:         "campaign name with newline",
			campaignName: "Test\nCampaign",
			expectError:  true,
		},
		{
			name:         "campaign name with carriage return",
			campaignName: "Test\rCampaign",
			expectError:  true,
		},
		{
			name:         "very long campaign name",
			campaignName: string(make([]byte, 600)), // 600 characters
			expectError:  true,
		},
		{
			name:         "max length campaign name",
			campaignName: string(make([]byte, 512)), // 512 characters
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCampaignName(tt.campaignName)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateCampaignName() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateBudget(t *testing.T) {
	tests := []struct {
		name        string
		budget      float64
		budgetMode  models.BudgetMode
		expectError bool
	}{
		{
			name:        "valid daily budget",
			budget:      50.0,
			budgetMode:  models.BudgetModeDaily,
			expectError: false,
		},
		{
			name:        "valid total budget",
			budget:      100.0,
			budgetMode:  models.BudgetModeTotal,
			expectError: false,
		},
		{
			name:        "zero budget",
			budget:      0.0,
			budgetMode:  models.BudgetModeDaily,
			expectError: true,
		},
		{
			name:        "negative budget",
			budget:      -10.0,
			budgetMode:  models.BudgetModeDaily,
			expectError: true,
		},
		{
			name:        "daily budget too low",
			budget:      10.0,
			budgetMode:  models.BudgetModeDaily,
			expectError: true,
		},
		{
			name:        "total budget too low",
			budget:      30.0,
			budgetMode:  models.BudgetModeTotal,
			expectError: true,
		},
		{
			name:        "invalid budget mode",
			budget:      50.0,
			budgetMode:  "INVALID_MODE",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBudget(tt.budget, tt.budgetMode)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateBudget() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "valid HTTPS URL",
			url:         "https://example.com",
			expectError: false,
		},
		{
			name:        "valid HTTP URL",
			url:         "http://example.com",
			expectError: false,
		},
		{
			name:        "empty URL",
			url:         "",
			expectError: true,
		},
		{
			name:        "invalid scheme",
			url:         "ftp://example.com",
			expectError: true,
		},
		{
			name:        "no host",
			url:         "https://",
			expectError: true,
		},
		{
			name:        "malformed URL",
			url:         "not-a-url",
			expectError: true,
		},
		{
			name:        "URL with path",
			url:         "https://example.com/path/to/resource",
			expectError: false,
		},
		{
			name:        "URL with query parameters",
			url:         "https://example.com?param=value",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.url)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateURL() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateAccessToken(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		expectError bool
	}{
		{
			name:        "valid token",
			token:       "abcdef123456789",
			expectError: false,
		},
		{
			name:        "empty token",
			token:       "",
			expectError: true,
		},
		{
			name:        "too short token",
			token:       "abc",
			expectError: true,
		},
		{
			name:        "token with invalid characters",
			token:       "token with spaces",
			expectError: true,
		},
		{
			name:        "token with special chars",
			token:       "token_with-dots.and_dashes",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAccessToken(tt.token)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateAccessToken() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateFields(t *testing.T) {
	allowedFields := []string{"field1", "field2", "field3"}

	tests := []struct {
		name        string
		fields      []string
		expectError bool
	}{
		{
			name:        "valid fields",
			fields:      []string{"field1", "field2"},
			expectError: false,
		},
		{
			name:        "empty fields",
			fields:      []string{},
			expectError: false,
		},
		{
			name:        "nil fields",
			fields:      nil,
			expectError: false,
		},
		{
			name:        "invalid field",
			fields:      []string{"field1", "invalid_field"},
			expectError: true,
		},
		{
			name:        "all invalid fields",
			fields:      []string{"invalid1", "invalid2"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFields(tt.fields, allowedFields)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateFields() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateDateRange(t *testing.T) {
	tests := []struct {
		name        string
		startDate   string
		endDate     string
		expectError bool
	}{
		{
			name:        "valid date range",
			startDate:   "2024-01-01",
			endDate:     "2024-01-31",
			expectError: false,
		},
		{
			name:        "same start and end date",
			startDate:   "2024-01-01",
			endDate:     "2024-01-01",
			expectError: false,
		},
		{
			name:        "empty start date",
			startDate:   "",
			endDate:     "2024-01-31",
			expectError: true,
		},
		{
			name:        "empty end date",
			startDate:   "2024-01-01",
			endDate:     "",
			expectError: true,
		},
		{
			name:        "invalid start date format",
			startDate:   "01-01-2024",
			endDate:     "2024-01-31",
			expectError: true,
		},
		{
			name:        "invalid end date format",
			startDate:   "2024-01-01",
			endDate:     "01-31-2024",
			expectError: true,
		},
		{
			name:        "start date after end date",
			startDate:   "2024-01-31",
			endDate:     "2024-01-01",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDateRange(tt.startDate, tt.endDate)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateDateRange() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidatePageSize(t *testing.T) {
	tests := []struct {
		name        string
		pageSize    int
		expectError bool
	}{
		{
			name:        "valid page size",
			pageSize:    50,
			expectError: false,
		},
		{
			name:        "minimum page size",
			pageSize:    1,
			expectError: false,
		},
		{
			name:        "maximum page size",
			pageSize:    1000,
			expectError: false,
		},
		{
			name:        "zero page size",
			pageSize:    0,
			expectError: true,
		},
		{
			name:        "negative page size",
			pageSize:    -1,
			expectError: true,
		},
		{
			name:        "page size too large",
			pageSize:    1001,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePageSize(tt.pageSize)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidatePageSize() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateObjectiveType(t *testing.T) {
	tests := []struct {
		name        string
		objective   models.ObjectiveType
		expectError bool
	}{
		{
			name:        "valid objective - reach",
			objective:   models.ObjectiveReach,
			expectError: false,
		},
		{
			name:        "valid objective - app promotion",
			objective:   models.ObjectiveAppPromotion,
			expectError: false,
		},
		{
			name:        "valid objective - conversions",
			objective:   models.ObjectiveConversions,
			expectError: false,
		},
		{
			name:        "invalid objective",
			objective:   "INVALID_OBJECTIVE",
			expectError: true,
		},
		{
			name:        "empty objective",
			objective:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateObjectiveType(tt.objective)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateObjectiveType() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateRequiredString(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		fieldName   string
		expectError bool
	}{
		{
			name:        "valid string",
			value:       "test value",
			fieldName:   "test_field",
			expectError: false,
		},
		{
			name:        "empty string",
			value:       "",
			fieldName:   "test_field",
			expectError: true,
		},
		{
			name:        "whitespace only",
			value:       "   ",
			fieldName:   "test_field",
			expectError: true,
		},
		{
			name:        "string with content and whitespace",
			value:       "  test  ",
			fieldName:   "test_field",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequiredString(tt.value, tt.fieldName)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateRequiredString() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestValidateStringLength(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		fieldName   string
		minLen      int
		maxLen      int
		expectError bool
	}{
		{
			name:        "valid length",
			value:       "test",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      10,
			expectError: false,
		},
		{
			name:        "too short",
			value:       "a",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      10,
			expectError: true,
		},
		{
			name:        "too long",
			value:       "this is too long",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      10,
			expectError: true,
		},
		{
			name:        "no max length limit",
			value:       "this is a very long string that should be valid",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      0, // 0 means no limit
			expectError: false,
		},
		{
			name:        "minimum length boundary",
			value:       "ab",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      10,
			expectError: false,
		},
		{
			name:        "maximum length boundary",
			value:       "1234567890",
			fieldName:   "test_field",
			minLen:      2,
			maxLen:      10,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStringLength(tt.value, tt.fieldName, tt.minLen, tt.maxLen)
			if (err != nil) != tt.expectError {
				t.Errorf("ValidateStringLength() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}
