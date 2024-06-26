package internal

import (
	"context"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
)

type UserRepo interface {
	// Get gets multiple or single user record if id provided
	Get(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	// Store create new user record
	Store(ctx context.Context, user entity.User) (uuid.UUID, error)
}

type BookRepo interface {
	// Get gets multiple or single book record if id provided
	Get(ctx context.Context, filter entity.BookFilter) ([]entity.Book, error)
	// Store create new book record or update if id provided, will return affected id
	Store(ctx context.Context, book entity.Book) (uuid.UUID, error)
}

type OrderRepo interface {
	// Get gets multiple or single order record if id provided
	Get(ctx context.Context, filter entity.OrderFilter) ([]entity.Order, error)
	// Store create new order record or update if id provided, will return affected id
	Store(ctx context.Context, order entity.Order) (uuid.UUID, error)
}
