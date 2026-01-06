package weex

import (
	"fmt"
	"net/http"
	"time"

	"github.com/weex-api/openapi-contract-go-sdk/weex/rest"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/account"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/market"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/trade"
)

// Client is the main SDK client for WEEX Contract API
type Client struct {
	config *Config
	auth   *Authenticator
	rest   *rest.Client
	logger Logger

	// Service accessors (lazy initialization)
	marketService  *market.Service
	accountService *account.Service
	tradeService   *trade.Service
}

// NewClient creates a new WEEX Contract API client
//
// Example:
//
//	config := weex.NewDefaultConfig().
//	    WithAPIKey("your-api-key").
//	    WithSecretKey("your-secret-key").
//	    WithPassphrase("your-passphrase")
//
//	client := weex.NewClient(config)
//	ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
func NewClient(config *Config) (*Client, error) {
	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create authenticator
	auth := NewAuthenticator(config.APIKey, config.SecretKey, config.Passphrase)

	// Create HTTP client
	httpClient := &http.Client{
		Timeout: config.HTTPTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Create retrier
	retrier := NewRetrier(
		config.MaxRetries,
		config.InitialBackoff,
		config.MaxBackoff,
		config.BackoffFactor,
		config.Logger,
	)

	// Create rate limiter
	rateLimiter := NewRateLimiter(
		config.EnableRateLimit,
		config.IPWeight,
		config.UIDWeight,
		config.Logger,
	)

	// Create REST client
	restClient := rest.NewClient(
		config.BaseURL,
		config.Locale,
		httpClient,
		auth,
		retrier,
		rateLimiter,
		config.Logger,
	)

	return &Client{
		config: config,
		auth:   auth,
		rest:   restClient,
		logger: config.Logger,
	}, nil
}

// NewPublicClient creates a new client for public endpoints only
// This client does not require API credentials
//
// Example:
//
//	config := weex.NewDefaultConfig()
//	client, err := weex.NewPublicClient(config)
//	ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
func NewPublicClient(config *Config) (*Client, error) {
	// Validate public configuration (no credentials required)
	if err := config.ValidatePublic(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create empty authenticator for public endpoints
	auth := NewAuthenticator("", "", "")

	// Create HTTP client
	httpClient := &http.Client{
		Timeout: config.HTTPTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Create retrier
	retrier := NewRetrier(
		config.MaxRetries,
		config.InitialBackoff,
		config.MaxBackoff,
		config.BackoffFactor,
		config.Logger,
	)

	// Create rate limiter
	rateLimiter := NewRateLimiter(
		config.EnableRateLimit,
		config.IPWeight,
		config.UIDWeight,
		config.Logger,
	)

	// Create REST client
	restClient := rest.NewClient(
		config.BaseURL,
		config.Locale,
		httpClient,
		auth,
		retrier,
		rateLimiter,
		config.Logger,
	)

	return &Client{
		config: config,
		auth:   auth,
		rest:   restClient,
		logger: config.Logger,
	}, nil
}

// Market returns the market data service
// Provides access to public market data endpoints
func (c *Client) Market() *market.Service {
	if c.marketService == nil {
		c.marketService = market.NewService(c.rest)
	}
	return c.marketService
}

// Account returns the account management service
// Provides access to account and position endpoints (requires authentication)
func (c *Client) Account() *account.Service {
	if c.accountService == nil {
		c.accountService = account.NewService(c.rest)
	}
	return c.accountService
}

// Trade returns the trading service
// Provides access to order and trading endpoints (requires authentication)
func (c *Client) Trade() *trade.Service {
	if c.tradeService == nil {
		c.tradeService = trade.NewService(c.rest)
	}
	return c.tradeService
}

// GetConfig returns a copy of the client configuration
func (c *Client) GetConfig() *Config {
	return c.config.Clone()
}

// GetLogger returns the logger instance
func (c *Client) GetLogger() Logger {
	return c.logger
}

// SetLogLevel sets the log level for the client
func (c *Client) SetLogLevel(level LogLevel) {
	c.logger.SetLevel(level)
}
