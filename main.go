package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on 0.0.0.0:8888")

	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("0.0.0.0:8888", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// Print request method and URL
	fmt.Printf("%s %s\n", r.Method, r.URL.String())

	// Print all headers
	for n, v := range r.Header {
		fmt.Printf("%s: %s\n", n, v)
	}
	fmt.Println()
}
