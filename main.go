package main

import (
	"log"
	"net/http"
	c "rest-apis/gorilla-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/books", c.GetBooks).Methods(http.MethodGet)
	api.HandleFunc("/books/{id}", c.GetBook).Methods(http.MethodGet)
	api.HandleFunc("/books", c.CreateBook).Methods(http.MethodPost)
	api.HandleFunc("/books/{id}", c.UpdateBook).Methods(http.MethodPut)
	api.HandleFunc("/books/{id}", c.DeleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", r))

}
