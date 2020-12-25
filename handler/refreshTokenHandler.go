package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/middleware"
)

//@Description refresh JWT Token
//@Security ApiKeyAuth
//@Accept  json
//@Produce  json
//@Success 200 {object} model.JwtTokenResponse "JWT Token Response"
//@Failure 401 {object} model.ErrorResponse "Token is expired"
//@Router /refresh_token [get]
func RefreshTokenHandler(c *gin.Context) {
	fmt.Println("test")
	middleware.AuthMiddleware().RefreshHandler(c)
}
