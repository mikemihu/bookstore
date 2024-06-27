package auth_jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthJWT(t *testing.T) {
	dummyUserID := uuid.New()

	type fields struct {
		secret []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AuthClaims
		wantErr error
	}{
		{
			name: "ok",
			fields: fields{
				secret: []byte("test-secret"),
			},
			args: args{
				tokenString: func() string {
					a := NewAuthJWT([]byte("test-secret"))
					tokenString, err := a.GenerateToken(dummyUserID)
					if err != nil {
						t.Fatal(err)
					}
					return tokenString
				}(),
			},
			want: AuthClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
				},
				UserID: dummyUserID,
			},
			wantErr: nil,
		},
		{
			name: "error - signature invalid",
			fields: fields{
				secret: []byte("test-secret"),
			},
			args: args{
				tokenString: func() string {
					a := NewAuthJWT([]byte("different-test-secret"))
					tokenString, err := a.GenerateToken(dummyUserID)
					if err != nil {
						t.Fatal(err)
					}
					return tokenString
				}(),
			},
			want:    AuthClaims{},
			wantErr: jwt.ErrSignatureInvalid,
		},
		{
			name: "error - token expired",
			fields: fields{
				secret: []byte("test-secret"),
			},
			args: args{
				tokenString: func() string {
					// Generate a token that's already expired
					claims := AuthClaims{
						UserID: dummyUserID,
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * -72)), // 72 hours in the past
						},
					}
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
					tokenString, err := token.SignedString([]byte("test-secret"))
					if err != nil {
						t.Fatal(err)
					}
					return tokenString
				}(),
			},
			want:    AuthClaims{},
			wantErr: jwt.ErrTokenExpired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAuthJWT(tt.fields.secret)
			got, err := a.ParseToken(tt.args.tokenString)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
