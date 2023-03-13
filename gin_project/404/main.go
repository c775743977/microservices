package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/user", func(c *gin.Context) {
        //指定默认值
        //http://localhost:8080/user 才会打印出来默认的值
        name := c.DefaultQuery("name", "default")
        c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
    })
    r.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "哎呀~页面丢失了")
    })
    r.Run()
}