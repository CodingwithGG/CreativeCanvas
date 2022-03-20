package routes

import (
	"awesomeProject/accounts/handlers"
	"github.com/gin-gonic/gin"
)

func (r routes) user(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	user.GET("/status", handlers.GetUserStatus())
	user.GET("/", handlers.GetUser())
	user.POST("/", handlers.CreateUser())
	user.PATCH("/", handlers.PatchUser())
}
