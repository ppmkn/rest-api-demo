package main

import (
  "github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) {
  r.HandleFunc("/books", getBooks).Methods("GET")
  r.HandleFunc("/books/{id}", getBooks).Methods("GET")
  r.HandleFunc("/books", createBook).Methods("POST")
  r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
}