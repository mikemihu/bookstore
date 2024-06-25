package middleware

import (
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"
	authPkg "gotu-bookstore/pkg/authentication"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware struct {
	cfg     *config.Cfg
	logger  *zap.Logger
	authJWT authPkg.AuthJWT
	userUC  internal.UserUC
}

func NewMiddleware(
	cfg *config.Cfg,
	logger *zap.Logger,
	authJWT authPkg.AuthJWT,
	userUC internal.UserUC,
) *Middleware {
	return &Middleware{
		cfg:     cfg,
		logger:  logger,
		authJWT: authJWT,
		userUC:  userUC,
	}
}

func (m *Middleware) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// remove "Bearer " prefix if exists
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := m.authJWT.ParseToken(m.cfg.Auth.JwtSecret, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// get user from db
		userResp, err := m.userUC.Get(c.Request.Context(), claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			c.Abort()
			return
		}
		c.Set(constant.CtxKeyUser, userResp)

		c.Next()
	}
}
