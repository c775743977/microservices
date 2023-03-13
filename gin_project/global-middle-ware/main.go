package main

import (
    "fmt"
    "time"
	"net/http"
    "github.com/gin-gonic/gin"
)

// 定义中间
func MiddleWare() gin.HandlerFunc { //工厂模式
    return func(c *gin.Context) {
        t := time.Now()
        fmt.Println("中间件1开始执行了")
        // 设置变量到Context的key中，可以通过Get()取
        c.Set("request1", "中间件1")
        status := c.Writer.Status()
        fmt.Println("中间件1执行完毕", status)
        t2 := time.Since(t)
        fmt.Println("time:", t2)
    }
}

func Middle(c *gin.Context) {
	t := time.Now()
	fmt.Println("中间件2开始执行了")
	// 设置变量到Context的key中，可以通过Get()取
	c.Set("request", "中间件2")
	status := c.Writer.Status()
	fmt.Println("中间件2执行完毕", status)
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}
//当存在两个中间件时，会依次执行，最后执行的会覆盖之前的（前提是有冲突）
func main() {
    // 1.创建路由
    // 默认使用了2个中间件Logger(), Recovery()
    r := gin.Default()
    // 注册中间件
    r.Use(MiddleWare())  //这个是定义全局中间件，即后面每次REST的执行都会经过中间件
    // {}为了代码规范
    {
        r.GET("/ce", func(c *gin.Context) {
            // 取值
            req, _ := c.Get("request1")
            fmt.Println("request:", req)
            // 页面接收
            c.JSON(200, gin.H{"request": req})
        })

		r.GET("/cc", Middle, func(c *gin.Context) {  //这样的中间件只有局部生效
            // 取值
            req, _ := c.Get("request")
			req1, _ := c.Get("request1")
            fmt.Println("request:", req)
            // 页面接收
            c.String(http.StatusOK, fmt.Sprintf("中间件已执行 %s", req))
			c.String(http.StatusOK, fmt.Sprintf("中间件1已执行 %s", req1))
        })

    }
    r.Run()
}