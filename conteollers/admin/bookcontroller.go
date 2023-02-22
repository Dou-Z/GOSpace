package admin

import "github.com/gin-gonic/gin"

type BookController struct{}

func (con BookController) Index(c *gin.Context) {
	c.String(200, "我是admin-Book页")
}

func (con BookController) Add(c *gin.Context) {
	c.String(200, "我是admin-book-add页")
}

func (con BookController) Edit(c *gin.Context) {
	c.String(200, "我是admin-book-edit页")
}
