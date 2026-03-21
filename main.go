package main

import (
	"log"
	"net/http"

	"Private-medical-clinic.backend/handlers"

	_ "Private-medical-clinic.backend/docs"
	"github.com/swaggo/http-swagger"
)

// @title Private Medical Clinic API
// @version 1.0
// @description Бұл жеке медициналық клиниканың API-сы
// @host localhost:8080
// @BasePath /

func main() {
	// Books эндпоинттері
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.BooksHandler(w, r)
		case http.MethodPost:
			handlers.CreateBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Бір кітаппен жұмыс істейтін эндпоинттер
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.BookByIDHandler(w, r)
		case http.MethodPut:
			handlers.UpdateBook(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Swagger UI
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	log.Println("Server started at :8080")
	log.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
}