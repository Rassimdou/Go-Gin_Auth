package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.authRoutes(router)
	router.userRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API v1 is working!",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API v2 is working!",
		})
	})
	router.Run(":" + port)
}
