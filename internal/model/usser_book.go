package model

import "time"

type UserBook struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	BookID      int64     `json:"book_id"`
	OrderID     int64     `json:"order_id"`
	PurchasedAt time.Time `json:"purchased_at"`
}
