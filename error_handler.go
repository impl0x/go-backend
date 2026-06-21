package mo

import "net/http"

type HTTPErrorHandler func(*Context, error)

func DefaultHTTPErrorHandler(exposeError bool) HTTPErrorHandler {
	return func(c *Context, err error) {
		if c.response.committed{
			return
		}
		switch e := err.(type) {
		case HttpErrorInterface:
			c.JSON(e.StatusCode(), e.JsonFormat())
		default:
			resp := map[string]any{
				"message": http.StatusText(http.StatusInternalServerError),
			}
			if exposeError{
				resp["error"]=e.Error()
			}
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}
