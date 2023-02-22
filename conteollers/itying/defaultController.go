package itying

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// fmt.Println(models.UnixToTime(1629788418))
	c.SetCookie("username", "Douz", 3600, "/", "www..com", false, true)

	c.SetCookie("hobby", "吃饭，睡觉奥", 5, "/", "127.0.0.1", false, true)

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "admin",
	})

}

func (con DefaultController) News(c *gin.Context) {
	// 获取cookies
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")
	c.String(200, "cookie="+username+"---hobby---"+hobby)
}

func (con DefaultController) Shop(c *gin.Context) {
	// 获取cookies
	username, _ := c.Cookie("username")
	c.String(200, "cookie="+username)
}
