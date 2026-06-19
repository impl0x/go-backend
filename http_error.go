package mo

import (
	"net/http"

	"github.com/impl0x/mo/responses"
)

type HttpError struct {
	StatusCode   int
	Response responses.Response
}

func NewError() HttpError {
	return HttpError{
		http.StatusBadRequest,
		nil,
	}
}

func (h *HttpError) Write(c *Context){
	c.response.WriteHeader(h.StatusCode)
	// if h.ResponseBody
	// c.response.Write(h.ResponseBody)
}