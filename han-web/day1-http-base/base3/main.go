package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main()  {
	engine := gee.New()
	engine.GET("/", indexHandler)
	err := engine.Run(":9999")
	fmt.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
