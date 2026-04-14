package models

import "time"

type FavoriteBook struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	BookID    uint      `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`

	Book Book `json:"book"`
}
