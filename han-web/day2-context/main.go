package main

import (
	"gee"
	"net/http"
)

func main()  {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "html")
	})
	engine.Run(":9999")
}
