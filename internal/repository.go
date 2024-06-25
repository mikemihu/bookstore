package internal

import (
	"context"
	"gotu-bookstore/internal/entity"

	"github.com/google/uuid"
)

type UserRepo interface {
	// Get gets multiple or single row if id provided
	Get(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	// Store create new user record
	Store(ctx context.Context, user entity.User) (uuid.UUID, error)
}
