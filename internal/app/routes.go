package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) registerRoutes() {
	a.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
