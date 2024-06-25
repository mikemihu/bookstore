package entity

import (
	"github.com/google/uuid"
)

type BookResponse struct {
	ID       uuid.UUID `json:"id"`
	ISBN     string    `json:"isbn"`
	Author   string    `json:"author"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
}

type BookFilter struct {
	Book
	Search string
	IDs    uuid.UUIDs
}

type BookGetListRequest struct {
	Search string `form:"search"`
}
