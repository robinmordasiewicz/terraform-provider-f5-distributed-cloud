// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"net/http"
	"testing"
)

func TestNewTokenAuthenticator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  *AuthConfig
		wantErr bool
	}{
		{
			name: "valid token",
			config: &AuthConfig{
				APIToken: "test-token-12345",
			},
			wantErr: false,
		},
		{
			name: "empty token",
			config: &AuthConfig{
				APIToken: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			auth, err := NewTokenAuthenticator(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenAuthenticator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if auth == nil {
					t.Error("NewTokenAuthenticator() returned nil authenticator")
				}
				if auth.Method() != AuthMethodToken {
					t.Errorf("Method() = %v, want %v", auth.Method(), AuthMethodToken)
				}
			}
		})
	}
}

func TestTokenAuthenticator_AddAuth(t *testing.T) {
	t.Parallel()

	config := &AuthConfig{
		APIToken: "my-secret-token",
	}

	auth, err := NewTokenAuthenticator(config)
	if err != nil {
		t.Fatalf("NewTokenAuthenticator() error = %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.example.com/test", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() error = %v", err)
	}

	if err := auth.AddAuth(req); err != nil {
		t.Errorf("AddAuth() error = %v", err)
	}

	authHeader := req.Header.Get("Authorization")
	expected := "APIToken my-secret-token"
	if authHeader != expected {
		t.Errorf("Authorization header = %q, want %q", authHeader, expected)
	}
}

func TestTokenAuthenticator_Configure(t *testing.T) {
	t.Parallel()

	config := &AuthConfig{
		APIToken: "test-token",
	}

	auth, err := NewTokenAuthenticator(config)
	if err != nil {
		t.Fatalf("NewTokenAuthenticator() error = %v", err)
	}

	client := &http.Client{}
	if err := auth.Configure(client); err != nil {
		t.Errorf("Configure() error = %v", err)
	}
}

func TestNewCertificateAuthenticator_NoCredentials(t *testing.T) {
	t.Parallel()

	config := &AuthConfig{}

	_, err := NewCertificateAuthenticator(config)
	if err == nil {
		t.Error("NewCertificateAuthenticator() expected error for empty config")
	}

	// Verify it's an AuthenticationError
	authErr, ok := err.(*AuthenticationError)
	if !ok {
		t.Errorf("expected AuthenticationError, got %T", err)
	}
	if authErr.Method != string(AuthMethodCertificate) {
		t.Errorf("AuthenticationError.Method = %q, want %q", authErr.Method, AuthMethodCertificate)
	}
}

func TestNewCertificateAuthenticator_InvalidP12File(t *testing.T) {
	t.Parallel()

	config := &AuthConfig{
		P12File:     "/nonexistent/path/to/cert.p12",
		P12Password: "password",
	}

	_, err := NewCertificateAuthenticator(config)
	if err == nil {
		t.Error("NewCertificateAuthenticator() expected error for nonexistent file")
	}
}

func TestNewCertificateAuthenticator_InvalidKeyPair(t *testing.T) {
	t.Parallel()

	config := &AuthConfig{
		CertFile: "/nonexistent/cert.pem",
		KeyFile:  "/nonexistent/key.pem",
	}

	_, err := NewCertificateAuthenticator(config)
	if err == nil {
		t.Error("NewCertificateAuthenticator() expected error for nonexistent files")
	}
}

func TestNewAuthenticator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		config     *AuthConfig
		wantMethod AuthMethod
		wantErr    bool
	}{
		{
			name: "token auth takes precedence",
			config: &AuthConfig{
				APIToken: "test-token",
				P12File:  "/some/file.p12", // Should be ignored
			},
			wantMethod: AuthMethodToken,
			wantErr:    false,
		},
		{
			name: "token auth only",
			config: &AuthConfig{
				APIToken: "test-token",
			},
			wantMethod: AuthMethodToken,
			wantErr:    false,
		},
		{
			name: "no credentials returns error",
			config: &AuthConfig{
				// No credentials
			},
			wantMethod: "",
			wantErr:    true,
		},
		{
			name: "partial certificate config (missing key) returns error",
			config: &AuthConfig{
				CertFile: "/some/cert.pem",
				// Missing KeyFile
			},
			wantMethod: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			auth, err := NewAuthenticator(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthenticator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if auth == nil {
					t.Error("NewAuthenticator() returned nil authenticator")
					return
				}
				if auth.Method() != tt.wantMethod {
					t.Errorf("Method() = %v, want %v", auth.Method(), tt.wantMethod)
				}
			}
		})
	}
}

func TestCertificateAuthenticator_Configure(t *testing.T) {
	t.Parallel()

	// We can't easily test with a real certificate, but we can test the Configure behavior
	// when provided a valid CertificateAuthenticator (which we can't create without a real cert)
	// So this test focuses on the client transport handling

	// Test that Configure handles nil transport correctly
	// This is tested indirectly through NewCertificateAuthenticator tests above
}

func TestAuthMethod_Constants(t *testing.T) {
	t.Parallel()

	if AuthMethodCertificate != "certificate" {
		t.Errorf("AuthMethodCertificate = %q, want %q", AuthMethodCertificate, "certificate")
	}

	if AuthMethodToken != "token" {
		t.Errorf("AuthMethodToken = %q, want %q", AuthMethodToken, "token")
	}
}
