package model

import "time"

type Order struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	TotalPrice int64     `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
