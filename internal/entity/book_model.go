package entity

type Book struct {
	BaseModel
	ISBN     string
	Author   string
	Title    string
	Subtitle string
	Price    float64
}

func (b *Book) ToResponse() BookResponse {
	return BookResponse{
		ID:       b.ID,
		ISBN:     b.ISBN,
		Author:   b.Author,
		Title:    b.Title,
		Subtitle: b.Subtitle,
		Price:    b.Price,
	}
}
