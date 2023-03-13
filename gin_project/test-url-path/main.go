package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func GetPath(c *gin.Context) {
	path := c.Param("username")
	c.String(200, fmt.Sprintf("path: %s", path))
}

func main() {
	r := gin.Default()
	r.GET("/home/:username", GetPath)
	//访问页面会显示出对应的"username"，即访问时在url输入的值
	r.Run()
}