package categoriesModels

import (
	"github.com/jinzhu/gorm"
)

type Categories struct {
	gorm.Model

	Category *string `gorm:"type:varchar(100);unique_index;json:category"`
}
