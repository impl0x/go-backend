package mo

import (
	"net/http"

	"github.com/impl0x/mo/errs"
)

type Mo struct{
	router *Router
}
func New()*Mo{
	return &Mo{
		router: &Router{},
	}
}
func (m *Mo)GET(path string, handler HandlerFunc){
	m.router.Routes = append(m.router.Routes, Route{path, http.MethodGet, handler})
}
func (m *Mo)Start(addr string)error{
	return http.ListenAndServe(":8080", m)
}




func (m *Mo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newContext := &Context{
		request:  r,
		response: w,
	}
	route,err:=m.router.Route(r.URL.Path)
	if err!=nil{
		println("Not found")
		// TODO: perform error handling function, default errorHandler gets called otherwise
	}
	route.Handler(newContext)
}


type Context struct {
	request  *http.Request
	response http.ResponseWriter
}


type Route struct {
	Path    string
	Method  string
	Handler HandlerFunc
}
type Router struct {
	Routes []Route
}
func (r *Router)Route(path string)(Route,error){
	for _,v:=range r.Routes{
		if path==v.Path{
			return v, nil
		}
	}
	return Route{},errs.NewHttpNotFoundError()
}

type HandlerFunc func(c *Context) error
type Middleware func(HandlerFunc) HandlerFunc

