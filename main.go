package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	fmt.Println("Server starting...")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website")
		// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}).Methods("GET")

	// fs := http.FileServer(http.Dir("static/"))

	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on: 8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
