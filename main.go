package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Index</h1>")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", health)
	fmt.Println("Server starting...")
	http.ListenAndServe(":5030", nil)
}
