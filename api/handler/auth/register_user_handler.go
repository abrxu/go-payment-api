package handler

import (
	"abrxu.com/gin/api/handler"
	request "abrxu.com/gin/api/request/auth"
	"abrxu.com/gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHandler(c *gin.Context) {
	userRequest := request.RegisterUserRequest{}

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": handler.ValidateError(err),
			})
		return
	}

	hashPassword, err := utils.HashPassword(userRequest.Password)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": handler.ValidateError(err),
		})
		return
	}

	_, err = user.Create
}
