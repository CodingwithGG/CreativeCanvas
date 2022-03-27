package routes

import (
	"awesomeProject/authentication/handlers"
	"github.com/gin-gonic/gin"
)

func (r routes) auth(rg *gin.RouterGroup) {

	user := rg.Group("/auth")

	user.POST("/", authHandlers.GetTokenHandler())
	user.GET("/logout", authHandlers.Logout())
}
