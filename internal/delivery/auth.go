package delivery

import (
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthDelivery struct {
	cfg    *config.Cfg
	userUC internal.UserUC
}

func NewAuthDelivery(
	cfg *config.Cfg,
	userUC internal.UserUC,
) internal.AuthDelivery {
	return &AuthDelivery{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (u *AuthDelivery) Login(c *gin.Context) {
	var req entity.AuthLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email and password is required"})
		return
	}

	token, err := u.userUC.Login(c, req)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	resp := entity.AuthLoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, resp)
}
