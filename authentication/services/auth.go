package authServices

import (
	"awesomeProject/accounts/enums"
	"awesomeProject/accounts/models"
	"awesomeProject/accounts/services"
	"awesomeProject/authentication/models"
	"awesomeProject/authentication/serializers"
	"awesomeProject/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetTokenService(credentials *authSerializers.Credentials, c *gin.Context) (config.Response, int) {
	var SecretKey = os.Getenv("SECRET_KEY")
	var jwtKey = []byte(SecretKey)
	db := config.GetDb()
	var credentialsObj *gorm.DB
	var userObj *gorm.DB
	var password authModels.Password
	var user userModels.User
	var userId uint
	if credentials.OnBoardingStatus == string(userEnums.OnBoardingStarted) {
		userId = userServices.CreateUserService(credentials)
		log.Println("New user is created")
	} else {
		if credentials.Mobile != "" {
			credentialsObj = db.Model(&authModels.Password{}).Where("mobile = ?", credentials.Mobile).First(&password)
			userObj = db.Model(&userModels.User{}).Where("mobile = ?", credentials.Mobile).First(&user)

		} else if credentials.Email != "" {
			credentialsObj = db.Model(&authModels.Password{}).Where("email = ?", credentials.Email).First(&password)
			userObj = db.Model(&userModels.User{}).Where("email = ?", credentials.Mobile).First(&user)
		} else {
			return config.BuildErrorResponse("give either mobile or email", "", nil),
				http.StatusBadRequest
		}
		if credentialsObj.Error != nil {
			return config.BuildErrorResponse("internal server error", credentialsObj.Error.Error(), nil),
				http.StatusInternalServerError
		}
		if userObj.Error != nil {
			return config.BuildErrorResponse("internal server error", credentialsObj.Error.Error(), nil),
				http.StatusInternalServerError
		}
		expectedPassword := password.Password
		if err := bcrypt.CompareHashAndPassword(expectedPassword, []byte(credentials.Password)); err != nil {
			return config.BuildErrorResponse("password incorrect", err.Error(), nil),
				http.StatusInternalServerError
		}
		userId = user.ID
	}
	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	claims := &authSerializers.Claims{
		UserId: strconv.Itoa(int(userId)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return config.BuildErrorResponse("error generating token", err.Error(), nil),
			http.StatusInternalServerError
	}

	c.SetCookie("jwt", tokenString, int(expirationTime),
		"/", "localhost", false, true)
	return config.BuildResponse("login successful", nil), http.StatusOK
	//http.SetCookie(w,
	//	&http.Cookie{
	//		Name:    "token",
	//		Value:   tokenString,
	//		Expires: expirationTime,
	//	})

}
func LogoutService(c *gin.Context) (config.Response, int) {
	c.SetCookie("jwt", "", int(time.Now().Add(-time.Hour).Unix()),
		"/", "localhost", false, true)

	return config.BuildResponse("logout successful", nil), http.StatusOK
}
