package delivery

import (
	"errors"
	"net/http"

	"gotu-bookstore/internal/constant"

	"github.com/gin-gonic/gin"
)

// handlerErrorResponse prevent sensitive error message being exposed to public
func handlerErrorResponse(c *gin.Context, err error) {
	if errors.Is(err, constant.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	if errors.Is(err, constant.ErrNoAccess) {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	if errors.Is(err, constant.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if errors.Is(err, constant.ErrUnimplemented) {
		c.JSON(http.StatusNotImplemented, gin.H{"message": err.Error()})
		return
	}
	if errors.Is(err, constant.ErrUserNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
}
