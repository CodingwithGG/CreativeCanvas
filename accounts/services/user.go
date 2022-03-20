package services

import (
	"awesomeProject/accounts/models"
	"awesomeProject/accounts/serializers"
	"awesomeProject/config"
	"github.com/jinzhu/gorm"
)

func GetUserStatusService(userRequestObj *serializers.UserStatusRequest) config.Response {
	db := config.GetDb()
	var userObj *gorm.DB
	var user []serializers.UserStatusResponse
	if userRequestObj.Email == "" {
		userObj = db.Model(&models.User{}).Find("mobile = ?", userRequestObj.Mobile)

	} else if userRequestObj.Mobile == "" {
		userObj = db.Model(&models.User{}).Find("email = ?", userRequestObj.Mobile)
	} else {
		return config.BuildErrorResponse("give either mobile or email", "error", nil)

	}
	if userObj.Error == nil {
		return config.BuildErrorResponse("internal server error", "internal server error", nil)
	}

	if len(user) == 0 {
		return config.BuildResponse("",
			serializers.UserStatusResponse{OnBoardingStatus: "on_boarding_started"})
	}
	return config.BuildResponse("", user)

}
