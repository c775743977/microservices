package main

import (
	"github.com/gin-gonic/gin"
	"gin/controller"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/**/*")
	r.Static("/static", "./static")
	r.Static("/pages", "./pages")
	r.GET("/index", controller.IndexHandler)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/login.html", nil)
	})
	r.POST("/login", controller.LoginHandler)
	r.GET("/regist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/regist.html", nil)
	})
	r.POST("/regist", controller.RegisterHandler)
	r.Run(":8080")
}