package handlers

import (
	"net/http"
	"strconv"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// =========================
// ❤️ GET /books/favorites
// =========================
// GetFavorites godoc
// @Summary Get user favorite books
// @Tags favorites
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.FavoriteBook
// @Router /books/favorites [get]
func GetFavorites(c *gin.Context) {
	var favorites []models.FavoriteBook

	// 🔥 user_id токеннен аламыз
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := uint(userIDValue.(float64))

	// pagination (қарапайым)
	limit := 10
	offset := 0

	config.DB.
		Where("user_id = ?", userID).
		Preload("Book").
		Limit(limit).
		Offset(offset).
		Find(&favorites)

	c.JSON(http.StatusOK, favorites)
}

// =========================
// ➕ PUT /books/:bookId/favorites
// =========================
// AddFavorite godoc
// @Summary Add book to favorites
// @Tags favorites
// @Security BearerAuth
// @Param bookId path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{bookId}/favorites [put]
func AddFavorite(c *gin.Context) {
	bookIDParam := c.Param("bookId")
	bookID := parseUint(bookIDParam)

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := uint(userIDValue.(float64))

	// 🔥 duplicate тексеру (өте маңызды)
	var existing models.FavoriteBook
	if err := config.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		First(&existing).Error; err == nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Already in favorites",
		})
		return
	}

	fav := models.FavoriteBook{
		UserID: userID,
		BookID: bookID,
	}

	if err := config.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Added to favorites",
	})
}

// =========================
// ❌ DELETE /books/:bookId/favorites
// =========================
// RemoveFavorite godoc
// @Summary Remove book from favorites
// @Tags favorites
// @Security BearerAuth
// @Param bookId path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{bookId}/favorites [delete]
func RemoveFavorite(c *gin.Context) {
	bookIDParam := c.Param("bookId")
	bookID := parseUint(bookIDParam)

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID := uint(userIDValue.(float64))

	if err := config.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Delete(&models.FavoriteBook{}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Removed from favorites",
	})
}

// =========================
// 🔧 HELPER
// =========================
func parseUint(s string) uint {
	id, _ := strconv.Atoi(s)
	return uint(id)
}
