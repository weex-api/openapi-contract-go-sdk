package weex

import (
	"fmt"
	"net/http"

	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// APIError represents an error returned by the WEEX Contract API
type APIError struct {
	Code        string               // Error code from API
	Message     string               // Error message from API
	HTTPStatus  int                  // HTTP status code
	RequestTime int64                // Request timestamp from API response
	Category    *types.ErrorCategory // Error category
	Underlying  error                // Underlying error if any
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Underlying != nil {
		return fmt.Sprintf("API error [%s]: %s (HTTP %d) - %v", e.Code, e.Message, e.HTTPStatus, e.Underlying)
	}
	return fmt.Sprintf("API error [%s]: %s (HTTP %d)", e.Code, e.Message, e.HTTPStatus)
}

// IsRetriable returns true if the error is retriable
func (e *APIError) IsRetriable() bool {
	return e.Category != nil && e.Category.Retriable
}

// IsAuthError returns true if the error is an authentication error
func (e *APIError) IsAuthError() bool {
	return e.Category != nil && e.Category.Type == types.ErrTypeAuth
}

// IsRateLimitError returns true if the error is a rate limiting error
func (e *APIError) IsRateLimitError() bool {
	return e.Category != nil && e.Category.Type == types.ErrTypeRateLimit
}

// IsValidationError returns true if the error is a validation error
func (e *APIError) IsValidationError() bool {
	return e.Category != nil && e.Category.Type == types.ErrTypeValidation
}

// IsSystemError returns true if the error is a system error
func (e *APIError) IsSystemError() bool {
	return e.Category != nil && e.Category.Type == types.ErrTypeSystem
}

// NewAPIError creates a new APIError from API response
func NewAPIError(code, message string, httpStatus int, requestTime int64) *APIError {
	return &APIError{
		Code:        code,
		Message:     message,
		HTTPStatus:  httpStatus,
		RequestTime: requestTime,
		Category:    types.GetErrorCategory(code),
	}
}

// WrapError wraps an underlying error with API error information
func WrapError(code, message string, httpStatus int, requestTime int64, underlying error) *APIError {
	err := NewAPIError(code, message, httpStatus, requestTime)
	err.Underlying = underlying
	return err
}

// NetworkError represents a network-related error
type NetworkError struct {
	Operation string // Operation being performed (e.g., "dial", "read", "write")
	URL       string // URL being accessed
	Err       error  // Underlying error
}

// Error implements the error interface
func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error during %s to %s: %v", e.Operation, e.URL, e.Err)
}

// Unwrap returns the underlying error
func (e *NetworkError) Unwrap() error {
	return e.Err
}

// IsRetriable returns true for network errors
func (e *NetworkError) IsRetriable() bool {
	return true
}

// NewNetworkError creates a new NetworkError
func NewNetworkError(operation, url string, err error) *NetworkError {
	return &NetworkError{
		Operation: operation,
		URL:       url,
		Err:       err,
	}
}

// IsRetriableHTTPStatus checks if an HTTP status code indicates a retriable error
func IsRetriableHTTPStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}

// Common error variables
var (
	// ErrInvalidCredentials is returned when API credentials are invalid
	ErrInvalidCredentials = fmt.Errorf("invalid API credentials")

	// ErrMissingCredentials is returned when API credentials are missing
	ErrMissingCredentials = fmt.Errorf("missing API credentials")

	// ErrInvalidConfig is returned when client configuration is invalid
	ErrInvalidConfig = fmt.Errorf("invalid client configuration")

	// ErrContextCanceled is returned when context is canceled
	ErrContextCanceled = fmt.Errorf("context canceled")

	// ErrContextDeadlineExceeded is returned when context deadline is exceeded
	ErrContextDeadlineExceeded = fmt.Errorf("context deadline exceeded")

	// ErrMaxRetriesExceeded is returned when maximum retry attempts are exceeded
	ErrMaxRetriesExceeded = fmt.Errorf("maximum retry attempts exceeded")

	// ErrWebSocketNotConnected is returned when WebSocket is not connected
	ErrWebSocketNotConnected = fmt.Errorf("websocket not connected")

	// ErrWebSocketAlreadyConnected is returned when WebSocket is already connected
	ErrWebSocketAlreadyConnected = fmt.Errorf("websocket already connected")

	// ErrInvalidSubscription is returned when subscription is invalid
	ErrInvalidSubscription = fmt.Errorf("invalid subscription")
)
