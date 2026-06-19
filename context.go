package mo

import "net/http"

type Context struct {
	request *http.Request
	response http.ResponseWriter
}