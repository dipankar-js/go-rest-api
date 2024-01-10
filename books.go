// Handlers for the router yet to be implemented
package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {
	err := json.NewEncoder(w).Encode((listBooks()))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {
	id := chi.URLParam(r, "id")
	book := getBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}