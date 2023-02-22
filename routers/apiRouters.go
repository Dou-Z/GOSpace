package routers

import (
	"dinGrom06/conteollers/itying"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiroute := r.Group("api")
	{
		apiroute.GET("/", func(c *gin.Context) {
			c.String(200, "我是API首页")
		})
		apiroute.GET("/news", func(c *gin.Context) {
			c.String(200, "我是API-news页")
		})
		apiroute.GET("/books", func(c *gin.Context) {
			c.String(200, "我是API-book页")
		})
		apiroute.GET("/user", itying.DefaultController{}.Index)
		// apiroute.POST("/user/add", itying.DefaultController{}.Doupload)

	}
}
