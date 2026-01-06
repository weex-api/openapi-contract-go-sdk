// Package weex provides the main client for the WEEX Contract API.
package weex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// Authenticator handles API authentication and signature generation
type Authenticator struct {
	apiKey     string
	secretKey  string
	passphrase string
}

// NewAuthenticator creates a new Authenticator instance
func NewAuthenticator(apiKey, secretKey, passphrase string) *Authenticator {
	return &Authenticator{
		apiKey:     apiKey,
		secretKey:  secretKey,
		passphrase: passphrase,
	}
}

// SignRequest generates the HMAC SHA256 signature for a REST API request
//
// The signature algorithm is:
//
//	Message = timestamp + method + requestPath + body
//	Signature = Base64(HMAC-SHA256(secretKey, Message))
//
// Parameters:
//   - timestamp: Unix timestamp in milliseconds
//   - method: HTTP method (GET, POST, PUT, DELETE)
//   - path: Request path (e.g., "/capi/v2/market/contracts")
//   - body: Request body as string (empty string for GET requests)
//
// Returns the base64-encoded signature string
func (a *Authenticator) SignRequest(timestamp int64, method, path, body string) string {
	message := fmt.Sprintf("%d%s%s%s", timestamp, method, path, body)
	return a.sign(message)
}

// SignWebSocket generates the HMAC SHA256 signature for WebSocket authentication
//
// The signature algorithm is:
//
//	Message = timestamp + method + requestPath + body
//	Signature = Base64(HMAC-SHA256(secretKey, Message))
//
// Parameters:
//   - timestamp: Unix timestamp in seconds (not milliseconds)
//   - method: HTTP method (typically "GET" for WebSocket)
//   - path: WebSocket path (e.g., "/users/self/verify")
//   - body: Request body as string (empty string for auth)
//
// Returns the base64-encoded signature string
func (a *Authenticator) SignWebSocket(timestamp int64, method, path, body string) string {
	message := fmt.Sprintf("%d%s%s%s", timestamp, method, path, body)
	return a.sign(message)
}

// SignWebSocketAuth generates the HMAC SHA256 signature for WebSocket authentication
//
// The signature algorithm is:
//
//	Message = timestamp + requestPath
//	Signature = Base64(HMAC-SHA256(secretKey, Message))
//
// Parameters:
//   - timestamp: Unix timestamp in milliseconds
//   - path: WebSocket path (e.g., "/v2/ws/private")
//
// Returns the base64-encoded signature string
func (a *Authenticator) SignWebSocketAuth(timestamp int64, path string) string {
	message := fmt.Sprintf("%d%s", timestamp, path)
	return a.sign(message)
}

// sign generates the HMAC SHA256 signature
func (a *Authenticator) sign(message string) string {
	h := hmac.New(sha256.New, []byte(a.secretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GetRESTHeaders returns the authentication headers for REST API requests
//
// Parameters:
//   - timestamp: Unix timestamp in milliseconds (if 0, current time is used)
//   - method: HTTP method (GET, POST, PUT, DELETE)
//   - path: Request path
//   - body: Request body as string
//
// Returns a map of header key-value pairs
func (a *Authenticator) GetRESTHeaders(timestamp int64, method, path, body string) map[string]string {
	if timestamp == 0 {
		timestamp = time.Now().UnixMilli()
	}

	signature := a.SignRequest(timestamp, method, path, body)

	return map[string]string{
		types.HeaderAccessKey:        a.apiKey,
		types.HeaderAccessSign:       signature,
		types.HeaderAccessPassphrase: a.passphrase,
		types.HeaderAccessTimestamp:  fmt.Sprintf("%d", timestamp),
		types.HeaderContentType:      types.ContentTypeJSON,
		types.HeaderUserAgent:        types.DefaultUserAgent,
	}
}

// GetWebSocketHeaders returns the authentication headers for WebSocket connections
//
// Parameters:
//   - timestamp: Unix timestamp in milliseconds (if 0, current time is used)
//   - path: WebSocket path (default: "/v2/ws/private")
//
// Returns a map of header key-value pairs
func (a *Authenticator) GetWebSocketHeaders(timestamp int64, path string) map[string]string {
	if timestamp == 0 {
		timestamp = time.Now().UnixMilli()
	}

	if path == "" {
		path = "/v2/ws/private"
	}

	signature := a.SignWebSocketAuth(timestamp, path)

	return map[string]string{
		types.HeaderAccessKey:        a.apiKey,
		types.HeaderAccessSign:       signature,
		types.HeaderAccessPassphrase: a.passphrase,
		types.HeaderAccessTimestamp:  fmt.Sprintf("%d", timestamp),
		types.HeaderUserAgent:        types.DefaultUserAgent,
	}
}

// GetAPIKey returns the API key
func (a *Authenticator) GetAPIKey() string {
	return a.apiKey
}

// GetPassphrase returns the passphrase
func (a *Authenticator) GetPassphrase() string {
	return a.passphrase
}

// ValidateTimestamp checks if a timestamp is within acceptable range
// The timestamp should be within 30 seconds of the current time
func ValidateTimestamp(timestamp int64) error {
	now := time.Now().UnixMilli()
	diff := now - timestamp
	if diff < 0 {
		diff = -diff
	}

	// Allow 30 seconds clock skew
	if diff > 30000 {
		return fmt.Errorf("timestamp %d is too far from current time %d (diff: %dms)", timestamp, now, diff)
	}

	return nil
}
