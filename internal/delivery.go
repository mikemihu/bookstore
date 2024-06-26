package internal

import "github.com/gin-gonic/gin"

type AuthDelivery interface {
	// Login returns user token
	Login(c *gin.Context)
}

type UserDelivery interface {
	// Register creates new user
	Register(c *gin.Context)
	// Me gets authenticated user info
	Me(c *gin.Context)
}

type BookDelivery interface {
	// GetList returns all books
	GetList(c *gin.Context)
	// Get returns single book
	Get(c *gin.Context)
}

type OrderDelivery interface {
	// 	GetList returns user's orders history
	GetList(c *gin.Context)
	// Get returns user's order detail
	Get(c *gin.Context)
	// Create creates new order
	Create(c *gin.Context)
}
