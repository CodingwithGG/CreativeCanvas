package categoriesServices

import (
	categoriesModels "awesomeProject/categories/models"
	categoriesSerializers "awesomeProject/categories/serializers"
	"awesomeProject/config"

	"github.com/jinzhu/gorm"
	"net/http"
)

func GetCategoriesService() (config.Response, int) {
	db := config.GetDb()
	var categoriesObj *gorm.DB
	var categories []categoriesModels.Categories
	categoriesObj = db.Model(&categoriesModels.Categories{}).Find(&categories)
	if categoriesObj.Error != nil {
		return config.BuildErrorResponse("internal server error", categoriesObj.Error.Error(), nil),
			http.StatusInternalServerError
	}
	var categoriesResponse categoriesSerializers.CategoriesResponse
	var categoriesResponseArray []categoriesSerializers.CategoriesResponse
	for _, categories := range categories {
		categoriesResponse.Category = *categories.Category
		categoriesResponseArray = append(categoriesResponseArray, categoriesResponse)
	}
	response := categoriesSerializers.CategoriesResponseArray{
		Categories: categoriesResponseArray,
	}
	return config.BuildResponse("Categories List", response), http.StatusOK
}
