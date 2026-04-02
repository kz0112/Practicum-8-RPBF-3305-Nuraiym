package handlers

import (
	"net/http"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// GetAuthors godoc
// @Summary Get all authors
// @Tags authors
// @Produce json
// @Success 200 {array} models.Author
// @Router /authors [get]
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	config.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

// CreateAuthor godoc
// @Summary Create author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body models.Author true "Author"
// @Success 201 {object} models.Author
// @Router /authors [post]
func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&author)

	c.JSON(201, author)
}

// GetAuthorByID godoc
// @Summary Get author by ID
// @Tags authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [get]
func GetAuthorByID(c *gin.Context) {
	var author models.Author
	id := c.Param("id")

	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Author not found",
		})
		return
	}

	c.JSON(http.StatusOK, author)
}
