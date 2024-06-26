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

	// order endpoint
	order := a.gin.Group("order")
	order.Use(a.middleware.Authenticated())
	order.GET("/list", a.orderDelivery.GetList)
	order.GET("/:id", a.orderDelivery.Get)
	order.POST("/create", a.orderDelivery.Create)
}
