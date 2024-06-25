package delivery

import (
	"net/http"

	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/entity"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	cfg    *config.Cfg
	userUC internal.UserUC
}

func NewUserDelivery(
	cfg *config.Cfg,
	userUC internal.UserUC,
) internal.UserDelivery {
	return &UserDelivery{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (u *UserDelivery) Register(c *gin.Context) {
	var req entity.UserRegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email and password is required"})
		return
	}

	err = u.userUC.Register(c, req)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserDelivery) Me(c *gin.Context) {
	user, err := u.userUC.Me(c)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
