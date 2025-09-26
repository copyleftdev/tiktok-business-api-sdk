package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/client"
	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/utils"
)

func main() {
	// Get access token from environment variable
	accessToken := os.Getenv("TIKTOK_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("TIKTOK_ACCESS_TOKEN environment variable is required")
	}

	advertiserID := os.Getenv("TIKTOK_ADVERTISER_ID")
	if advertiserID == "" {
		log.Fatal("TIKTOK_ADVERTISER_ID environment variable is required")
	}

	// Create client configuration with custom settings
	config := &client.Config{
		BaseURL:     "https://business-api.tiktok.com",
		AccessToken: accessToken,
		UserAgent:   "tiktok-go-sdk-advanced-example/1.0.0",
		Timeout:     30 * time.Second, // 30 seconds
		RetryConfig: &client.RetryConfig{
			MaxRetries:      3,
			BackoffStrategy: client.ExponentialBackoff,
		},
		RateLimit: &client.RateLimitConfig{
			RequestsPerSecond: 5,
			BurstSize:         10,
		},
	}

	// Create the client
	tiktokClient, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Comprehensive Account Management
	fmt.Println("=== Advanced Account Management ===")

	// Get detailed advertiser information
	advertisers, err := tiktokClient.Account().GetAdvertisers(ctx, &client.GetAdvertisersRequest{
		Fields: []string{
			"advertiser_id", "advertiser_name", "status", "currency",
			"timezone", "company_name", "industry", "balance", "create_time",
		},
		AdvertiserIDs: []string{advertiserID},
		Page:          1,
		PageSize:      10,
	})
	if err != nil {
		log.Printf("Failed to get advertisers: %v", err)
	} else {
		fmt.Printf("Found %d advertisers\n", len(advertisers.Data))
		for _, advertiser := range advertisers.Data {
			fmt.Printf("- %s (ID: %s)\n", advertiser.AdvertiserName, advertiser.AdvertiserID)
			fmt.Printf("  Company: %s, Industry: %s\n", advertiser.CompanyName, advertiser.Industry)
			fmt.Printf("  Currency: %s, Balance: %.2f\n", advertiser.Currency, advertiser.Balance)
		}
	}

	// Get account balance
	balance, err := tiktokClient.Account().GetAdvertiserBalance(ctx, &client.GetAdvertiserBalanceRequest{
		AdvertiserID: advertiserID,
	})
	if err != nil {
		log.Printf("Failed to get balance: %v", err)
	} else {
		fmt.Printf("Account Balance: %.2f %s\n", balance.Data.Balance, balance.Data.Currency)
	}

	// Example 2: Campaign Management with Validation
	fmt.Println("\n=== Advanced Campaign Management ===")

	// Validate campaign parameters before creating
	campaignName := "Go SDK Advanced Test Campaign"
	budget := 100.0
	budgetMode := models.BudgetModeDaily

	if err := utils.ValidateCampaignName(campaignName); err != nil {
		log.Printf("Campaign name validation failed: %v", err)
		return
	}

	if err := utils.ValidateBudget(budget, budgetMode); err != nil {
		log.Printf("Budget validation failed: %v", err)
		return
	}

	if err := utils.ValidateAdvertiserID(advertiserID); err != nil {
		log.Printf("Advertiser ID validation failed: %v", err)
		return
	}

	// Create campaign with comprehensive parameters
	campaign, err := tiktokClient.Campaign().Create(ctx, &client.CampaignCreateRequest{
		AdvertiserID:      advertiserID,
		CampaignName:      campaignName,
		ObjectiveType:     models.ObjectiveAppPromotion,
		Budget:            budget,
		BudgetMode:        budgetMode,
		AppPromotionType:  "APP_INSTALL",
		CampaignType:      "NORMAL_CAMPAIGN",
		SpecialIndustries: []string{},
	})
	if err != nil {
		fmt.Printf("Campaign creation failed (requires valid credentials): %v\n", err)
	} else {
		fmt.Printf("Campaign created successfully: %+v\n", campaign)

		// Get campaign details
		campaigns, err := tiktokClient.Campaign().Get(ctx, &client.CampaignGetRequest{
			AdvertiserID: advertiserID,
			CampaignIDs:  []string{campaign.Data.CampaignID},
			Fields: []string{
				"campaign_id", "campaign_name", "status", "objective_type",
				"budget", "budget_mode", "create_time", "modify_time",
			},
		})
		if err != nil {
			log.Printf("Failed to get campaign details: %v", err)
		} else {
			fmt.Printf("Campaign details retrieved: %+v\n", campaigns.Data)
		}
	}

	// Example 3: Tool API Usage
	fmt.Println("\n=== Tool API Usage ===")

	// Get supported languages
	languages, err := tiktokClient.Tool().GetLanguages(ctx, advertiserID)
	if err != nil {
		fmt.Printf("Failed to get languages (requires valid credentials): %v\n", err)
	} else {
		fmt.Printf("Supported languages: %d found\n", len(languages.Data))
		for i, lang := range languages.Data {
			if i < 5 { // Show first 5
				fmt.Printf("- %s (%s)\n", lang.LanguageName, lang.LanguageCode)
			}
		}
	}

	// Get supported currencies
	currencies, err := tiktokClient.Tool().GetCurrencies(ctx, advertiserID)
	if err != nil {
		fmt.Printf("Failed to get currencies (requires valid credentials): %v\n", err)
	} else {
		fmt.Printf("Supported currencies: %d found\n", len(currencies.Data))
		for i, curr := range currencies.Data {
			if i < 5 { // Show first 5
				fmt.Printf("- %s (%s)\n", curr.CurrencyName, curr.CurrencyCode)
			}
		}
	}

	// Get interest categories for targeting
	interests, err := tiktokClient.Tool().GetInterestCategories(ctx, &client.InterestCategoriesRequest{
		AdvertiserID: advertiserID,
		Version:      2,
		Language:     "en",
	})
	if err != nil {
		fmt.Printf("Failed to get interest categories (requires valid credentials): %v\n", err)
	} else {
		fmt.Printf("Interest categories: %d found\n", len(interests.Data))
		for i, interest := range interests.Data {
			if i < 3 { // Show first 3
				fmt.Printf("- %s (Level %d)\n", interest.InterestCategoryName, interest.Level)
			}
		}
	}

	// Example 4: Validation Utilities
	fmt.Println("\n=== Validation Examples ===")

	// Test URL validation
	testURL := "https://example.com/callback"
	if err := utils.ValidateURL(testURL); err != nil {
		fmt.Printf("URL validation failed: %v\n", err)
	} else {
		fmt.Printf("URL validation passed: %s\n", testURL)
	}

	// Test date range validation
	startDate := "2024-01-01"
	endDate := "2024-01-31"
	if err := utils.ValidateDateRange(startDate, endDate); err != nil {
		fmt.Printf("Date range validation failed: %v\n", err)
	} else {
		fmt.Printf("Date range validation passed: %s to %s\n", startDate, endDate)
	}

	// Test objective type validation
	if err := utils.ValidateObjectiveType(models.ObjectiveAppPromotion); err != nil {
		fmt.Printf("Objective type validation failed: %v\n", err)
	} else {
		fmt.Printf("Objective type validation passed: %s\n", models.ObjectiveAppPromotion)
	}

	// Example 5: Pagination Helper
	fmt.Println("\n=== Pagination Helper Example ===")

	helper := utils.NewPaginationHelper()

	// Calculate pagination info
	totalCount := 150
	pageSize := 20
	currentPage := 3

	offset := helper.CalculateOffset(currentPage, pageSize)
	totalPages := helper.CalculateTotalPages(totalCount, pageSize)
	pageInfo := helper.CreatePaginationInfo(currentPage, pageSize, totalCount)

	fmt.Printf("Pagination Info:\n")
	fmt.Printf("- Total Count: %d\n", totalCount)
	fmt.Printf("- Page Size: %d\n", pageSize)
	fmt.Printf("- Current Page: %d\n", currentPage)
	fmt.Printf("- Offset: %d\n", offset)
	fmt.Printf("- Total Pages: %d\n", totalPages)
	fmt.Printf("- Has More: %t\n", pageInfo.HasMore)

	// Example 6: Error Handling
	fmt.Println("\n=== Error Handling Examples ===")

	// Demonstrate different error types
	validationErr := models.NewValidationError("test_field", "test validation message")
	fmt.Printf("Validation Error: %v\n", validationErr)

	networkErr := models.NewNetworkError("test_operation", fmt.Errorf("connection failed"))
	fmt.Printf("Network Error: %v\n", networkErr)

	configErr := models.NewConfigurationError("test_config", "invalid configuration")
	fmt.Printf("Configuration Error: %v\n", configErr)

	fmt.Println("\n=== Advanced Example Completed ===")
	fmt.Println("This example demonstrates:")
	fmt.Println("- Advanced client configuration")
	fmt.Println("- Comprehensive API usage")
	fmt.Println("- Input validation")
	fmt.Println("- Pagination helpers")
	fmt.Println("- Error handling patterns")
	fmt.Println("- Best practices for Go SDK usage")
}
