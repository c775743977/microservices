package main

import (
    "net/http"
	_"fmt"
    "github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
    // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
    User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
    Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
    // 1.创建路由
    // 默认使用了2个中间件Logger(), Recovery()
    r := gin.Default()
    // JSON绑定
	r.LoadHTMLFiles("./parse-form/index.html")
	r.GET("/login", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})
    r.POST("/login", func(c *gin.Context) {
        // 声明接收的变量
        var form Login
        // Bind()默认解析并绑定form格式
        // 根据请求头中content-type自动推断
        if err := c.Bind(&form); err != nil {
			//通过这种方式我们可以直接判断用户输入是否为空，不用像之前标准库的写法还需要读取数据来进行判断
			if err.Error() == "Key: 'Login.User' Error:Field validation for 'User' failed on the 'required' tag\nKey: 'Login.Pssword' Error:Field validation for 'Pssword' failed on the 'required' tag" {
				c.String(http.StatusBadRequest, "用户名和密码不能为空")
				return
			}
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // 判断用户名密码是否正确
        if form.User != "root" || form.Pssword != "admin" {
            // c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			c.String(http.StatusBadRequest, "密码或用户名错误")
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "200"})
		c.String(http.StatusOK, "登录成功")
    })
    r.Run()
}