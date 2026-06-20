package mo

type Route struct {
	Path    string
	Method  string
	Handler HandlerFunc
	Middlewares []Middleware
}
type Router struct {
	Routes []Route
}

var emptyRoute = Route{}

func (r *Router) Route(path string, method string) (Route, HttpError) {
	for _, v := range r.Routes {
		if path == v.Path {
			if method == v.Method {
				return v, nil
			}
			return emptyRoute, ErrMethodNotAllowed
		}
	}
	return emptyRoute, ErrNotFound
}