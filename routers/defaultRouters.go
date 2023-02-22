package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultroute := r.Group("default")
	{
		defaultroute.GET("/", func(c *gin.Context) {
			// c.String(200, "我是Default首页")
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"title": "GROUP",
				"t":1671765352,
			})
		})
		defaultroute.GET("/news", func(c *gin.Context) {
			c.String(200, "我是Default-news页")
		})
		defaultroute.GET("/books", func(c *gin.Context) {
			c.String(200, "我是Default-book页")
		})
	}
}
