package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/muathendirangu/ratelimiter"
)

func TestNewRateLimiter(t *testing.T) {
	threshold := int(10)
	duration := time.Second

	rl := ratelimiter.NewRateLimiter(
		ratelimiter.WithThreshold(threshold),
		ratelimiter.WithDuration(duration),
	)
	if rl.Threshold != threshold {
		t.Errorf("Expected threshold to be %d, got %d", threshold, rl.Threshold)
	}
	if rl.Duration != duration {
		t.Errorf("Expected duration to be %d, got %d", duration, rl.Duration)
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	rl := ratelimiter.NewRateLimiter(
		ratelimiter.WithThreshold(10),
		ratelimiter.WithDuration(time.Second),
	)
	for i := 0; i < 10; i++ {
		if !rl.Allow() {
			t.Error("Expected request to be allowed")
		}
	}
	if rl.Allow() {
		t.Error("Expected request to be denied")
	}
	time.Sleep(time.Second)
	if !rl.Allow() {
		t.Error("Expected request to be allowed after reset")
	}
}
