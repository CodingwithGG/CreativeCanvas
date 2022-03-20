package serializers

import "awesomeProject/accounts/enums"

type UserStatusRequest struct {
	Email  string `form:"email"`
	Mobile string `form:"mobile"`
}

type UserStatusResponse struct {
	OnBoardingStatus enums.OnBoardingStatus `form:"on_boarding_status"`
}
