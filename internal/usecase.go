package internal

import (
	"context"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
)

type UserUC interface {
	// Register creates new user
	Register(ctx context.Context, req entity.UserRegisterRequest) error
	// Login returns user's auth token
	Login(ctx context.Context, req entity.AuthLoginRequest) (string, error)
	// Get returns single user record
	Get(ctx context.Context, id uuid.UUID) (entity.UserResponse, error)
	// Me returns current user's record
	Me(ctx context.Context) (entity.UserResponse, error)
}

type BookUC interface {
	// GetList returns all book
	GetList(ctx context.Context, filter entity.BookFilter) ([]entity.BookResponse, error)
	// Get returns single book record
	Get(ctx context.Context, filter entity.BookFilter) (entity.BookResponse, error)
}
