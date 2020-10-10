package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"rest-apis/gorilla-rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

var books []models.Book

func init() {
	books = append(books,
		models.Book{
			ID:    "1",
			Isbn:  "123456",
			Title: "Book 1",
			Author: &models.Author{
				FirstName: "John",
				LastName:  "Doe",
			}},
		models.Book{
			ID:    "2",
			Isbn:  "223456",
			Title: "Book 2",
			Author: &models.Author{
				FirstName: "Lucy",
				LastName:  "Lee",
			}})
}

// GetBooks return all book
// GET /api/books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook return a book by id
// GET /api/books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for _, book := range books {
		if book.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// CreateBook create a new book
// POST /api/books
// Reqest: Book struct
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newBook.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBook)
}

// UpdateBook update a book with id
// PUT /api/books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook models.Book

	// check incoming request
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if book exists
	id := mux.Vars(r)["id"]
	var idx int
	var book models.Book
	for i, b := range books {
		if b.ID == id {
			idx = i
			book = b
			break
		}
	}
	book.Title = updatedBook.Title
	book.Isbn = updatedBook.Isbn
	book.Author = updatedBook.Author
	books[idx] = book
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook delete a book by id
// DELETE /api/books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
