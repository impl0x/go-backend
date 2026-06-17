package core_types

import "net/http"

type Middleware func(next http.Handler)http.Handler

type MiddlewareHandler func(next http.Handler, w http.ResponseWriter, r *http.Request)
