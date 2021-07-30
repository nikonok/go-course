package main

import (
	"fmt"
	"net/http"
	"log"
)

var (
	// our message
	message = "Hello there!"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
		case http.MethodGet:
			fmt.Fprint(w, message)
		case http.MethodPost:
			// get new value for message
			message = r.PostFormValue("hello-msg")
			fmt.Fprint(w, message)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	// server port number
	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
