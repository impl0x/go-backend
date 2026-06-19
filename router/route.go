package router

import (
	"github.com/impl0x/mo"
)

type Route struct {
	Method      httpMethod
	Path        string
	Name        string
	Handler     mo.HandlerFunc
	PreMiddlewares []mo.MiddlewareFunc
	PostMiddlewares []mo.MiddlewareFunc
}

func t()  {
	
}