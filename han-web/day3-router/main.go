package main

import (
	"gee"
	"net/http"
)

func main()  {
	r := gee.New()
	r.GET("/a/b/c", func(c *gee.Context) {
		c.HTML(http.StatusOK, "abc")
	})
	r.Run(":9999")
}
