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
