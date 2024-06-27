package auth_jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthJWTImpl struct {
	secret    []byte
	expiresIn time.Duration
}

type AuthJWT interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(tokenString string) (AuthClaims, error)
}

func NewAuthJWT(secret []byte) AuthJWT {
	return &AuthJWTImpl{
		secret:    secret,
		expiresIn: time.Hour * 72, // expired in 3 days
	}
}

// GenerateToken return jwt token of user id
func (a *AuthJWTImpl) GenerateToken(userID uuid.UUID) (string, error) {
	claims := AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.expiresIn)),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (a *AuthJWTImpl) ParseToken(tokenString string) (AuthClaims, error) {
	var claims AuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return a.secret, nil
	})
	if err != nil {
		return AuthClaims{}, err
	}

	if !token.Valid {
		return AuthClaims{}, errors.New("token is not valid")
	}

	return claims, nil
}
