package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetListHandler).Methods("GET")
	r.HandleFunc("/books", NewBookHandler).Methods("POST")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")
	r.HandleFunc("/books/{id}", UpdateBookHandler).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// GetListHandler is handler for GET /books
// this returns the list of all books
func GetListHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Host, r.Method); err != nil {
		log.Fatal(err)
	}
}

// GetBookHandler is handler for GET /books/{id}
// returns a single book defined in {id}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Host, r.Method); err != nil {
		log.Fatal(err)
	}
}

// NewBookHandler is handler for POST /books
// adds a new book
func NewBookHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Host, r.Method); err != nil {
		log.Fatal(err)
	}
}

// UpdateBookHandler is handler for PUT /books/{id}
// updates the book information given by {id}
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Host, r.Method); err != nil {
		log.Fatal(err)
	}
}

// DeleteBookHandler is handler for DELETE /books/{id}
// deletes the book given by {id}
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, r.Host, r.Method); err != nil {
		log.Fatal(err)
	}
}
