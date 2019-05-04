package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book name
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author name
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Enterted to getBooks method")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Enterted to getBook method")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Enterted to createBook method")
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Enterted to updateBook method")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Enterted to deleteBook method")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to rest golang api")
}

// func init(books []Book) []Book {
// 	books = append(books, Book{ID: "1", ISBN: "UYH567", Title: "Wings of Fier", Author: &Author{FirstName: "APJ", LastName: "AbdulKalam"}})
// 	books = append(books, Book{ID: "2", ISBN: "WERF67", Title: "First Man", Author: &Author{FirstName: "Gorge", LastName: "De"}})
// 	return books
// }

func main() {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api").Subrouter()

	books = append(books, Book{ID: "1", ISBN: "UYH567", Title: "Wings of Fier", Author: &Author{FirstName: "APJ", LastName: "AbdulKalam"}})
	books = append(books, Book{ID: "2", ISBN: "WERF67", Title: "First Man", Author: &Author{FirstName: "Gorge", LastName: "De"}})

	sub.HandleFunc("/books", getBooks).Methods("GET")
	sub.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	sub.HandleFunc("/book/{id}", getBook).Methods("GET")
	sub.HandleFunc("/book", createBook).Methods("POST")
	sub.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	sub.HandleFunc("/index", api).Methods("GET")

	// To start server
	log.Println("Server started... port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
