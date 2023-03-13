package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		// name := c.Query("name")
		//DefaultQuery会设置一个默认值，当url中没有找到对应参数时就会使用这个默认值
		name := c.DefaultQuery("name", "default-name")
		c.String(200, fmt.Sprintf("NAME: %s", name))
	})
	r.Run()
}