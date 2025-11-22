// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewAPIError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		statusCode     int
		message        string
		requestID      string
		wantBaseErr    error
		wantErrContain string
	}{
		{
			name:           "not found error",
			statusCode:     http.StatusNotFound,
			message:        "resource not found",
			requestID:      "req-123",
			wantBaseErr:    ErrNotFound,
			wantErrContain: "status 404",
		},
		{
			name:           "unauthorized error",
			statusCode:     http.StatusUnauthorized,
			message:        "invalid token",
			requestID:      "",
			wantBaseErr:    ErrUnauthorized,
			wantErrContain: "status 401",
		},
		{
			name:           "forbidden error",
			statusCode:     http.StatusForbidden,
			message:        "access denied",
			requestID:      "req-456",
			wantBaseErr:    ErrForbidden,
			wantErrContain: "status 403",
		},
		{
			name:           "bad request error",
			statusCode:     http.StatusBadRequest,
			message:        "invalid parameters",
			requestID:      "",
			wantBaseErr:    ErrBadRequest,
			wantErrContain: "status 400",
		},
		{
			name:           "conflict error",
			statusCode:     http.StatusConflict,
			message:        "resource exists",
			requestID:      "",
			wantBaseErr:    ErrConflict,
			wantErrContain: "status 409",
		},
		{
			name:           "internal server error",
			statusCode:     http.StatusInternalServerError,
			message:        "server error",
			requestID:      "",
			wantBaseErr:    ErrInternalServer,
			wantErrContain: "status 500",
		},
		{
			name:           "service unavailable error",
			statusCode:     http.StatusServiceUnavailable,
			message:        "service down",
			requestID:      "",
			wantBaseErr:    ErrServiceUnavailable,
			wantErrContain: "status 503",
		},
		{
			name:           "rate limited error",
			statusCode:     http.StatusTooManyRequests,
			message:        "too many requests",
			requestID:      "",
			wantBaseErr:    ErrRateLimited,
			wantErrContain: "status 429",
		},
		{
			name:           "gateway timeout error",
			statusCode:     http.StatusGatewayTimeout,
			message:        "gateway timeout",
			requestID:      "",
			wantBaseErr:    ErrTimeout,
			wantErrContain: "status 504",
		},
		{
			name:           "request timeout error",
			statusCode:     http.StatusRequestTimeout,
			message:        "request timeout",
			requestID:      "",
			wantBaseErr:    ErrTimeout,
			wantErrContain: "status 408",
		},
		{
			name:           "unknown status code",
			statusCode:     418, // I'm a teapot
			message:        "teapot",
			requestID:      "",
			wantBaseErr:    nil,
			wantErrContain: "status 418",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			apiErr := NewAPIError(tt.statusCode, tt.message, tt.requestID)

			if apiErr.StatusCode != tt.statusCode {
				t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, tt.statusCode)
			}

			if apiErr.Message != tt.message {
				t.Errorf("Message = %q, want %q", apiErr.Message, tt.message)
			}

			if apiErr.RequestID != tt.requestID {
				t.Errorf("RequestID = %q, want %q", apiErr.RequestID, tt.requestID)
			}

			// Test Unwrap returns correct base error
			if tt.wantBaseErr != nil && !errors.Is(apiErr, tt.wantBaseErr) {
				t.Errorf("Unwrap() error = %v, want %v", apiErr.Unwrap(), tt.wantBaseErr)
			}

			// Test Error() string contains expected content
			errStr := apiErr.Error()
			if tt.wantErrContain != "" && !contains(errStr, tt.wantErrContain) {
				t.Errorf("Error() = %q, want to contain %q", errStr, tt.wantErrContain)
			}
		})
	}
}

func TestAPIError_WithRequestID(t *testing.T) {
	t.Parallel()

	apiErr := NewAPIError(http.StatusNotFound, "not found", "req-abc-123")
	errStr := apiErr.Error()

	if !contains(errStr, "req-abc-123") {
		t.Errorf("Error() = %q, want to contain request_id", errStr)
	}
}

func TestIsNotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "API error with not found",
			err:  NewAPIError(http.StatusNotFound, "not found", ""),
			want: true,
		},
		{
			name: "direct ErrNotFound",
			err:  ErrNotFound,
			want: true,
		},
		{
			name: "API error with unauthorized",
			err:  NewAPIError(http.StatusUnauthorized, "unauthorized", ""),
			want: false,
		},
		{
			name: "generic error",
			err:  errors.New("some error"),
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsNotFound(tt.err); got != tt.want {
				t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnauthorized(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "API error with unauthorized",
			err:  NewAPIError(http.StatusUnauthorized, "unauthorized", ""),
			want: true,
		},
		{
			name: "direct ErrUnauthorized",
			err:  ErrUnauthorized,
			want: true,
		},
		{
			name: "API error with not found",
			err:  NewAPIError(http.StatusNotFound, "not found", ""),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsUnauthorized(tt.err); got != tt.want {
				t.Errorf("IsUnauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRateLimited(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "API error with rate limited",
			err:  NewAPIError(http.StatusTooManyRequests, "rate limited", ""),
			want: true,
		},
		{
			name: "direct ErrRateLimited",
			err:  ErrRateLimited,
			want: true,
		},
		{
			name: "API error with not found",
			err:  NewAPIError(http.StatusNotFound, "not found", ""),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsRateLimited(tt.err); got != tt.want {
				t.Errorf("IsRateLimited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRetryable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "internal server error is retryable",
			err:  NewAPIError(http.StatusInternalServerError, "server error", ""),
			want: true,
		},
		{
			name: "service unavailable is retryable",
			err:  NewAPIError(http.StatusServiceUnavailable, "unavailable", ""),
			want: true,
		},
		{
			name: "rate limited is retryable",
			err:  NewAPIError(http.StatusTooManyRequests, "rate limited", ""),
			want: true,
		},
		{
			name: "timeout is retryable",
			err:  NewAPIError(http.StatusGatewayTimeout, "timeout", ""),
			want: true,
		},
		{
			name: "not found is not retryable",
			err:  NewAPIError(http.StatusNotFound, "not found", ""),
			want: false,
		},
		{
			name: "bad request is not retryable",
			err:  NewAPIError(http.StatusBadRequest, "bad request", ""),
			want: false,
		},
		{
			name: "unauthorized is not retryable",
			err:  NewAPIError(http.StatusUnauthorized, "unauthorized", ""),
			want: false,
		},
		{
			name: "direct ErrInternalServer",
			err:  ErrInternalServer,
			want: true,
		},
		{
			name: "direct ErrServiceUnavailable",
			err:  ErrServiceUnavailable,
			want: true,
		},
		{
			name: "direct ErrRateLimited",
			err:  ErrRateLimited,
			want: true,
		},
		{
			name: "direct ErrTimeout",
			err:  ErrTimeout,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsRetryable(tt.err); got != tt.want {
				t.Errorf("IsRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthenticationError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		method     string
		message    string
		wrappedErr error
	}{
		{
			name:       "certificate auth error",
			method:     "certificate",
			message:    "failed to load certificate",
			wrappedErr: errors.New("file not found"),
		},
		{
			name:       "token auth error",
			method:     "token",
			message:    "invalid token format",
			wrappedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authErr := NewAuthenticationError(tt.method, tt.message, tt.wrappedErr)

			if authErr.Method != tt.method {
				t.Errorf("Method = %q, want %q", authErr.Method, tt.method)
			}

			if authErr.Message != tt.message {
				t.Errorf("Message = %q, want %q", authErr.Message, tt.message)
			}

			// Test Error() contains method and message
			errStr := authErr.Error()
			if !contains(errStr, tt.method) {
				t.Errorf("Error() = %q, want to contain method %q", errStr, tt.method)
			}
			if !contains(errStr, tt.message) {
				t.Errorf("Error() = %q, want to contain message %q", errStr, tt.message)
			}

			// Test Unwrap
			if tt.wrappedErr != nil && authErr.Unwrap() != tt.wrappedErr {
				t.Errorf("Unwrap() = %v, want %v", authErr.Unwrap(), tt.wrappedErr)
			}
		})
	}
}

// contains checks if s contains substr
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}
