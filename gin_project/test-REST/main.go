package main

import "github.com/gin-gonic/gin"
import "net/http"

func GetInfo(c *gin.Context) {
	c.JSON(202, gin.H{
		"message" : "GET INFO",
	})
}

func PostInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "POST INFO",
	})
}

func PutInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "PUT INFO",
	})
}

func DelInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "DEL INFO",
	})
}

func main() {
	r := gin.Default()
	//REST风格的写法可以通过不同的访问方式来访问同一个地址，但是能够获取不同的结果
	r.GET("/test", GetInfo)
	r.POST("/test", PostInfo)
	r.PUT("/test", PutInfo)
	r.DELETE("/test", DelInfo)
	r.Run(":80") // 监听并在 0.0.0.0:8080 上启动服务
}