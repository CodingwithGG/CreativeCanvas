package userHandlers

import (
	"awesomeProject/accounts/serializers"
	"awesomeProject/accounts/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequestObj userSerializers.UserStatusRequest
		if err := c.ShouldBindQuery(&userRequestObj); err != nil {
			log.Println("unable to bind userRequestObj", err.Error())
		}
		response, statusCode := userServices.GetUserStatusService(&userRequestObj)
		c.IndentedJSON(statusCode, response)
	}

}

func PatchUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Status(http.StatusNoContent)
	}

}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequestObj userSerializers.UserStatusRequest
		if err := c.ShouldBindQuery(&userRequestObj); err != nil {
			log.Println("unable to bind userRequestObj", err.Error())
		}
		response, statusCode := userServices.GetUserStatusService(&userRequestObj)
		c.IndentedJSON(statusCode, response)
	}

}
