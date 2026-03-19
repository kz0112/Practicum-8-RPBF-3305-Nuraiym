package main

import (
	"log"
	"net/http"

	"Private-medical-clinic.backend/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.BooksHandler)
	http.HandleFunc("/books/", handlers.BookByIDHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
