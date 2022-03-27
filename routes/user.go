package routes

import (
	"awesomeProject/accounts/handlers"
	"awesomeProject/authentication/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) user(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	user.GET("/status", userHandlers.GetUserStatus())
	user.GET("/info", middlewares.AuthorizeJWT(), userHandlers.GetUserInfo())
	user.PATCH("/", middlewares.AuthorizeJWT(), userHandlers.PatchUser())
}
