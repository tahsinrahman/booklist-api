package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetListHandler).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")

	r.Handle("/books", authMiddleware(NewBookHandler())).Methods("POST")
	r.Handle("/books/{id}", authMiddleware(UpdateBookHandler())).Methods("PUT")
	r.Handle("/books/{id}", authMiddleware(DeleteBookHandler())).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
