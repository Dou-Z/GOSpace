package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare(c *gin.Context) {
	// 判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	// 当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文 (c *gin.Context) 必须使用其只读副本 (c.Copy())
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path" + cCp.Request.URL.Path)
	}()

}
