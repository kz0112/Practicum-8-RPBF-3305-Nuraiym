package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"Private-medical-clinic.backend/models"
	"Private-medical-clinic.backend/storage"
)

// BooksHandler кітаптар тізімін басқарады
// @Summary Кітаптар тізімін алу
// @Description Барлық кітаптарды фильтрлермен және пагинациямен қайтарады
// @Tags books
// @Accept json
// @Produce json
// @Param author query string false "Автор аты бойынша фильтр"
// @Param category query string false "Категория бойынша фильтр"
// @Param title query string false "Тақырып бойынша фильтр"
// @Param page query int false "Бет нөмірі" default(1)
// @Param limit query int false "Беттегі элементтер саны" default(5)
// @Success 200 {array} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [get]
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := storage.Books

	// Filters
	author := strings.ToLower(r.URL.Query().Get("author"))
	category := strings.ToLower(r.URL.Query().Get("category"))
	title := strings.ToLower(r.URL.Query().Get("title"))

	var filtered []models.Book

	for _, book := range books {
		if author != "" && !strings.Contains(strings.ToLower(book.Author), author) {
			continue
		}
		if category != "" && !strings.Contains(strings.ToLower(book.Category), category) {
			continue
		}
		if title != "" && !strings.Contains(strings.ToLower(book.Title), title) {
			continue
		}
		filtered = append(filtered, book)
	}

	// Pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	json.NewEncoder(w).Encode(filtered[start:end])
}

// CreateBook жаңа кітап жасау
// @Summary Жаңа кітап қосу
// @Description Жаңа кітапты дерекқорға қосады
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Кітап мәліметтері"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if newBook.Title == "" || newBook.Author == "" {
		http.Error(w, "Title and Author required", http.StatusBadRequest)
		return
	}

	newBook.ID = len(storage.Books) + 1
	storage.Books = append(storage.Books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// BookByIDHandler бір кітаппен жұмыс істейді
// @Summary Кітапты ID бойынша алу
// @Description ID бойынша кітапты қайтарады
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Кітап ID-і"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func BookByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Кітапты іздеу
	for _, book := range storage.Books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// UpdateBook кітапты жаңарту
// @Summary Кітапты жаңарту
// @Description ID бойынша кітапты толық жаңартады
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Кітап ID-і"
// @Param book body models.Book true "Жаңартылған кітап мәліметтері"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, book := range storage.Books {
		if book.ID == id {
			var updatedBook models.Book
			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			updatedBook.ID = id
			storage.Books[i] = updatedBook

			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// DeleteBook кітапты жою
// @Summary Кітапты жою
// @Description ID бойынша кітапты дерекқордан жояды
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Кітап ID-і"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, book := range storage.Books {
		if book.ID == id {
			storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}