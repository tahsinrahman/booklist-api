package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// StartServer starts our server with the specified port (default: 8080)
func StartServer(port int, timeout int) {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetListHandler).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")

	r.Handle("/books", authMiddleware(NewBookHandler())).Methods("POST")
	r.Handle("/books/{id}", authMiddleware(UpdateBookHandler())).Methods("PUT")
	r.Handle("/books/{id}", authMiddleware(DeleteBookHandler())).Methods("DELETE")

	listeningPort := fmt.Sprintf(":%v", port)
	server := &http.Server{
		Addr:    listeningPort,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	shutDownTime := time.Second * time.Duration(timeout)
	ctx, cancel := context.WithTimeout(context.Background(), shutDownTime)
	defer cancel()

	server.Shutdown(ctx)

	log.Println("Shutting Down")
}
