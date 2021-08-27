package main

import (
	"gee"
	"net/http"
)

func main()  {
	engine := gee.New()
	engine.GET("/hello", func (w http.ResponseWriter, req *http.Request) {

	})
	engine.Run(":9999")
}
