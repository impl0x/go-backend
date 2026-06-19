package router

import (
	mo "github.com/impl0x/go-backend"
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