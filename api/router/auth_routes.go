package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/nacl/auth"
)

func RegisterAuthRoutes(routes *gin.RouterGroup) {
	authRoutes := routes.Group("/auth")
	{
		authRoutes.POST("/register", auth.RegisterUserHandler)
		authRoutes.POST("/login", auth.LoginUserHandler)
	}
}
