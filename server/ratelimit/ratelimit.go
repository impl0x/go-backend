package ratelimit

import (
	"go-backend/server/core/core_types"
	"go-backend/server/core/helper"
	"go-backend/server/ratelimit/ratelimiters"
	"net/http"
)

type RatelimitConfig struct {
	StatusCode   int    // default to 429
	ErrorMessage string // default to "Too many requests!"
}

// returns a middleware which implements the ratelimiter
func NewRatelimiter(rl ratelimiters.Ratelimiter) (core_types.Middleware, *RatelimitConfig) {
	newRatelimitConfig := RatelimitConfig{
		StatusCode:   http.StatusTooManyRequests,
		ErrorMessage: "Too many requests",
	}

	return helper.NewMiddleware(
			func(next http.Handler, w http.ResponseWriter, r *http.Request) {
				if !rl.Allow(r) {
					// TODO: return a 429 and ErrorMessage in proper format
				}
				next.ServeHTTP(w, r)
			},
		),
		&newRatelimitConfig
}
