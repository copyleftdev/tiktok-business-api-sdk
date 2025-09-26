# TikTok Business API SDK - Go

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

The TikTok Business API SDK for Go provides a comprehensive and idiomatic Go interface to the [TikTok API for Business](https://ads.tiktok.com/marketing_api/). This SDK enables developers to integrate with TikTok's advertising platform efficiently and follows Go best practices.

This Go SDK complements the existing [Java](../java_sdk/), [Python](../python_sdk/), and [JavaScript](../js_sdk/) SDKs. While the other SDKs are generated using [Swagger Codegen](https://github.com/swagger-api/swagger-codegen), this Go SDK is hand-crafted to provide Go-idiomatic APIs with enhanced features like built-in rate limiting, retry logic, and comprehensive error handling.

## Features

### 🚀 **Premium Go Enhancements**
- **Built-in Rate Limiting**: Automatic protection against API rate limits with configurable strategies
- **Intelligent Retry Logic**: Exponential backoff with customizable retry policies
- **Context-First Design**: Full `context.Context` support for cancellation, timeouts, and tracing
- **Type-Safe Error Handling**: Comprehensive error categorization with actionable error information
- **Input Validation**: Client-side validation to catch errors before API calls
- **Memory Efficient**: Optimized for high-throughput applications with connection pooling

### 📋 **Standard SDK Features**
- **Complete API Coverage**: Full support for all TikTok Business API endpoints
- **Type Safety**: Strong typing for all API models and responses  
- **Authentication**: Built-in OAuth 2.0 support with token management
- **Comprehensive Testing**: Extensive test suite with HTTP mocking
- **Rich Documentation**: Complete API documentation and usage examples
- **Swagger Compatible**: Maintains compatibility with TikTok's generation strategy

## Why Choose the Premium Go SDK?

While TikTok's other SDKs are generated using Swagger Codegen for consistency, this Go SDK is **hand-crafted** to leverage Go's unique strengths:

| Feature | Generated SDK | **Premium Go SDK** |
|---------|---------------|-------------------|
| **Rate Limiting** | ❌ Manual implementation | ✅ **Built-in with configurable strategies** |
| **Retry Logic** | ❌ Manual implementation | ✅ **Intelligent exponential backoff** |
| **Error Handling** | ⚠️ Basic HTTP errors | ✅ **Categorized, actionable errors** |
| **Context Support** | ❌ Limited | ✅ **Full context.Context integration** |
| **Memory Efficiency** | ⚠️ Standard | ✅ **Optimized for high-throughput** |
| **Input Validation** | ❌ Server-side only | ✅ **Client-side validation** |
| **Go Idioms** | ❌ Generic patterns | ✅ **Native Go conventions** |

### 🎯 **Perfect for Production Go Applications**
- **High-throughput advertising platforms**
- **Microservices requiring reliable API integration**  
- **Applications needing advanced error recovery**
- **Systems requiring precise rate limit management**

## Installation

```bash
go get github.com/tiktok/tiktok-business-api-sdk/go_sdk
```

## Quick Start

### Prerequisites

1. [Create a TikTok For Business account](https://ads.tiktok.com/marketing_api/docs?id=1738855099573250)
2. [Register as a developer](https://ads.tiktok.com/marketing_api/docs?id=1738855176671234)
3. [Create a developer app](https://ads.tiktok.com/marketing_api/docs?id=1738855242728450)
4. [Obtain authorization](https://ads.tiktok.com/marketing_api/docs?id=1738373141733378)

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    tiktok "github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/client"
)

func main() {
    // Initialize the client
    client := tiktok.NewClient(&tiktok.Config{
        AccessToken: "your_access_token",
        BaseURL:     "https://business-api.tiktok.com",
    })

    // Get advertiser information
    ctx := context.Background()
    advertisers, err := client.Account().GetAdvertisers(ctx, &tiktok.GetAdvertisersRequest{
        Fields: []string{"advertiser_id", "advertiser_name", "status"},
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, advertiser := range advertisers.Data {
        fmt.Printf("Advertiser: %s (ID: %s)\n", advertiser.AdvertiserName, advertiser.AdvertiserID)
    }
}
```

### Authentication

```go
package main

import (
    "context"
    "fmt"
    "log"

    tiktok "github.com/tiktok/tiktok-business-api-sdk/go_sdk/pkg/client"
)

func main() {
    // OAuth 2.0 authentication
    auth := tiktok.NewAuthService(&tiktok.AuthConfig{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
        RedirectURI:  "your_redirect_uri",
    })

    // Get authorization URL
    authURL := auth.GetAuthorizationURL([]string{"user_info.basic", "video.list"})
    fmt.Printf("Visit this URL to authorize: %s\n", authURL)

    // Exchange authorization code for access token
    ctx := context.Background()
    token, err := auth.GetAccessToken(ctx, "authorization_code_from_callback")
    if err != nil {
        log.Fatal(err)
    }

    // Use the token with the client
    client := tiktok.NewClient(&tiktok.Config{
        AccessToken: token.AccessToken,
    })
}
```

## API Coverage

The Go SDK provides complete coverage of the TikTok Business API:

### Core APIs
- **Account Management**: Advertiser accounts, business centers
- **Campaign Management**: Campaigns, ad groups, ads
- **Creative Management**: Creative assets, videos, images
- **Audience Management**: Custom audiences, lookalike audiences
- **Reporting**: Performance reports, analytics
- **Tools**: Languages, currencies, targeting options

### Specialized APIs
- **ACO (Automated Creative Optimization)**: Smart ad creation
- **Catalog Management**: Product catalogs for e-commerce
- **Measurement**: Conversion tracking, attribution
- **Business Center**: Multi-account management
- **Comments**: Comment management and moderation

## Examples

### Campaign Creation

```go
campaign, err := client.Campaign().Create(ctx, &tiktok.CampaignCreateRequest{
    AdvertiserID:   "your_advertiser_id",
    CampaignName:   "My Go SDK Campaign",
    ObjectiveType:  "APP_PROMOTION",
    Budget:         50.0,
    BudgetMode:     "BUDGET_MODE_TOTAL",
})
```

### Reporting

```go
report, err := client.Reporting().GetBasicReports(ctx, &tiktok.ReportingRequest{
    AdvertiserID: "your_advertiser_id",
    ReportType:   "BASIC",
    DataLevel:    "AUCTION_CAMPAIGN",
    Dimensions:   []string{"campaign_id", "stat_time_day"},
    Metrics:      []string{"impressions", "clicks", "cost"},
    StartDate:    "2023-01-01",
    EndDate:      "2023-01-31",
})
```

### File Upload

```go
file, err := client.File().Upload(ctx, &tiktok.FileUploadRequest{
    AdvertiserID: "your_advertiser_id",
    FileType:     "IMAGE",
    FileName:     "creative.jpg",
    FileContent:  fileBytes,
})
```

## Configuration

### Client Configuration

```go
config := &tiktok.Config{
    BaseURL:     "https://business-api.tiktok.com",
    AccessToken: "your_access_token",
    Timeout:     30 * time.Second,
    RetryConfig: &tiktok.RetryConfig{
        MaxRetries: 3,
        BackoffStrategy: tiktok.ExponentialBackoff,
    },
    RateLimit: &tiktok.RateLimitConfig{
        RequestsPerSecond: 10,
        BurstSize:        20,
    },
}

client := tiktok.NewClient(config)
```

### Environment Variables

The SDK supports configuration via environment variables:

```bash
export TIKTOK_ACCESS_TOKEN="your_access_token"
export TIKTOK_BASE_URL="https://business-api.tiktok.com"
export TIKTOK_CLIENT_ID="your_client_id"
export TIKTOK_CLIENT_SECRET="your_client_secret"
```

## Error Handling

The SDK provides comprehensive error handling:

```go
campaign, err := client.Campaign().Create(ctx, request)
if err != nil {
    var apiErr *tiktok.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
        fmt.Printf("Request ID: %s\n", apiErr.RequestID)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests (requires API credentials)
go test -tags=integration ./tests/integration/...
```

### Mock Testing

```go
func TestCampaignCreate(t *testing.T) {
    mockClient := tiktok.NewMockClient()
    mockClient.Campaign().EXPECT().
        Create(gomock.Any(), gomock.Any()).
        Return(&tiktok.Campaign{ID: "123"}, nil)

    // Test your code with mockClient
}
```

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/tiktok/tiktok-business-api-sdk.git
cd tiktok-business-api-sdk/go_sdk

# Install dependencies
go mod download

# Run tests
go test ./...

# Run linter
golangci-lint run
```

## Documentation

- [API Reference](docs/api_reference.md)
- [Authentication Guide](docs/authentication.md)
- [Examples](docs/examples.md)
- [Migration Guide](docs/migration_guide.md)

## Support

- [TikTok API for Business Developer Portal](https://ads.tiktok.com/marketing_api/homepage)
- [GitHub Issues](https://github.com/tiktok/tiktok-business-api-sdk/issues)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes and version history.
