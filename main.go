package main

import (
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
