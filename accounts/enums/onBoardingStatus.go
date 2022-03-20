package enums

type OnBoardingStatus string

const (
	OnBoardingMobile OnBoardingStatus = "on_boarding_mobile"
	OnBoardingEmail                   = "on_boarding_email"
	DetailsFilled                     = "details_filled"
	PictureUploaded                   = "picture_uploaded"
	PictureSkipped                    = "picture_skipped"
)
