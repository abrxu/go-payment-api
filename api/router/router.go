package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Initialize() {
	router := gin.Default()

	UseCors(router)
	InitializeRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
