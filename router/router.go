package router

import (
	"net/http"

	"github.com/impl0x/mo"
)

type Router struct {
	routes Routes
	PreMiddlewares []mo.HandlerFunc
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	context:=mo.NewRequestContext(w,r)

	// applying pre middlewares
	var httpErr *mo.HttpError
	for _,val := range ro.PreMiddlewares{
		httpErr=val(context)
		if httpErr!=nil{
			return
		}
	}

	// routing
	// TODO

	// serving to handler
	// TODO, ex: ro.routes[0].Handler(context)
	
	
}



