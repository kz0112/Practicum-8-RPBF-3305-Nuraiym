package handlers

import (
	"net/http"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary Get all categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category"
// @Success 201 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&category)

	c.JSON(201, category)
}
