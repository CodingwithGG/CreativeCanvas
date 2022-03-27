package userModels

import (
	"awesomeProject/accounts/enums"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name               *string                      `json:"type:varchar(100);name"`
	Email              *string                      `gorm:"type:varchar(100);unique_index;json:email"`
	VerificationStatus userEnums.VerificationStatus `json:"verification_status"`
	OnBoardingStatus   userEnums.OnBoardingStatus   `json:"on_boarding_status"`
	IsActive           bool                         `json:"is_active" gorm:"default:true"`
	FirstName          *string                      `json:"first_name"`
	LastName           *string                      `json:"last_name"`
	IsStaff            bool                         `json:"is_staff"`
	Gender             userEnums.Gender             `json:"gender"`
	Mobile             *string                      `gorm:"unique_index;json:mobile"`
	ActiveAt           time.Time                    `json:"active_at"`
	Location           string                       `json:"location"`
	ProfilePic         string                       `json:"profile_pic"`
	RatingCount        int                          `gorm:"default:0;json:rating_count"`
	RatedBy            string                       `json:"rated_by"`
	IsCreator          bool                         `json:"is_creator"`
}
