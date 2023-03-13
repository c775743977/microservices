package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)
 //中间件判断cookie是否存在
func MiddleWare(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	fmt.Println("cookie:", cookie)
	if cookie == "" {
		c.String(404, "error!!")
		c.Abort()
	}
}

func Login(c *gin.Context) {
	c.SetCookie("user", "cdl", 0, "/", "localhost", true, true)
	c.String(200, "login success!")
}

func Home(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	c.String(200, fmt.Sprintf("user: %s", cookie))
}

func main() {
	r := gin.Default()
	r.GET("/home", MiddleWare, Home)
	r.GET("/login", Login)
	r.Run()
}