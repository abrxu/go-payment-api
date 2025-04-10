package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeRoutes(router *gin.Engine) {
	routes := router.Group("/api/" + os.Getenv("API_VERSION"))

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "pong"})
	})

	RegisterAuthRoutes(routes)
}
