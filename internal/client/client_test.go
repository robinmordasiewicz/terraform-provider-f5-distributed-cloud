// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  *ClientConfig
		wantErr bool
	}{
		{
			name: "valid config with token",
			config: &ClientConfig{
				APIURL: "https://api.example.com",
				Auth: &AuthConfig{
					APIToken: "test-token",
				},
			},
			wantErr: false,
		},
		{
			name: "missing API URL",
			config: &ClientConfig{
				APIURL: "",
				Auth: &AuthConfig{
					APIToken: "test-token",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid API URL",
			config: &ClientConfig{
				APIURL: "://invalid-url",
				Auth: &AuthConfig{
					APIToken: "test-token",
				},
			},
			wantErr: true,
		},
		{
			name: "missing auth config",
			config: &ClientConfig{
				APIURL: "https://api.example.com",
				Auth:   &AuthConfig{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client")
			}
		})
	}
}

func TestNewClient_Defaults(t *testing.T) {
	t.Parallel()

	config := &ClientConfig{
		APIURL: "https://api.example.com",
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	// Verify defaults are applied
	if client.userAgent != "terraform-provider-f5-distributed-cloud" {
		t.Errorf("default userAgent = %q, want %q", client.userAgent, "terraform-provider-f5-distributed-cloud")
	}
}

func TestNewClient_CustomSettings(t *testing.T) {
	t.Parallel()

	config := &ClientConfig{
		APIURL: "https://api.example.com",
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
		RetryMax:     5,
		RetryWaitMin: 2 * time.Second,
		RetryWaitMax: 60 * time.Second,
		Timeout:      120 * time.Second,
		UserAgent:    "custom-agent/1.0",
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client.userAgent != "custom-agent/1.0" {
		t.Errorf("userAgent = %q, want %q", client.userAgent, "custom-agent/1.0")
	}
}

func TestClient_Do_Success(t *testing.T) {
	t.Parallel()

	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q, want %q", r.Header.Get("Content-Type"), "application/json")
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Accept = %q, want %q", r.Header.Get("Accept"), "application/json")
		}
		if r.Header.Get("Authorization") != "APIToken test-token" {
			t.Errorf("Authorization = %q, want %q", r.Header.Get("Authorization"), "APIToken test-token")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var result map[string]string
	err = client.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}

	if result["status"] != "ok" {
		t.Errorf("result = %v, want status=ok", result)
	}
}

func TestClient_Do_NotFound(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "req-12345")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "resource not found"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var result map[string]string
	err = client.Get(context.Background(), "/notfound", &result)
	if err == nil {
		t.Error("Get() expected error for 404")
	}

	if !IsNotFound(err) {
		t.Errorf("IsNotFound(err) = false, want true")
	}
}

func TestClient_Do_Unauthorized(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid token"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "invalid-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Error("Get() expected error for 401")
	}

	if !IsUnauthorized(err) {
		t.Errorf("IsUnauthorized(err) = false, want true")
	}
}

func TestClient_Do_RateLimited(t *testing.T) {
	t.Parallel()

	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{"message": "rate limited"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
		RetryMax:     2,
		RetryWaitMin: 10 * time.Millisecond,
		RetryWaitMax: 20 * time.Millisecond,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Error("Get() expected error for rate limited")
	}

	// Should have retried due to rate limiting (initial + RetryMax retries)
	// RetryMax=2 means 3 total attempts (1 initial + 2 retries)
	if callCount < 2 {
		t.Errorf("callCount = %d, want >= 2 (retries)", callCount)
	}

	// Error should contain retry information
	errStr := err.Error()
	if !findSubstring(errStr, "giving up after") && !findSubstring(errStr, "rate") {
		t.Logf("Got error: %v", err)
	}
}

func TestClient_Post(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method = %q, want %q", r.Method, http.MethodPost)
		}

		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if body["name"] != "test" {
			t.Errorf("body.name = %q, want %q", body["name"], "test")
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": "123"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var result map[string]string
	err = client.Post(context.Background(), "/resources", map[string]string{"name": "test"}, &result)
	if err != nil {
		t.Errorf("Post() error = %v", err)
	}

	if result["id"] != "123" {
		t.Errorf("result.id = %q, want %q", result["id"], "123")
	}
}

func TestClient_Put(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method = %q, want %q", r.Method, http.MethodPut)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var result map[string]string
	err = client.Put(context.Background(), "/resources/123", map[string]string{"name": "updated"}, &result)
	if err != nil {
		t.Errorf("Put() error = %v", err)
	}
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method = %q, want %q", r.Method, http.MethodDelete)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = client.Delete(context.Background(), "/resources/123")
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}

func TestClient_ContextCancellation(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(1 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := &ClientConfig{
		APIURL: server.URL,
		Auth: &AuthConfig{
			APIToken: "test-token",
		},
		Timeout: 5 * time.Second,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err = client.Get(ctx, "/slow", nil)
	if err == nil {
		t.Error("Get() expected error for cancelled context")
	}
}

func TestDefaultConstants(t *testing.T) {
	t.Parallel()

	if DefaultRetryMax != 3 {
		t.Errorf("DefaultRetryMax = %d, want %d", DefaultRetryMax, 3)
	}

	if DefaultRetryWaitMin != 1*time.Second {
		t.Errorf("DefaultRetryWaitMin = %v, want %v", DefaultRetryWaitMin, 1*time.Second)
	}

	if DefaultRetryWaitMax != 30*time.Second {
		t.Errorf("DefaultRetryWaitMax = %v, want %v", DefaultRetryWaitMax, 30*time.Second)
	}

	if DefaultTimeout != 60*time.Second {
		t.Errorf("DefaultTimeout = %v, want %v", DefaultTimeout, 60*time.Second)
	}
}

// findSubstring checks if s contains substr
func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
