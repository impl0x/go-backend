package mo

import "net/http"

type Grouped struct {
	prefix      string
	Middlewares []Middleware
	m           *Mo
}

func (g *Grouped) add(path string, method string, handler HandlerFunc, mi []Middleware) *Route {
	return g.m.add(g.prefix+path, method, handler, append(g.Middlewares, mi...))
}

func (g *Grouped) GET(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodGet, handler, mi)
}
func (g *Grouped) POST(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodPost, handler, mi)
}
func (g *Grouped) PATCH(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodPatch, handler, mi)
}
func (g *Grouped) PUT(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodPut, handler, mi)
}
func (g *Grouped) OPTIONS(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodOptions, handler, mi)
}
func (g *Grouped) DELETE(path string, handler HandlerFunc, mi ...Middleware) *Route {
	return g.add(path, http.MethodDelete, handler, mi)
}
