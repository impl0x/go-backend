package mo


// type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c *Context) *HttpError

type Mo struct{

}


