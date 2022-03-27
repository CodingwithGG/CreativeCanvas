package userServices

import (
	"awesomeProject/accounts/enums"
	"awesomeProject/accounts/models"
	"awesomeProject/accounts/serializers"
	"awesomeProject/authentication/serializers"
	"awesomeProject/config"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetUserStatusService(userRequestObj *userSerializers.UserStatusRequest) (config.Response, int) {
	db := config.GetDb()
	var userObj *gorm.DB
	var users []userModels.User
	if userRequestObj.Mobile != "" {
		userObj = db.Model(&userModels.User{}).Where("mobile = ?", userRequestObj.Mobile).Find(&users)

	} else if userRequestObj.Email != "" {
		userObj = db.Model(&userModels.User{}).Where("email = ?", userRequestObj.Email).Find(&users)
	} else {
		return config.BuildErrorResponse("give either mobile or email", "", nil),
			http.StatusBadRequest

	}
	if userObj.Error != nil {
		return config.BuildErrorResponse("internal server error", userObj.Error.Error(), nil),
			http.StatusInternalServerError
	}

	if len(users) == 0 {
		return config.BuildResponse("",
			userSerializers.UserStatusResponse{OnBoardingStatus: userEnums.OnBoardingStarted}), http.StatusOK
	}
	var user userSerializers.UserStatusResponse
	user.OnBoardingStatus = users[0].OnBoardingStatus
	return config.BuildResponse("", user), http.StatusOK

}
func CreateUserService(userInfo *authSerializers.Credentials) uint {
	db := config.GetDb()
	user, password := userSerializers.CreateUserModels(userInfo)
	db.Create(user)
	db.Create(password)
	return user.ID

}
