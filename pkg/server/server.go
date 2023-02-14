package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(dbHandler *gorm.DB) {

	// Creates a router without any middleware by default
	r := gin.New()

	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Handler
	userHandler := bindUser(dbHandler)

	// Group of users
	user := r.Group("/user")
	{
		user.POST("/login")
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

