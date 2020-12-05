package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-oauth2/google"
	"log"
	"os"
	"strconv"
	"time"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	jwtTokenTimeout, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_TIMEOUT"))
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       os.Getenv("APP_NAME"),
		Key:         []byte(os.Getenv("SECRET")),
		Timeout:     time.Hour * time.Duration(jwtTokenTimeout),
		MaxRefresh:  time.Hour,
		IdentityKey: os.Getenv("JWT_IDENTITY_KEY"),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*google.User); ok {
				return jwt.MapClaims{
					os.Getenv("JWT_IDENTITY_KEY"): v.Sub,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &google.User{
				Sub: claims[os.Getenv("JWT_IDENTITY_KEY")].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			data, exists := c.Get("user")
			user, ok := data.(google.User)

			if !exists || !ok {
				return "", jwt.ErrFailedTokenCreation
			}

			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*google.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	return authMiddleware
}
