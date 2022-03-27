package userSerializers

import (
	"awesomeProject/accounts/enums"
	"awesomeProject/accounts/models"
	"awesomeProject/authentication/models"
	"awesomeProject/authentication/serializers"
	"golang.org/x/crypto/bcrypt"
)

type UserStatusRequest struct {
	Email  string `form:"email"`
	Mobile string `form:"mobile"`
}

type UserStatusResponse struct {
	OnBoardingStatus userEnums.OnBoardingStatus `form:"on_boarding_status"`
}

func CreateUserModels(userSerializer *authSerializers.Credentials) (*userModels.User, *authModels.Password) {
	var user userModels.User
	var password authModels.Password
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(userSerializer.Password), 14)
	password.Password = encryptedPassword
	user.VerificationStatus = userEnums.Unverified
	if userSerializer.Mobile != "" {
		user.Mobile = &userSerializer.Mobile
		password.Mobile = &userSerializer.Mobile
		user.OnBoardingStatus = userEnums.OnBoardingMobile
	} else {
		password.Email = &userSerializer.Email
		user.Email = &userSerializer.Email
		user.OnBoardingStatus = userEnums.OnBoardingEmail
	}

	return &user, &password
}
