package models

type Book struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Title      string  `json:"title" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	AuthorID   uint    `json:"author_id"`
	CategoryID uint    `json:"category_id"`

	Author   Author   `json:"author" gorm:"foreignKey:AuthorID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type CreateBookInput struct {
	Title      string  `json:"title" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	AuthorID   uint    `json:"author_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
}

type UpdateBookInput struct {
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	AuthorID   uint    `json:"author_id"`
	CategoryID uint    `json:"category_id"`
}
