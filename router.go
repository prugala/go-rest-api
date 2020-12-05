package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-oauth2/google"
	"main/handler"
	"main/handler/oauth"
	"main/middleware"
	"os"
)

func registerRouter(engine *gin.Engine) {
	authMiddleware := middleware.AuthMiddleware()

	engine.GET("/refresh_token", handler.RefreshTokenHandler)

	oauthGroup := engine.Group("/oauth")
	oauthGroup.Use(google.Session(os.Getenv("SESSION_NAME")))
	oauthGroup.GET("/login/google", oauth.GoogleLoginHandler)
	oauthGroup.Use(google.Auth())
	oauthGroup.GET("/check/google", oauth.GoogleCheckHandler)

	engine.Use(authMiddleware.MiddlewareFunc())
	engine.NoRoute(handler.NotFoundHandler)
	engine.GET("/test", handler.TestHandler)
}
