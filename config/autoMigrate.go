package config

import (
	userModels "awesomeProject/accounts/models"
	authModels "awesomeProject/authentication/models"
	categoriesModels "awesomeProject/categories/models"
)

func AutoMigrate() {
	db.AutoMigrate(&userModels.User{})
	db.AutoMigrate(&authModels.Password{})
	db.AutoMigrate(&authModels.AuthToken{})
	db.AutoMigrate(&categoriesModels.Categories{})
}
