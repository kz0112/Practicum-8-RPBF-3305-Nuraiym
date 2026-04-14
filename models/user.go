package models

import "gorm.io/gorm"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password" binding:"required"` // ❗ response-та шықпайды
	CreatedAt int64
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	// 🔗 Relations
	Appointments []Appointment `json:"appointments"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
