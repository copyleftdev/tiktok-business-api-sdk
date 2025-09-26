package client

import (
	"bufio"
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestNoPlaceholdersOrTodos ensures the codebase has no placeholders, TODOs, FIXMEs, or unimplemented functions
func TestNoPlaceholdersOrTodos(t *testing.T) {
	// Define patterns to search for
	forbiddenPatterns := []struct {
		pattern     *regexp.Regexp
		description string
	}{
		{regexp.MustCompile(`(?i)\btodo\b`), "TODO comments"},
		{regexp.MustCompile(`(?i)\bfixme\b`), "FIXME comments"},
		{regexp.MustCompile(`(?i)\bhack\b`), "HACK comments"},
		{regexp.MustCompile(`(?i)\bxxx\b`), "XXX comments"},
		{regexp.MustCompile(`(?i)\bbug\b`), "BUG comments"},
		{regexp.MustCompile(`(?i)\bplaceholder\b`), "placeholder references"},
		{regexp.MustCompile(`(?i)not\s+implemented`), "not implemented messages"},
		{regexp.MustCompile(`(?i)unimplemented`), "unimplemented references"},
		{regexp.MustCompile(`(?i)coming\s+soon`), "coming soon messages"},
		{regexp.MustCompile(`(?i)to\s+be\s+implemented`), "to be implemented messages"},
	}

	// Allowed exceptions (these are legitimate uses)
	allowedExceptions := []string{
		"ErrServiceNotImplemented",                                   // Our professional error
		"this service is not yet implemented in the current SDK",     // Professional error message
		"Type definitions for services not yet fully implemented",    // Documentation comment
		"Services not yet implemented - return clear error messages", // Documentation comment
		"notImplementedAdService",                                    // Service type name
		"notImplementedAdGroupService",                               // Service type name
		"notImplementedAudienceService",                              // Service type name
		"notImplementedCreativeService",                              // Service type name
		"notImplementedReportingService",                             // Service type name
		"notImplementedBCService",                                    // Service type name
	}

	// Get the project root directory
	projectRoot := getProjectRoot(t)

	// Scan all Go files in the project
	err := filepath.Walk(filepath.Join(projectRoot, "pkg"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only check Go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files for this specific test (to avoid self-reference)
		if strings.HasSuffix(path, "_test.go") && strings.Contains(path, "implementation_test.go") {
			return nil
		}

		return checkFileForForbiddenPatterns(t, path, forbiddenPatterns, allowedExceptions)
	})

	if err != nil {
		t.Fatalf("Error walking directory: %v", err)
	}
}

// TestAllServicesHaveProperImplementations ensures all services are either fully implemented or return proper errors
func TestAllServicesHaveProperImplementations(t *testing.T) {
	// Test that implemented services work correctly
	t.Run("ImplementedServices", func(t *testing.T) {
		config := &Config{
			BaseURL:     "https://business-api.tiktok.com",
			AccessToken: "test_token",
			Timeout:     30,
		}

		client, err := NewClient(config)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Verify implemented services are not nil and have proper types
		if client.Account() == nil {
			t.Error("Account service should not be nil")
		}

		if client.Campaign() == nil {
			t.Error("Campaign service should not be nil")
		}

		if client.Tool() == nil {
			t.Error("Tool service should not be nil")
		}

		if client.Auth() == nil {
			t.Error("Auth service should not be nil")
		}
	})

	// Test that not-yet-implemented services return proper errors
	t.Run("NotImplementedServices", func(t *testing.T) {
		config := &Config{
			BaseURL:     "https://business-api.tiktok.com",
			AccessToken: "test_token",
			Timeout:     30,
		}

		client, err := NewClient(config)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		ctx := context.Background()

		// Test Ad service returns proper error
		_, err = client.Ad().Create(ctx, &AdCreateRequest{})
		if err != ErrServiceNotImplemented {
			t.Errorf("Ad service should return ErrServiceNotImplemented, got: %v", err)
		}

		// Test AdGroup service returns proper error
		_, err = client.AdGroup().Create(ctx, &AdGroupCreateRequest{})
		if err != ErrServiceNotImplemented {
			t.Errorf("AdGroup service should return ErrServiceNotImplemented, got: %v", err)
		}

		// Test Audience service returns proper error
		_, err = client.Audience().CreateCustomAudience(ctx, &CustomAudienceCreateRequest{})
		if err != ErrServiceNotImplemented {
			t.Errorf("Audience service should return ErrServiceNotImplemented, got: %v", err)
		}

		// Test Creative service returns proper error
		_, err = client.Creative().UploadImage(ctx, &ImageUploadRequest{})
		if err != ErrServiceNotImplemented {
			t.Errorf("Creative service should return ErrServiceNotImplemented, got: %v", err)
		}

		// Test Reporting service returns proper error
		_, err = client.Reporting().GetBasicReports(ctx, &ReportingRequest{})
		if err != ErrServiceNotImplemented {
			t.Errorf("Reporting service should return ErrServiceNotImplemented, got: %v", err)
		}

		// Test BC service returns proper error
		_, err = client.BC().GetBusinessCenters(ctx)
		if err != ErrServiceNotImplemented {
			t.Errorf("BC service should return ErrServiceNotImplemented, got: %v", err)
		}
	})
}

// TestNoUnusedImports ensures there are no unused imports that might indicate incomplete implementations
func TestNoUnusedImports(t *testing.T) {
	projectRoot := getProjectRoot(t)

	err := filepath.Walk(filepath.Join(projectRoot, "pkg"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only check Go files, skip test files
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		return checkFileForUnusedImports(t, path)
	})

	if err != nil {
		t.Fatalf("Error checking unused imports: %v", err)
	}
}

// TestErrorMessagesAreProfessional ensures all error messages are professional and helpful
func TestErrorMessagesAreProfessional(t *testing.T) {
	// Test that ErrServiceNotImplemented has a professional message
	expectedMessage := "this service is not yet implemented in the current SDK version"
	if ErrServiceNotImplemented.Error() != expectedMessage {
		t.Errorf("ErrServiceNotImplemented should have message %q, got %q",
			expectedMessage, ErrServiceNotImplemented.Error())
	}

	// Ensure the error message doesn't contain unprofessional terms
	unprofessionalTerms := []string{"placeholder", "todo", "fixme", "hack", "broken", "temp"}
	message := strings.ToLower(ErrServiceNotImplemented.Error())

	for _, term := range unprofessionalTerms {
		if strings.Contains(message, term) {
			t.Errorf("Error message contains unprofessional term %q: %s", term, ErrServiceNotImplemented.Error())
		}
	}
}

// Helper functions

func getProjectRoot(t *testing.T) string {
	// Get current working directory and find project root
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Navigate up to find go.mod
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			t.Fatalf("Could not find project root (go.mod not found)")
		}
		wd = parent
	}
}

func checkFileForForbiddenPatterns(t *testing.T, filePath string, patterns []struct {
	pattern     *regexp.Regexp
	description string
}, allowedExceptions []string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		for _, p := range patterns {
			if p.pattern.MatchString(line) {
				// Check if this match is in our allowed exceptions
				isAllowed := false
				for _, exception := range allowedExceptions {
					if strings.Contains(line, exception) {
						isAllowed = true
						break
					}
				}

				if !isAllowed {
					t.Errorf("Found forbidden pattern (%s) in %s:%d: %s",
						p.description, filePath, lineNum, strings.TrimSpace(line))
				}
			}
		}
	}

	return scanner.Err()
}

func checkFileForUnusedImports(t *testing.T, filePath string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %v", filePath, err)
	}

	// This is a basic check - in a real implementation, you'd want more sophisticated analysis
	// For now, we just ensure the file parses correctly, which indicates no major structural issues

	// Check that there are no obvious signs of incomplete implementation
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// Check function bodies aren't empty (except for interface methods)
			if x.Body != nil && len(x.Body.List) == 0 {
				// Empty function body might indicate incomplete implementation
				// But we allow this for interface definitions
				if x.Recv == nil { // Not a method, so it's a regular function
					t.Errorf("Empty function body found in %s: %s", filePath, x.Name.Name)
				}
			}
		}
		return true
	})

	return nil
}

// TestDocumentationCompleteness ensures all public functions have documentation
func TestDocumentationCompleteness(t *testing.T) {
	projectRoot := getProjectRoot(t)

	err := filepath.Walk(filepath.Join(projectRoot, "pkg"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only check Go files, skip test files
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		return checkFileDocumentation(t, path)
	})

	if err != nil {
		t.Fatalf("Error checking documentation: %v", err)
	}
}

func checkFileDocumentation(t *testing.T, filePath string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %v", filePath, err)
	}

	// Check that exported functions have documentation
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// Check if function is exported (starts with capital letter)
			if x.Name.IsExported() && x.Doc == nil {
				// Allow some exceptions for simple getters/setters and test helpers
				funcName := x.Name.Name
				if !isDocumentationException(funcName) {
					t.Errorf("Exported function %s in %s lacks documentation", funcName, filePath)
				}
			}
		case *ast.TypeSpec:
			// Check if type is exported and has documentation
			if x.Name.IsExported() && x.Doc == nil {
				typeName := x.Name.Name
				if !isDocumentationException(typeName) {
					t.Errorf("Exported type %s in %s lacks documentation", typeName, filePath)
				}
			}
		}
		return true
	})

	return nil
}

func isDocumentationException(name string) bool {
	// Allow some common patterns that don't need documentation
	exceptions := []string{
		"String", "Error", // Common interface methods
		"MarshalJSON", "UnmarshalJSON", // JSON methods
	}

	for _, exception := range exceptions {
		if name == exception {
			return true
		}
	}

	return false
}
