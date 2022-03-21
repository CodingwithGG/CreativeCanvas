package serializers

import (
	"awesomeProject/accounts/enums"
	"awesomeProject/accounts/models"
)

type UserStatusRequest struct {
	Email  string `form:"email"`
	Mobile string `form:"mobile"`
}

type UserStatusResponse struct {
	OnBoardingStatus enums.OnBoardingStatus `form:"on_boarding_status"`
}

type CreateUser struct {
	Email            string                 `form:"email"`
	Mobile           string                 `form:"mobile"`
	Password         string                 `form:"password"`
	OnBoardingStatus enums.OnBoardingStatus `form:"on_boarding_status"`
}

func CreateUserModels(userSerializer *CreateUser) (*models.User, *models.Password) {
	var user models.User
	var password models.Password
	password.Password = userSerializer.Password
	password.Mobile = userSerializer.Mobile
	password.Email = userSerializer.Email
	user.Mobile = userSerializer.Mobile
	user.Email = userSerializer.Email
	user.OnBoardingStatus = userSerializer.OnBoardingStatus
	user.VerificationStatus = enums.Unverified
	return &user, &password
}
