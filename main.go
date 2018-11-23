package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

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

// Book defines a book object
type Book struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	ISBN    string   `json:"isbn,omitempty"`
	Authors []Author `json:"authors"`
}

// Author defines an author of a book
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
}

var count = 3
var books []Book
var mutex sync.Mutex

func init() {
	book1 := Book{
		ID:   1,
		Name: "Linux in a Nutshell",
		ISBN: "0596009305",
		Authors: []Author{
			Author{
				FirstName: "Ellen",
				LastName:  "Siever",
			},
			Author{
				FirstName: "Stephen",
				LastName:  "Figgins",
			},
		},
	}
	book2 := Book{
		ID:   2,
		Name: "The Linux Command Line",
		ISBN: "1593273894",
		Authors: []Author{
			Author{
				FirstName: "William",
				LastName:  "Scotts",
			},
		},
	}
	book3 := Book{
		ID:   3,
		Name: "The Linux Programming Interface",
		ISBN: "1593272200",
		Authors: []Author{
			Author{
				FirstName: "Michael",
				LastName:  "Kerrisk",
			},
		},
	}

	books = append(books, book1, book2, book3)
}

// GetListHandler is handler for GET /books
// this returns the list of all books
func GetListHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	b, err := json.Marshal(books)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// GetBookHandler is handler for GET /books/{id}
// returns a single book defined in {id}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, _ := findBook(id)

	if book == nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	b, err := json.Marshal(book)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// NewBookHandler is handler for POST /books
// adds a new book
func NewBookHandler(w http.ResponseWriter, r *http.Request) {
	book := new(Book)

	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}

	if err := checkNewBook(book); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	count++

	book.ID = count
	books = append(books, *book)

	b, err := json.Marshal(book)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

// UpdateBookHandler is handler for PUT /books/{id}
// updates the book information given by {id}
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	oldBook, index := findBook(id)
	if oldBook == nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	newBook := new(Book)
	if err := json.NewDecoder(r.Body).Decode(newBook); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	newBook.ID = oldBook.ID

	if err := checkNewBook(newBook); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	books[index] = *newBook

	b, err := json.Marshal(newBook)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// DeleteBookHandler is handler for DELETE /books/{id}
// deletes the book given by {id}
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, index := findBook(id)
	if book == nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	books = append(books[:index], books[index+1:]...)

	b, err := json.Marshal(book)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func findBook(id string) (*Book, int) {
	mutex.Lock()
	defer mutex.Unlock()

	for index, book := range books {
		if strconv.Itoa(book.ID) == id {
			return &book, index
		}
	}
	return nil, -1
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	message = fmt.Sprintf(`{"error": "%v"}`, message)

	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func checkNewBook(book *Book) error {
	if book.Name == "" {
		return errors.New("book name can't be empty")
	}
	for _, author := range book.Authors {
		if author.FirstName == "" {
			return errors.New("author firstname can't be empty")
		}
	}
	return nil
}
