package routers

import (
	"dinGrom06/conteollers/admin"
	"dinGrom06/conteollers/itying"
	"dinGrom06/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminroute := r.Group("/", middleware.InitMiddleWare)
	{
		adminroute.GET("/itying", itying.DefaultController{}.Index)
		adminroute.GET("/", admin.AdminController{}.Index)
		adminroute.GET("/add", admin.AdminController{}.Add)
		adminroute.GET("/edit", admin.AdminController{}.Edit)
		adminroute.GET("/del", admin.AdminController{}.Delete)

	}

}
