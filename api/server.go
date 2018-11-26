package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartServer starts our server with the specified port (default: 8080)
func StartServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetListHandler).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")

	r.Handle("/books", authMiddleware(NewBookHandler())).Methods("POST")
	r.Handle("/books/{id}", authMiddleware(UpdateBookHandler())).Methods("PUT")
	r.Handle("/books/{id}", authMiddleware(DeleteBookHandler())).Methods("DELETE")

	listeningPort := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(listeningPort, r); err != nil {
		log.Fatal(err)
	}
}
