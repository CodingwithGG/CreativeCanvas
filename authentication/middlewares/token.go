package middlewares

import (
	"awesomeProject/authentication/serializers"
	"awesomeProject/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// AuthorizeJWT validates and authorizes the requests
func AuthorizeJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")

		if err != nil {
			response := config.BuildErrorResponse("error getting cookie", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claims := &authSerializers.Claims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response := config.BuildErrorResponse("unauthorised", err.Error(), nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			response := config.BuildErrorResponse("bad request", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if !token.Valid {
			response := config.BuildErrorResponse("unauthorised", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}
		c.Set("user_id", claims.UserId)
		c.Set("is_admin", claims.IsAdmin)
	}
}
