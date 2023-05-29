package web

import (
	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine) *gin.Engine {
	r.Static("/", "static/")
	/*r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})*/

	return r
}
