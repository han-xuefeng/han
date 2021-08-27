package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	http.ListenAndServe(":9990", nil)
}
