package router

type httpMethod string

const (
	GET     httpMethod = "GET"
	HEAD    httpMethod = "HEAD"
	POST    httpMethod = "POST"
	PUT     httpMethod = "PUT"
	PATCH   httpMethod = "PATCH"
	DELETE  httpMethod = "DELETE"
	CONNECT httpMethod = "CONNECT"
	OPTIONS httpMethod = "OPTIONS"
	TRACE   httpMethod = "TRACE"
)
