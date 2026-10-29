package model

type OrderItem struct {
	ID      int64 `json:"id"`
	OrderID int64 `json:"order_id"`
	BookID  int64 `json:"book_id"`
	Price   int64 `json:"price"`
}
