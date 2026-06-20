package main

import "github.com/impl0x/mo"

// import "github.com/labstack/echo/v5"

func main() {
	n:=mo.New()
	n.GET("/",func(c *mo.Context) error {println("Root test");return nil})
	n.Start(":8080")
}