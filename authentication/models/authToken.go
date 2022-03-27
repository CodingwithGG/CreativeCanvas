package authModels

import (
	"awesomeProject/accounts/models"
	"github.com/jinzhu/gorm"
)

type AuthToken struct {
	gorm.Model
	UserId userModels.User `gorm:"references:Code;json:user_id"`
	Token  *string         `gorm:"type:varchar(100);unique_index;json:token"`
}
