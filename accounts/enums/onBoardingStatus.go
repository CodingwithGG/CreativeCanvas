package userEnums

type OnBoardingStatus string

const (
	OnBoardingStarted OnBoardingStatus = "on_boarding_started"
	OnBoardingMobile                   = "on_boarding_mobile"
	OnBoardingEmail                    = "on_boarding_email"
	DetailsFilled                      = "details_filled"
	PictureUploaded                    = "picture_uploaded"
	PictureSkipped                     = "picture_skipped"
)
