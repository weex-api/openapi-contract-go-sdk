// Package rest provides the REST API client for WEEX Contract API
package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/weex/openapi-contract-go-sdk/weex/types"
)

// Logger interface for logging (to avoid importing weex package)
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// Authenticator interface (to avoid importing weex package)
type Authenticator interface {
	GetRESTHeaders(timestamp int64, method, path, body string) map[string]string
}

// Retrier interface (to avoid importing weex package)
type Retrier interface {
	DoWithRetry(ctx context.Context, fn func() error) error
}

// RateLimiter interface (to avoid importing weex package)
type RateLimiter interface {
	WaitForCapacity(ctx context.Context, ipWeight, uidWeight int) error
}

// Client is the REST API client
type Client struct {
	baseURL     string
	locale      string
	auth        Authenticator
	httpClient  *http.Client
	retrier     Retrier
	rateLimiter RateLimiter
	logger      Logger
}

// NewClient creates a new REST API client
func NewClient(baseURL, locale string, httpClient *http.Client, auth Authenticator, retrier Retrier, rateLimiter RateLimiter, logger Logger) *Client {
	return &Client{
		baseURL:     baseURL,
		locale:      locale,
		auth:        auth,
		httpClient:  httpClient,
		retrier:     retrier,
		rateLimiter: rateLimiter,
		logger:      logger,
	}
}

// DoRequest performs an HTTP request with authentication, retry, and rate limiting
func (c *Client) DoRequest(ctx context.Context, method, path string, body interface{}, result interface{}, ipWeight, uidWeight int) error {
	return c.retrier.DoWithRetry(ctx, func() error {
		return c.doRequestOnce(ctx, method, path, body, result, ipWeight, uidWeight)
	})
}

// doRequestOnce performs a single HTTP request attempt
func (c *Client) doRequestOnce(ctx context.Context, method, path string, body interface{}, result interface{}, ipWeight, uidWeight int) error {
	// Wait for rate limit capacity
	if err := c.rateLimiter.WaitForCapacity(ctx, ipWeight, uidWeight); err != nil {
		return fmt.Errorf("rate limit wait failed: %w", err)
	}

	// Prepare request body
	var bodyBytes []byte
	var bodyStr string
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyStr = string(bodyBytes)
	}

	// Build full URL
	url := c.baseURL + types.DefaultAPIPathPrefix + path

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add authentication headers
	timestamp := time.Now().UnixMilli()
	headers := c.auth.GetRESTHeaders(timestamp, method, types.DefaultAPIPathPrefix+path, bodyStr)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Add locale header
	req.Header.Set(types.HeaderLocale, c.locale)

	// Log request
	c.logger.Debug("REST request: %s %s (IP weight: %d, UID weight: %d)", method, path, ipWeight, uidWeight)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Log response
	c.logger.Debug("REST response: %s %s - Status: %d, Body: %s", method, path, resp.StatusCode, string(respBody))

	// Parse response
	return c.parseResponse(resp.StatusCode, respBody, result)
}

// parseResponse parses the API response and handles errors
func (c *Client) parseResponse(statusCode int, body []byte, result interface{}) error {
	// Parse API response wrapper
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for API errors
	if apiResp.Code != "" && apiResp.Code != "0" {
		return fmt.Errorf("API error [%s]: %s (status: %d, time: %d)", apiResp.Code, apiResp.Msg, statusCode, apiResp.RequestTime)
	}

	// Check HTTP status code
	if statusCode >= 400 {
		// If we don't have an error code from API, create a generic error
		if apiResp.Code == "" {
			return fmt.Errorf("HTTP error: %d (time: %d)", statusCode, apiResp.RequestTime)
		}
	}

	// Parse data if result is provided
	if result != nil && len(apiResp.Data) > 0 {
		if err := json.Unmarshal(apiResp.Data, result); err != nil {
			return fmt.Errorf("failed to unmarshal response data: %w", err)
		}
	}

	return nil
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, result interface{}, ipWeight, uidWeight int) error {
	return c.DoRequest(ctx, http.MethodGet, path, nil, result, ipWeight, uidWeight)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}, ipWeight, uidWeight int) error {
	return c.DoRequest(ctx, http.MethodPost, path, body, result, ipWeight, uidWeight)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}, ipWeight, uidWeight int) error {
	return c.DoRequest(ctx, http.MethodPut, path, body, result, ipWeight, uidWeight)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string, body interface{}, result interface{}, ipWeight, uidWeight int) error {
	return c.DoRequest(ctx, http.MethodDelete, path, body, result, ipWeight, uidWeight)
}

// APIResponse represents the standard API response wrapper
type APIResponse struct {
	Code        string          `json:"code"`        // Error code ("0" means success)
	Msg         string          `json:"msg"`         // Error message
	RequestTime int64           `json:"requestTime"` // Request timestamp
	Data        json.RawMessage `json:"data"`        // Response data (can be object or array)
}
