package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Pwd int
}

func HtmlTemplate(c *gin.Context) {
	user := &User{
		Name : "cdl",
		Pwd : 123,
	}
	c.HTML(http.StatusOK, "index.html", user) //也可以通过gin.H直接输入，替换user这个结构
	// c.HTML(http.StatusOK, "index.html", gin.H{"Name" : "cdl", "Pwd" : "123"}) //这里的字段就可以不用首字母大写了
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./test-html-template/index.html")
	r.GET("/home", HtmlTemplate)
	r.Run()
}