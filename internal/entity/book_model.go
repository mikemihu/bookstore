package entity

import "github.com/google/uuid"

type Book struct {
	BaseModel
	ID       uuid.UUID
	ISBN     string
	Author   string
	Title    string
	Subtitle string
}

func (b *Book) ToResponse() BookResponse {
	return BookResponse{
		ID:       b.ID,
		ISBN:     b.ISBN,
		Author:   b.Author,
		Title:    b.Title,
		Subtitle: b.Subtitle,
	}
}
