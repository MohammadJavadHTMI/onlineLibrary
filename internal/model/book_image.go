package model

import "time"

type BookImage struct {
	ID        int64     `json:"id"`
	BookID    int64     `json:"book_id"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}
