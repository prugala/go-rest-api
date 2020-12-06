package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zalando/gin-oauth2/google"
	_ "main/docs"
	"os"
)

//@title Example API
//@version 1.0
//@description This is a example API with oauth and jwt.
//@termsOfService http://swagger.io/terms/
//
//@contact.name Piotr Rugala
//@contact.url https://github.com/prugala
//@contact.email piotr@isedo.pl
//
//@license.name MIT License
//@license.url https://opensource.org/licenses/MIT
//
//@securityDefinitions.apikey JSON Web Token
//@in header
//@name Authorization
//@host localhost:8080
//@BasePath /
func main() {
	hostname := os.Getenv("URL_SCHEME") + os.Getenv("HOST") + ":" + os.Getenv("PORT")
	gin.SetMode(os.Getenv("GIN_MODE"))
	g := gin.Default()

	google.Setup(
		hostname+os.Getenv("OAUTH_REDIRECT_URL"),
		os.Getenv("OAUTH_CREDENTIALS_FILE"),
		[]string{os.Getenv("OAUTH_SCOPE")},
		[]byte(os.Getenv("SECRET")),
	)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	registerRouter(g)

	g.Run(":" + os.Getenv("PORT"))
}
