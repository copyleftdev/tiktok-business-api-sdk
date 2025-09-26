package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
)

// ValidateAdvertiserID validates an advertiser ID format
func ValidateAdvertiserID(advertiserID string) error {
	if advertiserID == "" {
		return models.NewValidationError("advertiser_id", "advertiser ID cannot be empty")
	}

	// TikTok advertiser IDs are typically numeric strings
	matched, err := regexp.MatchString(`^\d+$`, advertiserID)
	if err != nil {
		return fmt.Errorf("regex validation failed: %w", err)
	}

	if !matched {
		return models.NewValidationError("advertiser_id", "advertiser ID must be numeric")
	}

	return nil
}

// ValidateCampaignName validates a campaign name
func ValidateCampaignName(name string) error {
	if name == "" {
		return models.NewValidationError("campaign_name", "campaign name cannot be empty")
	}

	if len(name) > 512 {
		return models.NewValidationError("campaign_name", "campaign name cannot exceed 512 characters")
	}

	// Check for invalid characters (basic validation)
	if strings.Contains(name, "\n") || strings.Contains(name, "\r") {
		return models.NewValidationError("campaign_name", "campaign name cannot contain newline characters")
	}

	return nil
}

// ValidateBudget validates a budget amount
func ValidateBudget(budget float64, budgetMode models.BudgetMode) error {
	if budget <= 0 {
		return models.NewValidationError("budget", "budget must be greater than 0")
	}

	// Different minimum budgets based on budget mode
	var minBudget float64
	switch budgetMode {
	case models.BudgetModeDaily:
		minBudget = 20.0 // Minimum daily budget
	case models.BudgetModeTotal:
		minBudget = 50.0 // Minimum total budget
	default:
		return models.NewValidationError("budget_mode", "invalid budget mode")
	}

	if budget < minBudget {
		return models.NewValidationError("budget",
			fmt.Sprintf("budget must be at least %.2f for %s mode", minBudget, budgetMode))
	}

	return nil
}

// ValidateURL validates a URL format
func ValidateURL(urlStr string) error {
	if urlStr == "" {
		return models.NewValidationError("url", "URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return models.NewValidationError("url", fmt.Sprintf("invalid URL format: %v", err))
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return models.NewValidationError("url", "URL must use http or https scheme")
	}

	if parsedURL.Host == "" {
		return models.NewValidationError("url", "URL must have a valid host")
	}

	return nil
}

// ValidateAccessToken validates an access token format
func ValidateAccessToken(token string) error {
	if token == "" {
		return models.NewValidationError("access_token", "access token cannot be empty")
	}

	if len(token) < 10 {
		return models.NewValidationError("access_token", "access token appears to be too short")
	}

	// Basic format validation - TikTok tokens are typically alphanumeric with some special chars
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._-]+$`, token)
	if err != nil {
		return fmt.Errorf("regex validation failed: %w", err)
	}

	if !matched {
		return models.NewValidationError("access_token", "access token contains invalid characters")
	}

	return nil
}

// ValidateFields validates a list of field names
func ValidateFields(fields []string, allowedFields []string) error {
	if len(fields) == 0 {
		return nil // Empty fields list is usually allowed
	}

	allowedMap := make(map[string]bool)
	for _, field := range allowedFields {
		allowedMap[field] = true
	}

	var invalidFields []string
	for _, field := range fields {
		if !allowedMap[field] {
			invalidFields = append(invalidFields, field)
		}
	}

	if len(invalidFields) > 0 {
		return models.NewValidationError("fields",
			fmt.Sprintf("invalid fields: %s", strings.Join(invalidFields, ", ")))
	}

	return nil
}

// ValidateDateRange validates a date range
func ValidateDateRange(startDate, endDate string) error {
	if startDate == "" {
		return models.NewValidationError("start_date", "start date cannot be empty")
	}

	if endDate == "" {
		return models.NewValidationError("end_date", "end date cannot be empty")
	}

	// Validate date format (YYYY-MM-DD)
	dateRegex := `^\d{4}-\d{2}-\d{2}$`

	matched, err := regexp.MatchString(dateRegex, startDate)
	if err != nil {
		return fmt.Errorf("regex validation failed: %w", err)
	}
	if !matched {
		return models.NewValidationError("start_date", "start date must be in YYYY-MM-DD format")
	}

	matched, err = regexp.MatchString(dateRegex, endDate)
	if err != nil {
		return fmt.Errorf("regex validation failed: %w", err)
	}
	if !matched {
		return models.NewValidationError("end_date", "end date must be in YYYY-MM-DD format")
	}

	// Basic chronological validation
	if startDate > endDate {
		return models.NewValidationError("date_range", "start date must be before or equal to end date")
	}

	return nil
}

// ValidatePageSize validates pagination page size
func ValidatePageSize(pageSize int) error {
	if pageSize <= 0 {
		return models.NewValidationError("page_size", "page size must be greater than 0")
	}

	if pageSize > 1000 {
		return models.NewValidationError("page_size", "page size cannot exceed 1000")
	}

	return nil
}

// ValidateObjectiveType validates campaign objective type
func ValidateObjectiveType(objective models.ObjectiveType) error {
	validObjectives := []models.ObjectiveType{
		models.ObjectiveReach,
		models.ObjectiveTraffic,
		models.ObjectiveVideoViews,
		models.ObjectiveLeadGeneration,
		models.ObjectiveAppPromotion,
		models.ObjectiveConversions,
		models.ObjectiveProductSales,
		models.ObjectiveEngagement,
	}

	for _, valid := range validObjectives {
		if objective == valid {
			return nil
		}
	}

	return models.NewValidationError("objective_type",
		fmt.Sprintf("invalid objective type: %s", objective))
}

// ValidateRequiredString validates that a string field is not empty
func ValidateRequiredString(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return models.NewValidationError(fieldName, fmt.Sprintf("%s is required", fieldName))
	}
	return nil
}

// ValidateStringLength validates string length constraints
func ValidateStringLength(value, fieldName string, minLen, maxLen int) error {
	length := len(value)

	if length < minLen {
		return models.NewValidationError(fieldName,
			fmt.Sprintf("%s must be at least %d characters", fieldName, minLen))
	}

	if maxLen > 0 && length > maxLen {
		return models.NewValidationError(fieldName,
			fmt.Sprintf("%s cannot exceed %d characters", fieldName, maxLen))
	}

	return nil
}
