package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/client"
	"github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/models"
)

func main() {
	// Get access token from environment variable
	accessToken := os.Getenv("TIKTOK_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("TIKTOK_ACCESS_TOKEN environment variable is required")
	}

	// Create client configuration
	config := &client.Config{
		BaseURL:     "https://business-api.tiktok.com",
		AccessToken: accessToken,
		UserAgent:   "tiktok-go-sdk-example/1.0.0",
	}

	// Create the client
	tiktokClient, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Get advertiser information
	fmt.Println("=== Getting Advertiser Information ===")
	advertisers, err := tiktokClient.Account().GetAdvertisers(ctx, &client.GetAdvertisersRequest{
		Fields: []string{"advertiser_id", "advertiser_name", "status"},
	})
	if err != nil {
		log.Printf("Failed to get advertisers: %v", err)
	} else {
		fmt.Printf("Found %d advertisers\n", len(advertisers.Data))
		for _, advertiser := range advertisers.Data {
			fmt.Printf("- %s (ID: %s, Status: %s)\n",
				advertiser.AdvertiserName,
				advertiser.AdvertiserID,
				advertiser.Status)
		}
	}

	// Example 2: Create a campaign (requires valid advertiser ID and credentials)
	fmt.Println("\n=== Creating Campaign ===")
	campaign, err := tiktokClient.Campaign().Create(ctx, &client.CampaignCreateRequest{
		AdvertiserID:     "your_advertiser_id",
		CampaignName:     "Go SDK Test Campaign",
		ObjectiveType:    models.ObjectiveAppPromotion,
		Budget:           100.0,
		BudgetMode:       models.BudgetModeDaily,
		AppPromotionType: "APP_INSTALL",
	})
	if err != nil {
		fmt.Printf("Campaign creation failed (requires valid credentials): %v\n", err)
	} else {
		fmt.Printf("Campaign created: %+v\n", campaign)
	}

	// Example 3: Authentication flow demonstration
	fmt.Println("\n=== Authentication Example ===")
	authConfig := &client.AuthConfig{
		ClientID:     os.Getenv("TIKTOK_CLIENT_ID"),
		ClientSecret: os.Getenv("TIKTOK_CLIENT_SECRET"),
		RedirectURI:  "https://your-app.com/callback",
	}

	if authConfig.ClientID != "" && authConfig.ClientSecret != "" {
		authService := client.NewAuthService(authConfig)

		// Generate authorization URL
		scopes := []string{"user_info.basic", "video.list", "ad_management"}
		authURL := authService.GetAuthorizationURL(scopes)
		fmt.Printf("Authorization URL: %s\n", authURL)

		// Validate current token
		validation, err := authService.ValidateToken(ctx, accessToken)
		if err != nil {
			fmt.Printf("Token validation failed: %v\n", err)
		} else {
			fmt.Printf("Token validation result: Valid=%t, Scope=%s\n",
				validation.Valid, validation.Scope)
		}
	} else {
		fmt.Println("Set TIKTOK_CLIENT_ID and TIKTOK_CLIENT_SECRET for auth examples")
	}

	fmt.Println("\n=== Example completed ===")
}
