package model

import "time"

type Book struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	AuthorID    int64     `json:"author_id"`
	Category    string    `json:"category"`
	Path        string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
}
