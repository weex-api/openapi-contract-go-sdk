package weex

import (
	"fmt"
	"time"

	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// Config holds the configuration for the WEEX Contract API client
type Config struct {
	// API credentials
	APIKey     string // API key
	SecretKey  string // Secret key for signing
	Passphrase string // API key passphrase

	// API endpoints
	BaseURL      string // REST API base URL (default: https://api-contract.weex.com)
	WSPublicURL  string // WebSocket public URL (default: wss://ws-contract.weex.com/v2/ws/public)
	WSPrivateURL string // WebSocket private URL (default: wss://ws-contract.weex.com/v2/ws/private)

	// HTTP client settings
	HTTPTimeout time.Duration // HTTP request timeout (default: 10 seconds)
	MaxRetries  int           // Maximum number of retries for failed requests (default: 3)

	// Rate limiting
	EnableRateLimit bool // Enable rate limiting (default: true)
	IPWeight        int  // Max IP weight per 5 minutes (default: 300)
	UIDWeight       int  // Max UID weight per 5 minutes (default: 100)

	// Retry settings
	InitialBackoff time.Duration // Initial backoff duration for retries (default: 1 second)
	MaxBackoff     time.Duration // Maximum backoff duration for retries (default: 30 seconds)
	BackoffFactor  float64       // Backoff multiplier (default: 2.0)

	// WebSocket settings
	WSReadBufferSize  int           // WebSocket read buffer size (default: 4096)
	WSWriteBufferSize int           // WebSocket write buffer size (default: 4096)
	WSPingInterval    time.Duration // WebSocket ping interval (default: 20 seconds)
	WSPongWait        time.Duration // WebSocket pong wait time (default: 30 seconds)
	WSReconnect       bool          // Enable automatic reconnection (default: true)
	WSMaxReconnect    int           // Maximum reconnection attempts (default: 10)
	WSReconnectDelay  time.Duration // Initial reconnection delay (default: 1 second)

	// Logging
	Logger   Logger   // Custom logger (default: DefaultLogger with Info level)
	LogLevel LogLevel // Log level (default: Info)

	// Locale
	Locale string // API locale (default: "en")
}

// NewDefaultConfig creates a new Config with default values
func NewDefaultConfig() *Config {
	return &Config{
		BaseURL:      types.DefaultBaseURL,
		WSPublicURL:  types.DefaultWSPublicURL,
		WSPrivateURL: types.DefaultWSPrivateURL,

		HTTPTimeout: 10 * time.Second,
		MaxRetries:  3,

		EnableRateLimit: true,
		IPWeight:        300,
		UIDWeight:       100,

		InitialBackoff: 1 * time.Second,
		MaxBackoff:     30 * time.Second,
		BackoffFactor:  2.0,

		WSReadBufferSize:  4096,
		WSWriteBufferSize: 4096,
		WSPingInterval:    20 * time.Second,
		WSPongWait:        30 * time.Second,
		WSReconnect:       true,
		WSMaxReconnect:    10,
		WSReconnectDelay:  1 * time.Second,

		Logger:   NewDefaultLogger(LogLevelInfo),
		LogLevel: LogLevelInfo,

		Locale: types.DefaultLocale,
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// API credentials validation (required for private endpoints)
	if c.APIKey == "" || c.SecretKey == "" || c.Passphrase == "" {
		return fmt.Errorf("%w: APIKey, SecretKey, and Passphrase are required for authenticated requests", ErrInvalidConfig)
	}

	// URL validation
	if c.BaseURL == "" {
		return fmt.Errorf("%w: BaseURL cannot be empty", ErrInvalidConfig)
	}

	// Timeout validation
	if c.HTTPTimeout <= 0 {
		return fmt.Errorf("%w: HTTPTimeout must be greater than 0", ErrInvalidConfig)
	}

	// Retry validation
	if c.MaxRetries < 0 {
		return fmt.Errorf("%w: MaxRetries cannot be negative", ErrInvalidConfig)
	}

	// Backoff validation
	if c.InitialBackoff <= 0 {
		return fmt.Errorf("%w: InitialBackoff must be greater than 0", ErrInvalidConfig)
	}
	if c.MaxBackoff <= 0 {
		return fmt.Errorf("%w: MaxBackoff must be greater than 0", ErrInvalidConfig)
	}
	if c.BackoffFactor <= 1.0 {
		return fmt.Errorf("%w: BackoffFactor must be greater than 1.0", ErrInvalidConfig)
	}

	// Logger validation
	if c.Logger == nil {
		c.Logger = NewDefaultLogger(c.LogLevel)
	}

	return nil
}

// ValidatePublic checks if the configuration is valid for public endpoints only
// Public endpoints don't require API credentials
func (c *Config) ValidatePublic() error {
	// URL validation
	if c.BaseURL == "" {
		return fmt.Errorf("%w: BaseURL cannot be empty", ErrInvalidConfig)
	}

	// Timeout validation
	if c.HTTPTimeout <= 0 {
		return fmt.Errorf("%w: HTTPTimeout must be greater than 0", ErrInvalidConfig)
	}

	// Logger validation
	if c.Logger == nil {
		c.Logger = NewDefaultLogger(c.LogLevel)
	}

	return nil
}

// Clone creates a copy of the configuration
func (c *Config) Clone() *Config {
	clone := *c
	return &clone
}

// WithAPIKey sets the API key and returns the config for chaining
func (c *Config) WithAPIKey(apiKey string) *Config {
	c.APIKey = apiKey
	return c
}

// WithSecretKey sets the secret key and returns the config for chaining
func (c *Config) WithSecretKey(secretKey string) *Config {
	c.SecretKey = secretKey
	return c
}

// WithPassphrase sets the passphrase and returns the config for chaining
func (c *Config) WithPassphrase(passphrase string) *Config {
	c.Passphrase = passphrase
	return c
}

// WithBaseURL sets the base URL and returns the config for chaining
func (c *Config) WithBaseURL(baseURL string) *Config {
	c.BaseURL = baseURL
	return c
}

// WithHTTPTimeout sets the HTTP timeout and returns the config for chaining
func (c *Config) WithHTTPTimeout(timeout time.Duration) *Config {
	c.HTTPTimeout = timeout
	return c
}

// WithMaxRetries sets the maximum retries and returns the config for chaining
func (c *Config) WithMaxRetries(maxRetries int) *Config {
	c.MaxRetries = maxRetries
	return c
}

// WithLogger sets the logger and returns the config for chaining
func (c *Config) WithLogger(logger Logger) *Config {
	c.Logger = logger
	return c
}

// WithLogLevel sets the log level and returns the config for chaining
func (c *Config) WithLogLevel(level LogLevel) *Config {
	c.LogLevel = level
	if c.Logger != nil {
		c.Logger.SetLevel(level)
	}
	return c
}

// WithLocale sets the locale and returns the config for chaining
func (c *Config) WithLocale(locale string) *Config {
	c.Locale = locale
	return c
}
