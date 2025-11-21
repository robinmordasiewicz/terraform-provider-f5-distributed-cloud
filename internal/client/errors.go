// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"errors"
	"fmt"
	"net/http"
)

// Common error types for the F5 XC API client.
var (
	ErrNotFound           = errors.New("resource not found")
	ErrUnauthorized       = errors.New("unauthorized: invalid or missing credentials")
	ErrForbidden          = errors.New("forbidden: insufficient permissions")
	ErrBadRequest         = errors.New("bad request: invalid request parameters")
	ErrConflict           = errors.New("conflict: resource already exists or version mismatch")
	ErrInternalServer     = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrRateLimited        = errors.New("rate limited: too many requests")
	ErrTimeout            = errors.New("request timeout")
)

// APIError represents an error returned by the F5 XC API.
type APIError struct {
	StatusCode int
	Message    string
	RequestID  string
	Err        error
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("API error (status %d, request_id: %s): %s", e.StatusCode, e.RequestID, e.Message)
	}
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

// Unwrap returns the underlying error.
func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new APIError from an HTTP response.
func NewAPIError(statusCode int, message string, requestID string) *APIError {
	var baseErr error
	switch statusCode {
	case http.StatusNotFound:
		baseErr = ErrNotFound
	case http.StatusUnauthorized:
		baseErr = ErrUnauthorized
	case http.StatusForbidden:
		baseErr = ErrForbidden
	case http.StatusBadRequest:
		baseErr = ErrBadRequest
	case http.StatusConflict:
		baseErr = ErrConflict
	case http.StatusInternalServerError:
		baseErr = ErrInternalServer
	case http.StatusServiceUnavailable:
		baseErr = ErrServiceUnavailable
	case http.StatusTooManyRequests:
		baseErr = ErrRateLimited
	case http.StatusGatewayTimeout, http.StatusRequestTimeout:
		baseErr = ErrTimeout
	}

	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		RequestID:  requestID,
		Err:        baseErr,
	}
}

// IsNotFound returns true if the error is a not found error.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsUnauthorized returns true if the error is an unauthorized error.
func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

// IsRateLimited returns true if the error is a rate limited error.
func IsRateLimited(err error) bool {
	return errors.Is(err, ErrRateLimited)
}

// IsRetryable returns true if the error is retryable.
func IsRetryable(err error) bool {
	return errors.Is(err, ErrInternalServer) ||
		errors.Is(err, ErrServiceUnavailable) ||
		errors.Is(err, ErrRateLimited) ||
		errors.Is(err, ErrTimeout)
}

// AuthenticationError represents an authentication-specific error.
type AuthenticationError struct {
	Method  string // "certificate" or "token"
	Message string
	Err     error
}

// Error implements the error interface.
func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("authentication error (%s): %s", e.Method, e.Message)
}

// Unwrap returns the underlying error.
func (e *AuthenticationError) Unwrap() error {
	return e.Err
}

// NewAuthenticationError creates a new AuthenticationError.
func NewAuthenticationError(method string, message string, err error) *AuthenticationError {
	return &AuthenticationError{
		Method:  method,
		Message: message,
		Err:     err,
	}
}
