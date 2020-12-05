package oauth

import (
	"github.com/gin-gonic/gin"
	"main/middleware"
)

//@Description get JWT Token after successfully google login
//@Security ApiKeyAuth
//@Accept  json
//@Produce  json
//@Success 200 {object} model.JwtTokenResponse "JWT Token Response"
//@Failure 401 {object} model.ErrorResponse "failed to create JWT Token"
//@Router /oauth/check/google [get]
func GoogleCheckHandler(c *gin.Context) {
	middleware.AuthMiddleware().LoginHandler(c)
}
