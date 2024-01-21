package models

import "github.com/google/uuid"

// News represents the News entity
type News struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Title string    `db:"title" json:"title" binding:"required"`
	Body  string    `db:"body" json:"body" binding:"required"`
}
