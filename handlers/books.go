package handlers

import (
	"net/http"
	"strconv"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// =========================
// 📚 GET /books (pagination + filter)
// =========================
// GetBooks godoc
// @Summary Get all books
// @Description Get books with pagination and filters
// @Tags books
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param category query int false "Category ID"
// @Param author query int false "Author ID"
// @Param title query string false "Book title"
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book

	// Query params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	category := c.Query("category")
	author := c.Query("author")
	title := c.Query("title")

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	query := config.DB

	// Filters
	if category != "" {
		query = query.Where("category_id = ?", category)
	}
	if author != "" {
		query = query.Where("author_id = ?", author)
	}
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// 🔥 preload (JOIN)
	query = query.Preload("Author").Preload("Category")

	if err := query.
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Find(&books).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch books",
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

// =========================
// ➕ POST /books
// =========================
// CreateBook godoc
// @Summary Create book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.CreateBookInput true "Book"
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validation
	if input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if input.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	book := models.Book{
		Title:      input.Title,
		Price:      input.Price,
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create book",
		})
		return
	}

	// 🔥 МЫНА ЖЕРДІ ҚОСАСЫҢ
	if err := config.DB.Preload("Author").Preload("Category").First(&book, book.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load relations",
		})
		return
	}

	c.JSON(http.StatusCreated, book)

}

// =========================
// 🔍 GET /books/:id
// =========================
// GetBookByID godoc
// @Summary Get book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.Preload("Author").Preload("Category").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// =========================
// ✏️ PUT /books/:id
// =========================
// UpdateBook godoc
// @Summary Update book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.UpdateBookInput true "Updated book"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// partial update
	if input.Title != "" {
		book.Title = input.Title
	}
	if input.Price != 0 {
		book.Price = input.Price
	}
	if input.AuthorID != 0 {
		book.AuthorID = input.AuthorID
	}
	if input.CategoryID != 0 {
		book.CategoryID = input.CategoryID
	}

	config.DB.Save(&book)
	if err := config.DB.Preload("Author").Preload("Category").First(&book, book.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load relations",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

// =========================
// ❌ DELETE /books/:id
// =========================
// DeleteBook godoc
// @Summary Delete book
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
