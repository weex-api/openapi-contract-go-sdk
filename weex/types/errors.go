package types

// ErrorType represents the category of an error
type ErrorType int

const (
	ErrTypeUnknown    ErrorType = iota // Unknown error
	ErrTypeAuth                        // Authentication error (40001-40014, 40753)
	ErrTypeRateLimit                   // Rate limiting error (429)
	ErrTypeValidation                  // Validation error (40017, 40019, 40020)
	ErrTypeNetwork                     // Network error (connection, timeout, etc.)
	ErrTypeSystem                      // System error (50xxx)
	ErrTypePermission                  // Permission error (40022, 50003, 50004)
	ErrTypeBusiness                    // Business logic error
)

// String returns the string representation of ErrorType
func (e ErrorType) String() string {
	switch e {
	case ErrTypeAuth:
		return "AUTHENTICATION_ERROR"
	case ErrTypeRateLimit:
		return "RATE_LIMIT_ERROR"
	case ErrTypeValidation:
		return "VALIDATION_ERROR"
	case ErrTypeNetwork:
		return "NETWORK_ERROR"
	case ErrTypeSystem:
		return "SYSTEM_ERROR"
	case ErrTypePermission:
		return "PERMISSION_ERROR"
	case ErrTypeBusiness:
		return "BUSINESS_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}

// ErrorCategory contains error classification information
type ErrorCategory struct {
	Type      ErrorType // Error type
	Retriable bool      // Whether the error is retriable
}

// Common error codes mapping to categories
// Reference: /contract/ErrorCodes/ExampleOfErrorCode.md
var ErrorCodeMap = map[string]*ErrorCategory{
	// Authentication errors (not retriable)
	"40001": {Type: ErrTypeAuth, Retriable: false}, // ACCESS-KEY header cannot be empty
	"40002": {Type: ErrTypeAuth, Retriable: false}, // ACCESS-SIGN header cannot be empty
	"40003": {Type: ErrTypeAuth, Retriable: false}, // ACCESS-PASSPHRASE header cannot be empty
	"40004": {Type: ErrTypeAuth, Retriable: false}, // ACCESS-TIMESTAMP header cannot be empty
	"40005": {Type: ErrTypeAuth, Retriable: false}, // Invalid ACCESS-TIMESTAMP
	"40006": {Type: ErrTypeAuth, Retriable: false}, // Invalid API key
	"40007": {Type: ErrTypeAuth, Retriable: false}, // Invalid signature
	"40008": {Type: ErrTypeAuth, Retriable: false}, // Timestamp expired (>30s difference)
	"40009": {Type: ErrTypeAuth, Retriable: false}, // API key doesn't exist
	"40010": {Type: ErrTypeAuth, Retriable: false}, // Incorrect passphrase
	"40011": {Type: ErrTypeAuth, Retriable: false}, // API key expired
	"40012": {Type: ErrTypeAuth, Retriable: false}, // API key frozen
	"40013": {Type: ErrTypeAuth, Retriable: false}, // IP not in whitelist
	"40014": {Type: ErrTypeAuth, Retriable: false}, // API key not bound to subaccount
	"40753": {Type: ErrTypeAuth, Retriable: false}, // Invalid locale parameter

	// Permission errors (not retriable)
	"40022": {Type: ErrTypePermission, Retriable: false}, // Insufficient permissions
	"50003": {Type: ErrTypePermission, Retriable: false}, // Operation not allowed
	"50004": {Type: ErrTypePermission, Retriable: false}, // Endpoint access forbidden

	// Validation errors (not retriable)
	"40017": {Type: ErrTypeValidation, Retriable: false}, // Parameter validation failed
	"40019": {Type: ErrTypeValidation, Retriable: false}, // Missing required parameter
	"40020": {Type: ErrTypeValidation, Retriable: false}, // Invalid parameter value

	// Rate limiting (retriable after delay)
	"429": {Type: ErrTypeRateLimit, Retriable: true}, // Too many requests

	// System errors (retriable)
	"40015": {Type: ErrTypeSystem, Retriable: true}, // System error
	"40018": {Type: ErrTypeSystem, Retriable: true}, // Invalid IP address
	"50000": {Type: ErrTypeSystem, Retriable: true}, // Internal server error
	"50001": {Type: ErrTypeSystem, Retriable: true}, // Service temporarily unavailable
	"50002": {Type: ErrTypeSystem, Retriable: true}, // Service degradation

	// Business errors (usually not retriable)
	"50005": {Type: ErrTypeBusiness, Retriable: false}, // Order not found
	"50006": {Type: ErrTypeBusiness, Retriable: false}, // Position not found
	"50007": {Type: ErrTypeBusiness, Retriable: false}, // Leverage exceeds limit
	"50008": {Type: ErrTypeBusiness, Retriable: false}, // Insufficient balance
	"50009": {Type: ErrTypeBusiness, Retriable: false}, // Position size exceeds limit
	"50010": {Type: ErrTypeBusiness, Retriable: false}, // Risk limit exceeded
}

// GetErrorCategory returns the error category for a given error code
func GetErrorCategory(code string) *ErrorCategory {
	if cat, ok := ErrorCodeMap[code]; ok {
		return cat
	}
	// Default to unknown, not retriable
	return &ErrorCategory{Type: ErrTypeUnknown, Retriable: false}
}

// IsRetriableError checks if an error code represents a retriable error
func IsRetriableError(code string) bool {
	cat := GetErrorCategory(code)
	return cat.Retriable
}

// IsAuthError checks if an error code represents an authentication error
func IsAuthError(code string) bool {
	cat := GetErrorCategory(code)
	return cat.Type == ErrTypeAuth
}

// IsRateLimitError checks if an error code represents a rate limiting error
func IsRateLimitError(code string) bool {
	cat := GetErrorCategory(code)
	return cat.Type == ErrTypeRateLimit
}
