package main

import (
	"fmt"
	"net/http"
)

func main() {
	const URLPrefix = "http://"
	const DOMAIN = "localhost"
	const PORT = ":8080"

	http.HandleFunc("/", helloHandler)
	fmt.Printf("Server starting!\nOpen %s%s%s in your browser.", URLPrefix, DOMAIN, PORT)

	err := http.ListenAndServe(DOMAIN+PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Bonjour Hello World!")
}
