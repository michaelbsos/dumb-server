package main

import (
	"fmt"
	"io/ioutil"
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

	// Print body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
