package main

import (
	"fmt"
	"net/http"
)

type Engine struct {

}

func (e Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	default:
		writer.Write([]byte("jack"))
		
	}
}

func main()  {
	engine := &Engine{}
	http.ListenAndServe(":9999", engine)

}

