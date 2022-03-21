package services

import (
	"awesomeProject/accounts/models"
	"awesomeProject/accounts/serializers"
	"awesomeProject/config"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetUserStatusService(userRequestObj *serializers.UserStatusRequest) (config.Response, int) {
	db := config.GetDb()
	var userObj *gorm.DB
	var users []models.User
	if userRequestObj.Mobile != "" {
		userObj = db.Model(&models.User{}).Where("mobile = ?", userRequestObj.Mobile).Find(&users)

	} else if userRequestObj.Email != "" {
		userObj = db.Model(&models.User{}).Where("email = ?", userRequestObj.Email).Find(&users)
	} else {
		return config.BuildErrorResponse("give either mobile or email", "error", nil),
			http.StatusBadRequest

	}
	if userObj.Error != nil {
		return config.BuildErrorResponse("internal server error", "internal server error", nil),
			http.StatusBadRequest
	}

	if len(users) == 0 {
		return config.BuildResponse("",
			serializers.UserStatusResponse{OnBoardingStatus: "on_boarding_started"}), http.StatusOK
	}
	var user serializers.UserStatusResponse
	user.OnBoardingStatus = users[0].OnBoardingStatus
	return config.BuildResponse("", user), http.StatusOK

}
func CreateUserService(userInfo *serializers.CreateUser) {
	db := config.GetDb()
	user, password := serializers.CreateUserModels(userInfo)
	db.Create(user)
	db.Create(password)

}
