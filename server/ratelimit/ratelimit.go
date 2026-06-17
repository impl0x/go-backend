package ratelimit

import (
	"go-backend/server/core/core_types"
	"go-backend/server/core/helper"
	"go-backend/server/ratelimit/ratelimiters"
	"net/http"
)

type Ratelimit struct {
	StatusCode   uint16    // default to 429
	ErrorMessage string // default to "Too many requests!"
}

// returns a middleware which implements the ratelimiter
func (r *Ratelimit) NewRatelimiter(rl ratelimiters.Ratelimiter) core_types.Middleware {
	// validation logic
	if r.ErrorMessage == "" {
		r.ErrorMessage = "Too many requests!"
	} else if r.StatusCode == 0 {
		if r.StatusCode > 599 {
			panic("Status code cannot be greater than 599")
		}
		r.StatusCode = 429
	}

	return helper.NewMiddleware(
		func(next http.Handler, w http.ResponseWriter, r *http.Request){
			if !rl.Allow(r){
				// TODO: return a 429 and ErrorMessage in proper format
			}
			next.ServeHTTP(w,r)
		},
	)
}
