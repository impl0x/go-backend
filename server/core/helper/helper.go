package helper

import (
	"go-backend/server/core/core_types"
	"net/http"
)

func Middleware(Handler core_types.MiddlewareHandler) core_types.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Handler(next, w, r)
		})
	}
}
