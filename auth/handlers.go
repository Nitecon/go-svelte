package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/auth/login/google", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login route",
		})
	})

	// Add more auth routes here
}
