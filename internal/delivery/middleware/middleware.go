package middleware

import (
	"errors"
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"
	"gotu-bookstore/internal/contexts"
	"gotu-bookstore/internal/entity"
	authPkg "gotu-bookstore/pkg/authentication"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg      *config.Cfg
	logger   *zap.Logger
	authJWT  authPkg.AuthJWT
	userRepo internal.UserRepo
}

func NewMiddleware(
	cfg *config.Cfg,
	logger *zap.Logger,
	authJWT authPkg.AuthJWT,
	userRepo internal.UserRepo,
) *Middleware {
	return &Middleware{
		cfg:      cfg,
		logger:   logger,
		authJWT:  authJWT,
		userRepo: userRepo,
	}
}

func (m *Middleware) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		// remove "Bearer " prefix if exists
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := m.authJWT.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		// get user from db
		filter := entity.UserFilter{
			User: entity.User{
				BaseModel: entity.BaseModel{ID: claims.UserID},
			},
		}
		users, err := m.userRepo.Get(c, filter)
		if err != nil {
			if !errors.Is(err, constant.ErrNotFound) {
				m.logger.Error("failed userRepo.Get", zap.Error(err),
					zap.String("user_id", claims.UserID.String()))
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
			c.Abort()
			return
		}

		// put user into context
		c.Set(contexts.CtxKeyUser, users[0])

		c.Next()
	}
}
