package entity

import "github.com/google/uuid"

type Order struct {
	BaseModel
	UserID     uuid.UUID
	TotalQty   int
	TotalPrice float64

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	BaseModel
	OrderID uuid.UUID
	BookID  uuid.UUID
	Qty     int
	Price   float64

	Order Order `gorm:"foreignKey:OrderID"`
	Book  Book  `gorm:"foreignKey:BookID"`
}

func (o *Order) ToResponse() OrderResponse {
	resp := OrderResponse{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		UserID:     o.UserID,
		TotalQty:   o.TotalQty,
		TotalPrice: o.TotalPrice,
		Items:      make([]OrderItemResponse, len(o.Items)),
	}
	for i := range o.Items {
		resp.Items[i] = o.Items[i].ToResponse()
	}
	return resp
}

func (o *OrderItem) ToResponse() OrderItemResponse {
	resp := OrderItemResponse{
		ID:      o.ID,
		OrderID: o.OrderID,
		BookID:  o.BookID,
		Qty:     o.Qty,
		Price:   o.Price,
	}
	if o.Book.ID != uuid.Nil {
		bookResp := o.Book.ToResponse()
		resp.Book = &bookResp
	}
	return resp
}
