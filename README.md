# Mo - api framework

A backend server template which is made on top of net/http and is lightweight

inspired heavily from echo

### Finished making !

## Features:
- Single context which gives access to Request and ResponseWriter objects
- Has middlewares for ratelimiting, 2 types as of now, token bucket algorithm and window counter
- Has in built validator which validates structs


## Documentation:
**Demo usage**
```go
import "github.com/impl0x/mo"

func main() {
	m := mo.New()           // New instance of Mo
	m.GET("/", rootHandler) // Registering a handler function for path "/"
	m.Start(":8080")        // start and serve at port 8080
}

func rootHandler(c *mo.Context) error { // signature for a mo.HandlerFunc
	println(`Got request at "/"`)
	return c.JSON(200, map[string]any{"status": "success"}) // returns a json response with status code 200 and {"status":"success"}
}
```
**Middleware usage**
```go
import (
	"github.com/impl0x/mo"
	"github.com/impl0x/mo/middlewares"
)

func main() {
	m := mo.New()
	m.Use(middlewares.Logger) // m.Use() takes a parameter of type mo.MiddlewareFunc
	m.GET("/", func(c *mo.Context) error {return nil})
	m.Start(":8080")
}
```
