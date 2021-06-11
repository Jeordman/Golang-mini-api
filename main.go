package main

// allow console printing
// "fmt"
// "encoding/json"
// "log"
// "net/http"
// "math/rand"
// "strconv"
// "github.com/gorilla/mux"
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request)     {}
func createBooks(w http.ResponseWriter, r *http.Request) {}
func updateBook(w http.ResponseWriter, r *http.Request)  {}
func deleteBook(w http.ResponseWriter, r *http.Request)  {}

func main() {
	// Init Router (type inferance)
	r := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "456456", Title: "Book One", Author: &Author{Firstname: "Jeordin", Lastname: "Callister"}})
	books = append(books, Book{ID: "2", Isbn: "465456", Title: "Book Two", Author: &Author{Firstname: "Brenna", Lastname: "Lastname"}})

	r.HandleFunc("api/books", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books", createBooks).Methods("POST")
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":8080", r))
}
