package ratelimiters

import (
	"net/http"
	"sync"
	"time"
)

type windowCounterConfig struct {
	limit      int
	windowSize time.Duration
}

type windowCounter struct {
	reqCount    int
	windowStart time.Time
	Config      windowCounterConfig
	mu          sync.RWMutex
}

// Non-IP based
// Fixed value per second, anything over is rejected.
func NewWindowCounter(maxRequests int, per time.Duration) *windowCounter {
	return &windowCounter{
		Config: windowCounterConfig{
			limit:      maxRequests,
			windowSize: per,
		},
		windowStart: time.Now(),
	}
}

func (wc *windowCounter) Allow(*http.Request) bool {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	if time.Since(wc.windowStart) >= wc.Config.windowSize {
		wc.reqCount = 0
		wc.windowStart = time.Now()
	}
	if wc.reqCount < wc.Config.limit {
		wc.reqCount++
		return true
	}
	return false
}
