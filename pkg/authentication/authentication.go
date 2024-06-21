package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

// GenerateToken return jwt token of user id
func GenerateToken(secret string, userID uuid.UUID) (string, error) {
	authClaims := AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
