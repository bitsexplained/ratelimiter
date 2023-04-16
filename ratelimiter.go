package ratelimiter

import "time"

// RateLimiterOption is a functional option for configuring RateLimiter.
type RateLimiterOption func(*RateLimiter)

// RateLimiter defines a simple rate limiter.
type RateLimiter struct {
	threshold    int
	duration     time.Duration
	lastReset    time.Time
	requestCount int
}

// WithThreshold sets the request threshold for the rate limiter.
func WithThreshold(threshold int) RateLimiterOption {
	return func(rl *RateLimiter) {
		rl.threshold = threshold
	}
}

// WithDuration sets the duration for which the request threshold is allowed.
func WithDuration(duration time.Duration) RateLimiterOption {
	return func(rl *RateLimiter) {
		rl.duration = duration
	}
}

// NewRateLimiter creates a new rate limiter with the provided options.
func NewRateLimiter(options ...RateLimiterOption) *RateLimiter {
	rl := &RateLimiter{
		threshold:    500,
		duration:     time.Second,
		lastReset:    time.Now(),
		requestCount: 0,
	}

	// Apply options
	for _, opt := range options {
		opt(rl)
	}

	return rl
}

// Allow checks if a request is allowed based on the rate limiter's settings.
func (rl *RateLimiter) Allow() bool {
	// Check if time duration since last reset is greater than the duration
	if time.Since(rl.lastReset) >= rl.duration {
		// Reset request count
		rl.requestCount = 0
		rl.lastReset = time.Now()
	}

	// Check if request count is below threshold
	if rl.requestCount < rl.threshold {
		// Increment request count and allow request
		rl.requestCount++
		return true
	}

	return false
}

// Reset resets the rate limiter by setting the request count to 0 and updating the last reset time.
func (rl *RateLimiter) Reset() {
	rl.requestCount = 0
	rl.lastReset = time.Now()
}
