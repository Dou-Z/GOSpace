package main

import (
	"dinGrom06/middleware"
	"dinGrom06/models"
	"dinGrom06/routers"
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 时间戳转换日期

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

func main() {
	r := gin.Default()
	// 配置全局中间件
	r.Use(middleware.InitMiddleWare)
	r.SetFuncMap(template.FuncMap{
		"unixToTime": models.UnixToTime,
		"Println":    Println,
	})
	// 加载模板
	r.LoadHTMLGlob("templates/**/**")
	r.Static("/static", "./static")

	// get请求
	routers.DefaultRouterInit(r)

	routers.AdminRouterInit(r)

	routers.ApiRouterInit(r)
	r.Run(":80")
}
