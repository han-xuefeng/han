package main

import (
	"gee"
	"net/http"
)

func main()  {
	r := gee.New()
	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//
	//r.GET("/hello", func(c *gee.Context) {
	//	// expect /hello?name=geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})
	//
	//r.GET("/hello/:name", func(c *gee.Context) {
	//	// expect /hello/geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//
	//r.GET("/assets/*filepath", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	//})

	group := r.Group("vi")
	group.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page#########</h1>")
	})

	r.Run(":9999")
}
