package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) registerRoutes() {
	a.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// authentication endpoint
	a.gin.POST("/login", a.authDelivery.Login)

	// user endpoint
	user := a.gin.Group("user")
	user.POST("/register", a.userDelivery.Register)
	userAuth := user.Use(a.middleware.Authenticated())
	userAuth.GET("/me", a.userDelivery.Me)

	// book endpoint
	book := a.gin.Group("book")
	book.GET("/list", a.bookDelivery.GetList)
	book.GET("/:id", a.bookDelivery.Get)
}
