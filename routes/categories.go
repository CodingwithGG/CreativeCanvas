package routes

import (
	"awesomeProject/authentication/middlewares"
	categoriesHandlers "awesomeProject/categories/handlers"
	"github.com/gin-gonic/gin"
)

func (r routes) categories(rg *gin.RouterGroup) {

	user := rg.Group("/categories")

	user.GET("/", middlewares.AuthorizeJWT(), categoriesHandlers.GetCategories())

}
