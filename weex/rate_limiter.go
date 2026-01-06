package weex

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TokenBucket implements a token bucket rate limiter
type TokenBucket struct {
	capacity       int           // Maximum number of tokens
	tokens         int           // Current number of tokens
	refillRate     int           // Tokens to add per refill interval
	refillInterval time.Duration // How often to refill tokens
	lastRefill     time.Time     // Last refill time
	mu             sync.Mutex    // Mutex for thread safety
}

// NewTokenBucket creates a new TokenBucket
//
// Parameters:
//   - capacity: Maximum tokens (e.g., 300 for IP weight, 100 for UID weight)
//   - refillInterval: Time window for refill (e.g., 5 minutes)
//
// The bucket refills to full capacity after each interval
func NewTokenBucket(capacity int, refillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     capacity,
		refillInterval: refillInterval,
		lastRefill:     time.Now(),
	}
}

// Take attempts to take n tokens from the bucket
// Returns true if successful, false if not enough tokens
func (tb *TokenBucket) Take(n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}

	return false
}

// Wait waits until n tokens are available, respecting context cancellation
// Returns error if context is canceled or deadline exceeded
func (tb *TokenBucket) Wait(ctx context.Context, n int) error {
	// Fast path: tokens available immediately
	if tb.Take(n) {
		return nil
	}

	// Need to wait
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if tb.Take(n) {
				return nil
			}
		}
	}
}

// refill adds tokens based on elapsed time since last refill
// Must be called with mutex held
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	if elapsed >= tb.refillInterval {
		// Full refill
		tb.tokens = tb.capacity
		tb.lastRefill = now
	}
}

// Available returns the number of tokens currently available
func (tb *TokenBucket) Available() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	return tb.tokens
}

// RateLimiter manages rate limiting using token buckets
type RateLimiter struct {
	ipBucket  *TokenBucket // IP weight limiter
	uidBucket *TokenBucket // UID weight limiter
	enabled   bool         // Whether rate limiting is enabled
	logger    Logger
}

// NewRateLimiter creates a new RateLimiter
//
// Parameters:
//   - enabled: Whether to enable rate limiting
//   - ipWeight: Maximum IP weight per 5 minutes (default: 300)
//   - uidWeight: Maximum UID weight per 5 minutes (default: 100)
//   - logger: Logger instance
func NewRateLimiter(enabled bool, ipWeight, uidWeight int, logger Logger) *RateLimiter {
	return &RateLimiter{
		ipBucket:  NewTokenBucket(ipWeight, 5*time.Second),
		uidBucket: NewTokenBucket(uidWeight, 5*time.Second),
		enabled:   enabled,
		logger:    logger,
	}
}

// WaitForCapacity waits until the specified weight is available
//
// Parameters:
//   - ctx: Context for cancellation
//   - ipWeight: IP weight for the request
//   - uidWeight: UID weight for the request
//
// Returns error if rate limit cannot be satisfied or context is canceled
func (rl *RateLimiter) WaitForCapacity(ctx context.Context, ipWeight, uidWeight int) error {
	if !rl.enabled {
		return nil
	}

	// Wait for IP capacity
	if ipWeight > 0 {
		rl.logger.Debug("Waiting for IP weight capacity: %d", ipWeight)
		if err := rl.ipBucket.Wait(ctx, ipWeight); err != nil {
			return fmt.Errorf("failed to acquire IP weight: %w", err)
		}
	}

	// Wait for UID capacity
	if uidWeight > 0 {
		rl.logger.Debug("Waiting for UID weight capacity: %d", uidWeight)
		if err := rl.uidBucket.Wait(ctx, uidWeight); err != nil {
			return fmt.Errorf("failed to acquire UID weight: %w", err)
		}
	}

	return nil
}

// TryAcquire attempts to acquire the specified weight without waiting
// Returns true if successful, false otherwise
func (rl *RateLimiter) TryAcquire(ipWeight, uidWeight int) bool {
	if !rl.enabled {
		return true
	}

	ipOk := true
	uidOk := true

	if ipWeight > 0 {
		ipOk = rl.ipBucket.Take(ipWeight)
	}

	if uidWeight > 0 {
		uidOk = rl.uidBucket.Take(uidWeight)
	}

	return ipOk && uidOk
}

// GetStatus returns the current status of the rate limiter
func (rl *RateLimiter) GetStatus() (ipAvailable, uidAvailable int) {
	return rl.ipBucket.Available(), rl.uidBucket.Available()
}
