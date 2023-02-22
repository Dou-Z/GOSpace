package admin

import (
	"dinGrom06/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (con AdminController) Index(c *gin.Context) {
	// c.String(200, "用户列表")
	// 查询数据
	userList := []models.User{}
	// 查询全部
	models.DB.Find(&userList)
	// 查询age大于20的
	// models.DB.Where("age>?", 20).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
	// c.HTML(http.StatusOK,"data/index.html",gin.H{
	// 	"result":userList,
	// })
}

func (con AdminController) Add(c *gin.Context) {
	userList := []models.User{{Username: "gorm", Age: 15, Email: "qwe123@qq.com", AddTime: 1238520}, {Username: "go", Age: 19, Email: "go123@qq.com", AddTime: 1235788}}
	models.DB.Create(&userList)
	c.String(200, "成功", gin.H{
		"result": "成功",
	})
}

func (con AdminController) Edit(c *gin.Context) {
	c.String(200, "更新")
}
func (con AdminController) Delete(c *gin.Context) {
	c.String(200, "删除")
}
