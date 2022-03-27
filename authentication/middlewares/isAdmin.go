package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckIfSystemAdmin middleware checks if the user is system admin
func CheckIfSystemAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("is_admin")
		if role == false {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
