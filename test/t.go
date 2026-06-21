package main

import "github.com/impl0x/mo"

func main() {
	m:=mo.New()
	m.GET("/",func(c *mo.Context) error {
		c.Mo.HTTPErrorHandler(c, nil)
	})
}
