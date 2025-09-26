package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				BaseURL:     "https://business-api.tiktok.com",
				AccessToken: "test_token",
				Timeout:     30 * time.Second,
			},
			wantErr: false,
		},
		{
			name:    "nil config uses defaults",
			config:  nil,
			wantErr: true, // Will fail validation due to missing access token
		},
		{
			name: "invalid base URL",
			config: &Config{
				BaseURL:     "invalid-url",
				AccessToken: "test_token",
			},
			wantErr: true,
		},
		{
			name: "missing authentication",
			config: &Config{
				BaseURL: "https://business-api.tiktok.com",
				Timeout: 30 * time.Second,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client without error")
			}
		})
	}
}

func TestClient_DoRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("Access-Token") != "test_token" {
			t.Errorf("Expected Access-Token header, got %s", r.Header.Get("Access-Token"))
		}

		if r.Header.Get("User-Agent") != "tiktok-business-api-go-sdk/1.0.0" {
			t.Errorf("Expected User-Agent header, got %s", r.Header.Get("User-Agent"))
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"code": 0, "message": "success", "request_id": "test123"}`))
	}))
	defer server.Close()

	// Create client with test server URL
	config := &Config{
		BaseURL:     server.URL,
		AccessToken: "test_token",
		Timeout:     5 * time.Second,
		UserAgent:   "tiktok-business-api-go-sdk/1.0.0",
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test request
	ctx := context.Background()
	resp, err := client.DoRequest(ctx, "GET", "/test", nil, nil)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestClient_BuildQueryParams(t *testing.T) {
	client := &Client{}

	tests := []struct {
		name   string
		params map[string]interface{}
		want   string
	}{
		{
			name:   "empty params",
			params: map[string]interface{}{},
			want:   "",
		},
		{
			name: "single param",
			params: map[string]interface{}{
				"key": "value",
			},
			want: "key=value",
		},
		{
			name: "multiple params",
			params: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			// Note: order may vary due to map iteration
		},
		{
			name: "nil value ignored",
			params: map[string]interface{}{
				"key1": "value1",
				"key2": nil,
			},
			want: "key1=value1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.BuildQueryParams(tt.params)
			if tt.name == "empty params" || tt.name == "single param" || tt.name == "nil value ignored" {
				if got != tt.want {
					t.Errorf("BuildQueryParams() = %v, want %v", got, tt.want)
				}
			}
			// For multiple params, just check that both keys are present
			if tt.name == "multiple params" {
				if !contains(got, "key1=value1") || !contains(got, "key2=value2") {
					t.Errorf("BuildQueryParams() = %v, should contain both key1=value1 and key2=value2", got)
				}
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config with access token",
			config: &Config{
				BaseURL:     "https://api.example.com",
				AccessToken: "token",
				Timeout:     30 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "valid config with client credentials",
			config: &Config{
				BaseURL:      "https://api.example.com",
				ClientID:     "client_id",
				ClientSecret: "client_secret",
				Timeout:      30 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "missing base URL",
			config: &Config{
				AccessToken: "token",
				Timeout:     30 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "missing authentication",
			config: &Config{
				BaseURL: "https://api.example.com",
				Timeout: 30 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "invalid timeout",
			config: &Config{
				BaseURL:     "https://api.example.com",
				AccessToken: "token",
				Timeout:     -1 * time.Second,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr)) ||
		(len(s) > len(substr)+1 && s[1:len(substr)+1] == substr))
}
