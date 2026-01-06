package weex

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

// Retrier handles retry logic with exponential backoff
type Retrier struct {
	maxRetries     int
	initialBackoff time.Duration
	maxBackoff     time.Duration
	backoffFactor  float64
	logger         Logger
}

// NewRetrier creates a new Retrier instance
func NewRetrier(maxRetries int, initialBackoff, maxBackoff time.Duration, backoffFactor float64, logger Logger) *Retrier {
	return &Retrier{
		maxRetries:     maxRetries,
		initialBackoff: initialBackoff,
		maxBackoff:     maxBackoff,
		backoffFactor:  backoffFactor,
		logger:         logger,
	}
}

// DoWithRetry executes a function with retry logic
//
// The function will be retried if:
//   - It returns a retriable error (APIError with IsRetriable() == true)
//   - It returns a NetworkError
//   - The context is not canceled
//
// Parameters:
//   - ctx: Context for cancellation
//   - fn: Function to execute
//
// Returns the error from the last attempt if all retries fail
func (r *Retrier) DoWithRetry(ctx context.Context, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		// Check context before attempting
		select {
		case <-ctx.Done():
			r.logger.Debug("Context canceled, stopping retries")
			return ctx.Err()
		default:
		}

		// Execute the function
		err := fn()
		if err == nil {
			// Success
			if attempt > 0 {
				r.logger.Info("Request succeeded after %d retries", attempt)
			}
			return nil
		}

		lastErr = err

		// Check if error is retriable
		if !r.isRetriable(err) {
			r.logger.Debug("Error is not retriable: %v", err)
			return err
		}

		// Don't sleep after the last attempt
		if attempt >= r.maxRetries {
			r.logger.Warn("Max retries (%d) exceeded, giving up", r.maxRetries)
			break
		}

		// Calculate backoff duration
		backoff := r.calculateBackoff(attempt)
		r.logger.Info("Request failed (attempt %d/%d), retrying after %v: %v",
			attempt+1, r.maxRetries+1, backoff, err)

		// Wait with context support
		select {
		case <-time.After(backoff):
			// Continue to next retry
		case <-ctx.Done():
			r.logger.Debug("Context canceled during backoff")
			return ctx.Err()
		}
	}

	return fmt.Errorf("%w: %v", ErrMaxRetriesExceeded, lastErr)
}

// isRetriable determines if an error is retriable
func (r *Retrier) isRetriable(err error) bool {
	if err == nil {
		return false
	}

	// Check for context errors (not retriable)
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	// Check for APIError
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.IsRetriable()
	}

	// Check for NetworkError (always retriable)
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return true
	}

	// Default: not retriable
	return false
}

// calculateBackoff calculates the backoff duration for a given attempt
// Uses exponential backoff: initialBackoff * (backoffFactor ^ attempt)
func (r *Retrier) calculateBackoff(attempt int) time.Duration {
	backoff := float64(r.initialBackoff) * math.Pow(r.backoffFactor, float64(attempt))

	// Cap at maxBackoff
	if backoff > float64(r.maxBackoff) {
		backoff = float64(r.maxBackoff)
	}

	return time.Duration(backoff)
}

// ShouldRetry is a helper function to check if an error should be retried
func ShouldRetry(err error) bool {
	if err == nil {
		return false
	}

	// Context errors are not retriable
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	// Check APIError
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.IsRetriable()
	}

	// Check NetworkError
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return true
	}

	return false
}
