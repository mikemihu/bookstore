package auth_jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}
