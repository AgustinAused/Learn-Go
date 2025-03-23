package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Attach a path with a handler
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start the server
	http.ListenAndServe(":8080", r)
}