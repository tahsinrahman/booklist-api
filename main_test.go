package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/books", GetListHandler).Methods("GET")
	router.HandleFunc("/books", NewBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")
	router.HandleFunc("/books/{id}", UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
}

type test struct {
	method     string
	url        string
	request    string
	response   string
	statusCode int
}

func TestGetListHandler(t *testing.T) {
	testSuite := []test{
		test{
			method:     "GET",
			url:        "/books",
			request:    "",
			response:   `[{"id":1,"name":"Linux in a Nutshell","isbn":"0596009305","authors":[{"first_name":"Ellen","last_name":"Siever"},{"first_name":"Stephen","last_name":"Figgins"}]},{"id":2,"name":"The Linux Command Line","isbn":"1593273894","authors":[{"first_name":"William","last_name":"Scotts"}]},{"id":3,"name":"The Linux Programming Interface","isbn":"1593272200","authors":[{"first_name":"Michael","last_name":"Kerrisk"}]}]`,
			statusCode: 200,
		},
	}

	runTest(t, testSuite)
}

func TestGetBookHandler(t *testing.T) {
	testSuite := []test{
		test{
			method:     "GET",
			url:        "/books/1",
			request:    "",
			response:   `{"id":1,"name":"Linux in a Nutshell","isbn":"0596009305","authors":[{"first_name":"Ellen","last_name":"Siever"},{"first_name":"Stephen","last_name":"Figgins"}]}`,
			statusCode: 200,
		},
		test{
			method:     "GET",
			url:        "/books/4",
			request:    "",
			response:   `{"error": "book not found"}`,
			statusCode: 404,
		},
	}

	runTest(t, testSuite)
}

func TestNewBookHandler(t *testing.T) {
	testSuite := []test{
		test{
			method:     "POST",
			url:        "/books",
			request:    `{"name": "Understanding the Linux Kernel", "isbn":"0596005652", "authors": [{"first_name": "Daniel", "last_name": "Bovet"}, {"first_name": "Marco", "last_name": "cesati"}]}`,
			response:   `{"id":4,"name":"Understanding the Linux Kernel","isbn":"0596005652","authors":[{"first_name":"Daniel","last_name":"Bovet"},{"first_name":"Marco","last_name":"cesati"}]}`,
			statusCode: 201,
		},
		test{
			method:     "GET",
			url:        "/books/4",
			request:    "",
			response:   `{"id":4,"name":"Understanding the Linux Kernel","isbn":"0596005652","authors":[{"first_name":"Daniel","last_name":"Bovet"},{"first_name":"Marco","last_name":"cesati"}]}`,
			statusCode: 200,
		},
	}

	runTest(t, testSuite)
}

func TestUpdateBookHandler(t *testing.T) {
	testSuite := []test{
		test{
			method:     "PUT",
			url:        "/books/4",
			request:    `{"name": "Understanding the Linux Kernel", "isbn":"new_isbn", "authors": [{"first_name": "Daniel", "last_name": "Bovet"}, {"first_name": "Marco", "last_name": "cesati"}]}`,
			response:   `{"id":4,"name":"Understanding the Linux Kernel","isbn":"new_isbn","authors":[{"first_name":"Daniel","last_name":"Bovet"},{"first_name":"Marco","last_name":"cesati"}]}`,
			statusCode: 200,
		},
		test{
			method:     "GET",
			url:        "/books/4",
			request:    "",
			response:   `{"id":4,"name":"Understanding the Linux Kernel","isbn":"new_isbn","authors":[{"first_name":"Daniel","last_name":"Bovet"},{"first_name":"Marco","last_name":"cesati"}]}`,
			statusCode: 200,
		},
	}

	runTest(t, testSuite)
}

func TestDeleteBookHandler(t *testing.T) {
	testSuite := []test{
		test{
			method:     "DELETE",
			url:        "/books/4",
			request:    `{"name": "Understanding the Linux Kernel", "isbn":"new_isbn", "authors": [{"first_name": "Daniel", "last_name": "Bovet"}, {"first_name": "Marco", "last_name": "cesati"}]}`,
			response:   `{"id":4,"name":"Understanding the Linux Kernel","isbn":"new_isbn","authors":[{"first_name":"Daniel","last_name":"Bovet"},{"first_name":"Marco","last_name":"cesati"}]}`,
			statusCode: 200,
		},
		test{
			method:     "GET",
			url:        "/books/4",
			request:    "",
			response:   `{"error": "book not found"}`,
			statusCode: 404,
		},
	}

	runTest(t, testSuite)
}

func runTest(t *testing.T, testSuite []test) {
	for _, mytest := range testSuite {

		r, err := http.NewRequest(mytest.method, mytest.url, strings.NewReader(mytest.request))
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		if w.Code != mytest.statusCode {
			t.Error(
				"\nexpected status code", mytest.statusCode,
				"\nfound status code", w.Code,
				"\nfor request", mytest.method, mytest.url, mytest.request,
			)
		}

		if w.Body.String() != mytest.response {
			t.Error(
				"\nexpected response", mytest.response,
				"\nfound response", w.Body.String(),
				"\nfor request", mytest.method, mytest.url, mytest.request,
			)
		}
	}

}
