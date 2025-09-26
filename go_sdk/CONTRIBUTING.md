# Contributing to TikTok Business API Go SDK

Thank you for your interest in contributing to the TikTok Business API Go SDK! This document provides guidelines for contributing to this project.

## Development Setup

### Prerequisites
- Go 1.19 or later
- Git

### Getting Started
1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/tiktok-business-api-sdk.git
   cd tiktok-business-api-sdk/go_sdk
   ```
3. Install dependencies:
   ```bash
   make dev-setup
   ```

## Development Workflow

### Running Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests with race detection
make test-race
```

### Code Quality
```bash
# Format code
make fmt

# Run linter
make lint

# Run all quality checks
make check
```

### Building
```bash
# Build the SDK
make build
```

## Code Guidelines

### Go Standards
- Follow standard Go conventions and idioms
- Use `gofmt` for code formatting
- Write comprehensive tests for new functionality
- Include godoc comments for all public APIs
- Handle errors appropriately using Go's error handling patterns

### API Implementation Patterns
When implementing new API endpoints, follow these patterns:

1. **Define types in `pkg/client/types.go`**:
   ```go
   type NewAPIRequest struct {
       AdvertiserID string `json:"advertiser_id"`
       // ... other fields
   }
   
   type NewAPIResponse struct {
       models.BaseResponse
       Data []NewAPIData `json:"data"`
   }
   ```

2. **Add interface method in `pkg/client/interfaces.go`**:
   ```go
   type NewAPIService interface {
       NewMethod(ctx context.Context, req *NewAPIRequest) (*NewAPIResponse, error)
   }
   ```

3. **Implement in `pkg/api/newapi.go`**:
   ```go
   func (s *newAPIService) NewMethod(ctx context.Context, req *NewAPIRequest) (*NewAPIResponse, error) {
       // Implementation following existing patterns
   }
   ```

4. **Add tests in `pkg/api/newapi_test.go`**:
   ```go
   func TestNewAPIService_NewMethod(t *testing.T) {
       // Test implementation with HTTP mocking
   }
   ```

### Testing Guidelines
- Write unit tests for all new functionality
- Use HTTP mocking for API integration tests
- Test error scenarios and edge cases
- Maintain test coverage above 80%

### Documentation
- Include godoc comments for all public functions and types
- Update README.md for new features
- Add usage examples for new APIs
- Update CHANGELOG.md for all changes

## Submitting Changes

### Pull Request Process
1. Create a feature branch from `main`:
   ```bash
   git checkout -b feature/new-api-implementation
   ```

2. Make your changes following the code guidelines

3. Run the full test suite:
   ```bash
   make ci
   ```

4. Commit your changes with clear, descriptive messages:
   ```bash
   git commit -m "Add NewAPI implementation with comprehensive tests"
   ```

5. Push to your fork and create a pull request

### Pull Request Requirements
- [ ] All tests pass
- [ ] Code follows Go conventions
- [ ] New functionality includes tests
- [ ] Documentation is updated
- [ ] CHANGELOG.md is updated
- [ ] No breaking changes (unless discussed)

### Commit Message Format
Use clear, descriptive commit messages:
- `Add: New feature or API implementation`
- `Fix: Bug fix or correction`
- `Update: Changes to existing functionality`
- `Test: Adding or updating tests`
- `Doc: Documentation updates`

## API Coverage

### Currently Implemented
- ✅ Account API (6 endpoints)
- ✅ Campaign API (5 endpoints)
- ✅ Tool API (6 endpoints)
- ✅ Authentication (OAuth 2.0)

### Pending Implementation
- Ad API
- AdGroup API
- Audience API
- Creative API
- Reporting API
- Business Center API
- And others...

## Swagger Compatibility

This Go SDK maintains compatibility with TikTok's Swagger-based generation approach:

### **Design Philosophy**
- **API Surface Compatibility**: Same endpoints, parameters, and responses as generated SDKs
- **Enhanced Implementation**: Go-idiomatic features built on top of standard API patterns
- **Coexistence Ready**: Can work alongside generated SDKs during migration

### **Swagger Integration**
- `.swagger-codegen-ignore` protects hand-crafted enhancements
- `.swagger-codegen/VERSION` maintains version compatibility
- API models follow OpenAPI specification patterns

### **When Contributing New APIs**
1. **Check Swagger Spec**: Ensure new endpoints match the official OpenAPI specification
2. **Maintain Compatibility**: Keep request/response structures consistent
3. **Add Go Enhancements**: Layer Go-specific features (validation, error handling) on top
4. **Update Ignore File**: Protect custom implementations from generation overwrites

## Getting Help

If you need help or have questions:
1. Check existing issues and documentation
2. Create an issue for bugs or feature requests
3. Follow the existing code patterns and conventions

## Code of Conduct

Please be respectful and professional in all interactions. This project follows the standard open source community guidelines for respectful collaboration.

Thank you for contributing to the TikTok Business API Go SDK!
