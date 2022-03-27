package middlewares

//func RefreshTokenService(c *gin.Context) {
//	cookie, err := c.Cookie("jwt")
//
//	if err != nil {
//		response := config.BuildErrorResponse("error getting cookie", err.Error(), nil)
//		c.IndentedJSON(http.StatusUnauthorized, response)
//		return
//	}
//
//	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("SECRET_KEY")), nil
//	})
//
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			response := config.BuildErrorResponse("unauthorised", err.Error(), nil)
//			c.IndentedJSON(http.StatusUnauthorized, response)
//			return
//		}
//		response := config.BuildErrorResponse("bad request", err.Error(), nil)
//		c.IndentedJSON(http.StatusBadRequest, response)
//	}
//	if !token.Valid {
//		response := config.BuildErrorResponse("unauthorised", err.Error(), nil)
//		c.IndentedJSON(http.StatusUnauthorized, response)
//		return
//	}
//	//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
//	//	response := config.BuildErrorResponse("bad request", err.Error(), nil)
//	//	c.IndentedJSON(http.StatusBadRequest, response)
//	//	return
//	//}
//	expirationTime := time.Now().Add(time.Hour * 24).Unix()
//	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := refreshToken.SignedString(jwtKey)
//	if err != nil {
//		response := config.BuildErrorResponse("error generating new", err.Error(), nil)
//		c.IndentedJSON(http.StatusInternalServerError, response)
//		return
//	}
//
//	c.SetCookie("jwt", tokenString, int(expirationTime),
//		"/", "localhost", false, true)
//	return
//}
