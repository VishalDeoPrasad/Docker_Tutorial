package handlers

import (
	"golang/middleware" 

	"github.com/gin-gonic/gin"
)

// API function to define and return a Gin engine
func API() *gin.Engine {
	// Create a new Gin engine
	r := gin.New()

	// Use the middleware and recovery globally
	r.Use(middleware.Logger(), gin.Recovery())

	// Define routes
	r.GET("/welcome", welcome)
	r.POST("/whoru", name)

	return r
}

// Handler function for the GET /welcome endpoint
func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome!",
	})
}

// Handler function for the POST /whoru endpoint
func name(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, who are you?",
	})
}
