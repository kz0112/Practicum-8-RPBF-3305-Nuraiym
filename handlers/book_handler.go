package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"Private-medical-clinic.backend/models"
	"Private-medical-clinic.backend/storage"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(storage.Books)

	case http.MethodPost:
		var newBook models.Book
		json.NewDecoder(r.Body).Decode(&newBook)

		newBook.ID = len(storage.Books) + 1
		storage.Books = append(storage.Books, newBook)

		json.NewEncoder(w).Encode(newBook)
	}
}

func BookByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i, book := range storage.Books {
		if book.ID == id {
			switch r.Method {
			case http.MethodGet:
				json.NewEncoder(w).Encode(book)

			case http.MethodPut:
				var updatedBook models.Book
				json.NewDecoder(r.Body).Decode(&updatedBook)

				updatedBook.ID = id
				storage.Books[i] = updatedBook

				json.NewEncoder(w).Encode(updatedBook)

			case http.MethodDelete:
				storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
				w.Write([]byte("Deleted"))
			}
			return
		}
	}

	http.NotFound(w, r)
}
