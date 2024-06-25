package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthJWTImpl struct {
	secret []byte
}

type AuthJWT interface {
	GenerateToken(secret string, userID uuid.UUID) (string, error)
	ParseToken(secret string, tokenString string) (AuthClaims, error)
}

func NewAuthJWT(secret []byte) AuthJWT {
	return &AuthJWTImpl{
		secret: secret,
	}
}

type AuthClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

// GenerateToken return jwt token of user id
func (a *AuthJWTImpl) GenerateToken(secret string, userID uuid.UUID) (string, error) {
	claims := AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (a *AuthJWTImpl) ParseToken(secret string, tokenString string) (AuthClaims, error) {
	var claims AuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return AuthClaims{}, err
	}

	if !token.Valid {
		return AuthClaims{}, err
	}

	return claims, nil
}
