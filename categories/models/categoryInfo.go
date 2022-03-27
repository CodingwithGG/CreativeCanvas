package categoriesModels

import (
	"container/list"
	"github.com/jinzhu/gorm"
)

type CategoriesInfo struct {
	gorm.Model
	//Category_id forein key
	//User_id	forein key
	ExpectedPriceFrom float64 `json:"expected_price_from"`
	ExpectedPriceTo   float64 `json:"expected_price_to"`
	//Calender extension
	//Specific category info
	Photos list.Element `json:"photos"`
	//Upto 4 photos to build profile
	About  string `json:"about"`
	Doc    string `json:"doc"`
	Videos string `json:"videos"`
	//Two videos
	Experience string `json:"experience"`
	//certificates/awards
}
