package admin

import "github.com/gin-gonic/gin"

type NewController struct{}

func (con NewController) NewIndex(c *gin.Context) {
	c.String(200, "我是admin-news页")
}

func (con NewController) NewAdd(c *gin.Context) {
	c.String(200, "我是admin-news-add页")
}

func (con NewController) NewEdit(c *gin.Context) {
	c.String(200, "我是admin-news-edit页")
}
