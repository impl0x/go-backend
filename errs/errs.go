package errs

import "net/http"

type HttpNotFoundError struct {
	StatusCode int
	Message string
}

func (h *HttpNotFoundError) Error() string {
	return h.Message
}
func NewHttpNotFoundError() *HttpNotFoundError {
	return &HttpNotFoundError{
		StatusCode: http.StatusNotFound,
		Message: "Not found",
	}
}

type HttpInternalServerError struct {
	StatusCode int
	Message string
}
func (h *HttpInternalServerError) Error() string {
	return h.Message
}
func NewHttpInternalServerError() *HttpInternalServerError {
	return &HttpInternalServerError{
		StatusCode: http.StatusInternalServerError,
		Message: "Internal server error",
	}
}