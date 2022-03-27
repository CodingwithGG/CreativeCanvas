package categoriesHandlers

import (
	"awesomeProject/categories/services"
	"github.com/gin-gonic/gin"
)

func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		response, statusCode := categoriesServices.GetCategoriesService()
		c.IndentedJSON(statusCode, response)
	}
}
