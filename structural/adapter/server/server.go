package main

import (
	"fmt"
	"log"
	"net/http"
)

//HandlerFunc is a adapter turn func to handler
//type HandlerFunc func(ResponseWriter, *Request)
//
//// ServeHTTP calls f(w, r).
//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//	f(w, r)
//}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
