package main

//
import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
	r.LoadHTMLFiles("./test-html-postform/index.html")
	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//可以看到使用REST就可以对同一个地址有不同的操作，而根据GO标准库来写，一个地址只能对应一个Handler
    r.POST("/form", func(c *gin.Context) {
        types := c.DefaultPostForm("type", "post")
        username := c.PostForm("username")
        password := c.PostForm("userpassword")
        // c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
        c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
    })
    r.Run()
}