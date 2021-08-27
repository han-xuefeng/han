package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Static("/static", "F:\\code\\go\\han\\han-web\\day6-template\\static")
	// 或相对路径 r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title" : "han",
			"name" : "thanks geektutu",
		})
	})

	r.Run(":9999")
}