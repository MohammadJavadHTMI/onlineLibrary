package model

import "time"

type Payment struct {
	ID        int64      `json:"id"`
	OrderID   int64      `json:"order_id"`
	Amount    int64      `json:"amount"`
	Authority *string    `json:"authority"`
	RefID     *string    `json:"ref_id"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	PaidAt    *time.Time `json:"paid_at"`
}
