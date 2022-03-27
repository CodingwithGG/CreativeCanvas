package authSerializers

import (
	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Email            string `json:"email"`
	Mobile           string `json:"mobile"`
	Password         string `json:"password"`
	OnBoardingStatus string `json:"on_boarding_status"`
}
type Claims struct {
	UserId  string
	IsAdmin bool
	jwt.StandardClaims
}
