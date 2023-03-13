package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

func RunTime(c *gin.Context) {
	res, _ := c.Get("RunTime")
	c.String(http.StatusOK, fmt.Sprintf("%s\nPATH:%s", res, c.Request.URL))
}

func MiddleWare(c *gin.Context) {
	now := time.Now()
	status := c.Writer.Status()
	fmt.Println("waiting...")
	gap := time.Since(now)
	fmt.Println("time:",gap)
	c.Set("RunTime", fmt.Sprintf("RunTime: %s\tstatus:%v", gap, status))
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare)
	user1 := r.Group("/u1")
	user2 := r.Group("/u2")
	user1.GET("/test", RunTime)
	user2.GET("/test", RunTime)
	r.Run()
}