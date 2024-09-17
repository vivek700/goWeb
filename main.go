package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website")
		// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
