package entity

import (
	"github.com/google/uuid"
)

type BookFilter struct {
	Book
	Search string
	IDs    uuid.UUIDs
}

type BookResponse struct {
	ID       uuid.UUID `json:"id"`
	ISBN     string    `json:"isbn"`
	Author   string    `json:"author"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Price    float64   `json:"price"`
}

type BookGetListRequest struct {
	Search string   `form:"search"`
	IDs    []string `form:"ids"`
}
