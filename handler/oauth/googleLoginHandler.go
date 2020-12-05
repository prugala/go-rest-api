package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-oauth2/google"
	"net/http"
)

//@Description oauth google login
//@Security ApiKeyAuth
//@Accept  json
//@Produce  json
//@Success 301
//@Router /oauth/login/google [get]
func GoogleLoginHandler(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, google.GetLoginURL(state))
}
