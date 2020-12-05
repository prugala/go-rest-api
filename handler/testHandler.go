package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"main/model"
	"net/http"
	"os"
)

//@Description test endpoint
//@Security ApiKeyAuth
//@Accept  json
//@Produce  json
//@Success 200 {object} model.ExampleResponse "Example Response"
//@Failure 401 {object} model.ErrorResponse "Token is expired"
//@Router /test [get]
func TestHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	//user, _ := c.Get(os.Getenv("JWT_IDENTITY_KEY"))
	c.JSON(http.StatusOK, model.ExampleResponse{
		Sub: claims[os.Getenv("JWT_IDENTITY_KEY")].(string),
	})
}
