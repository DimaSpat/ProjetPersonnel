package main

import (
	"fmt"
	"net/http"
)

func main() {
	const URLPrefix = "http://:"
	const DOMAIN = "localhost"
	const PORT = ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome to the Pong!\nAdd /ping to the search and get a response")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "pong")
	})

	fmt.Printf("Server starting!\nOpen %s%s%s in your browser.", URLPrefix, DOMAIN, PORT)
	err := http.ListenAndServe(DOMAIN+PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
