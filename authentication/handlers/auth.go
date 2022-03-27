package authHandlers

import (
	"awesomeProject/authentication/serializers"
	"awesomeProject/authentication/services"
	"github.com/gin-gonic/gin"
	"log"
)

func GetTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials authSerializers.Credentials
		if err := c.BindJSON(&credentials); err != nil {
			log.Println("unable to bind credentials", err.Error())
		}
		response, statusCode := authServices.GetTokenService(&credentials, c)
		c.IndentedJSON(statusCode, response)
	}

}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, statusCode := authServices.LogoutService(c)
		c.IndentedJSON(statusCode, response)
	}

}
