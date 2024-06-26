package entity

import (
	"time"

	"github.com/google/uuid"
)

type OrderFilter struct {
	Order
	PreloadDetail bool
}

type OrderResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     uuid.UUID `json:"user_id"`
	TotalQty   int       `json:"total_qty"`
	TotalPrice float64   `json:"total_price"`

	Items []OrderItemResponse `json:"items,omitempty"`
}

type OrderItemResponse struct {
	ID      uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id"`
	BookID  uuid.UUID `json:"book_id"`
	Qty     int       `json:"qty"`
	Price   float64   `json:"price"`

	Book *BookResponse `json:"book,omitempty"`
}

type OrderCreateRequest struct {
	Items []struct {
		BookID uuid.UUID `json:"book_id"`
		Qty    int       `json:"qty"`
	} `json:"items"`
}
