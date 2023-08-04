package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-svelte/config"
	"net/http"
)

func AuthRedirects(r *gin.RouterGroup) {
	r.GET("/auth/login/google", RedirectToGoogleLogin)
}

func RedirectToGoogleLogin(c *gin.Context) {
	redirLoc := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&scope=email&response_type=code", config.Get().OAuth.ClientID, config.Get().OAuth.RedirectURI)
	c.Redirect(http.StatusSeeOther, redirLoc)
}
