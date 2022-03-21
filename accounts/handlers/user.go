package handlers

import (
	"awesomeProject/accounts/serializers"
	"awesomeProject/accounts/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequestObj serializers.UserStatusRequest
		if err := c.ShouldBindQuery(&userRequestObj); err != nil {
			log.Println("unable to bind userRequestObj", err.Error())
		}
		response, statusCode := services.GetUserStatusService(&userRequestObj)
		c.IndentedJSON(statusCode, response)
	}

}
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInfo serializers.CreateUser
		if err := c.BindJSON(&userInfo); err != nil {
			log.Println("unable to bind userRequestObj", err.Error())
		}
		services.CreateUserService(&userInfo)
		c.Status(http.StatusNoContent)
	}

}
func PatchUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Status(http.StatusNoContent)
	}

}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, nil)
	}

}
