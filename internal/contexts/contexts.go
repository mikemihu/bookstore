package contexts

import (
	"context"
	"gotu-bookstore/internal/entity"
)

const (
	CtxKeyUser = "user"
)

// GetUser returns current user from context
func GetUser(ctx context.Context) entity.User {
	user, ok := ctx.Value(CtxKeyUser).(entity.User)
	if !ok {
		return entity.User{}
	}
	return user
}
