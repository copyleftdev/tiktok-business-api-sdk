# Changelog

All notable changes to the TikTok Business API Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-01

### Added
- Initial release of TikTok Business API Go SDK
- Complete Account API implementation with 6 endpoints
- Complete Campaign API implementation with 5 endpoints  
- Complete Tool API implementation with 6 endpoints
- OAuth 2.0 authentication support
- Built-in rate limiting and retry mechanisms
- Comprehensive error handling with typed errors
- Input validation utilities
- Pagination helper functions
- Context support for all API calls
- Comprehensive test suite with HTTP mocking
- Basic and advanced usage examples
- Complete documentation and API reference

### Features
- **Account Management**: Get advertisers, account info, balance, and fund information
- **Campaign Management**: Full CRUD operations for campaigns
- **Tool API**: Access to languages, currencies, regions, and targeting options
- **Authentication**: OAuth 2.0 flow with token management
- **Error Handling**: Detailed error information with categorization
- **Rate Limiting**: Built-in protection against API rate limits
- **Validation**: Input validation for all API parameters
- **Testing**: Mock HTTP server for testing integrations

### Technical Details
- Go 1.19+ support
- Minimal external dependencies (only golang.org/x/time/rate)
- Type-safe API with comprehensive Go documentation
- Professional project structure following Go conventions
- Makefile for development workflow automation
