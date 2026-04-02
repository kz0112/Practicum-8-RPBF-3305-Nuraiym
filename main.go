package main

import (
	"log"

	"Private-medical-clinic.backend/config"
	"Private-medical-clinic.backend/handlers"
	"Private-medical-clinic.backend/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Private-medical-clinic.backend/docs"
)

// @title Private Medical Clinic API
// @version 1.0
// @description Бұл жеке медициналық клиниканың API-сы
// @host localhost:8080
// @BasePath /

func main() {

	// 🔌 Database қосу
	config.ConnectDB()

	// 🔄 Migration
	config.DB.AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.Category{},
	)

	// 🚀 Gin router
	r := gin.Default()

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// =========================
	// 📚 BOOK ROUTES
	// =========================
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// =========================
	// 👤 AUTHOR ROUTES
	// =========================
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)
	r.GET("/authors/:id", handlers.GetAuthorByID)

	// =========================
	// 🏷 CATEGORY ROUTES
	// =========================
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	// =========================
	// ▶️ SERVER START
	// =========================
	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
