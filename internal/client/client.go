// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// DefaultRetryMax is the default maximum number of retries.
const DefaultRetryMax = 3

// DefaultRetryWaitMin is the default minimum wait time between retries.
const DefaultRetryWaitMin = 1 * time.Second

// DefaultRetryWaitMax is the default maximum wait time between retries.
const DefaultRetryWaitMax = 30 * time.Second

// DefaultTimeout is the default HTTP client timeout.
const DefaultTimeout = 60 * time.Second

// ClientConfig holds the configuration for the F5 XC API client.
type ClientConfig struct {
	// APIURL is the base URL for the F5 XC API.
	APIURL string

	// Auth contains authentication configuration.
	Auth *AuthConfig

	// RetryMax is the maximum number of retries.
	RetryMax int

	// RetryWaitMin is the minimum wait time between retries.
	RetryWaitMin time.Duration

	// RetryWaitMax is the maximum wait time between retries.
	RetryWaitMax time.Duration

	// Timeout is the HTTP client timeout.
	Timeout time.Duration

	// UserAgent is the User-Agent header value.
	UserAgent string
}

// Client is the F5 XC API client.
type Client struct {
	baseURL       *url.URL
	httpClient    *retryablehttp.Client
	authenticator Authenticator
	userAgent     string
}

// NewClient creates a new F5 XC API client.
func NewClient(config *ClientConfig) (*Client, error) {
	if config.APIURL == "" {
		return nil, fmt.Errorf("api_url is required")
	}

	baseURL, err := url.Parse(config.APIURL)
	if err != nil {
		return nil, fmt.Errorf("invalid api_url: %w", err)
	}

	// Create authenticator
	authenticator, err := NewAuthenticator(config.Auth)
	if err != nil {
		return nil, err
	}

	// Set defaults
	retryMax := config.RetryMax
	if retryMax == 0 {
		retryMax = DefaultRetryMax
	}

	retryWaitMin := config.RetryWaitMin
	if retryWaitMin == 0 {
		retryWaitMin = DefaultRetryWaitMin
	}

	retryWaitMax := config.RetryWaitMax
	if retryWaitMax == 0 {
		retryWaitMax = DefaultRetryWaitMax
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	userAgent := config.UserAgent
	if userAgent == "" {
		userAgent = "terraform-provider-f5-distributed-cloud"
	}

	// Create retryable HTTP client
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = retryMax
	retryClient.RetryWaitMin = retryWaitMin
	retryClient.RetryWaitMax = retryWaitMax
	retryClient.HTTPClient.Timeout = timeout
	retryClient.CheckRetry = customRetryPolicy
	retryClient.Logger = nil // Disable default logging

	// Configure authentication on the underlying HTTP client
	if err := authenticator.Configure(retryClient.HTTPClient); err != nil {
		return nil, fmt.Errorf("failed to configure authentication: %w", err)
	}

	return &Client{
		baseURL:       baseURL,
		httpClient:    retryClient,
		authenticator: authenticator,
		userAgent:     userAgent,
	}, nil
}

// customRetryPolicy implements a custom retry policy.
func customRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// Do not retry on context cancellation
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// Retry on connection errors
	if err != nil {
		return true, nil
	}

	// Retry on specific status codes
	switch resp.StatusCode {
	case http.StatusTooManyRequests,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusBadGateway:
		return true, nil
	}

	return false, nil
}

// Do performs an HTTP request.
func (c *Client) Do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// Build URL by joining base URL with path
	// Note: url.ResolveReference doesn't work correctly when base URL has a path component
	// so we manually join the paths
	reqURL := c.baseURL.String() + path

	// Prepare body
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create request
	req, err := retryablehttp.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	// Add authentication
	if err := c.authenticator.AddAuth(req.Request); err != nil {
		return fmt.Errorf("failed to add authentication: %w", err)
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		requestID := resp.Header.Get("X-Request-Id")
		message := string(respBody)

		// Try to parse error response
		var errorResp struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}
		if json.Unmarshal(respBody, &errorResp) == nil {
			if errorResp.Message != "" {
				message = errorResp.Message
			} else if errorResp.Error != "" {
				message = errorResp.Error
			}
		}

		return NewAPIError(resp.StatusCode, message, requestID)
	}

	// Parse response
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	return c.Do(ctx, http.MethodGet, path, nil, result)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodPost, path, body, result)
}

// Put performs a PUT request.
func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, http.MethodPut, path, body, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.Do(ctx, http.MethodDelete, path, nil, nil)
}
