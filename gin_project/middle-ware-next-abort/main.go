package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main(){
	router := gin.New()
 
	mid1 := func(c * gin.Context){
		fmt.Println("mid1 start")
		c.Next() //Next方法会使后面的中间件立刻开始执行，等待后面的中间件执行完毕后才会执行该中间件后面的代码
		fmt.Println("mid1 end")
	}
	mid2 := func(c * gin.Context){
		fmt.Println("mid2 start")
		c.Abort() //Abort方法会使该中间件执行完毕后结束中间件执行流程 注意GET方法也算一个中间件
		c.Next()
		fmt.Println("mid2 end")
	}
	mid3 := func(c * gin.Context){
		fmt.Println("mid3 start")
		c.Next()
		fmt.Println("mid3 end")
	}
	router.Use(mid1,mid2,mid3)
	router.GET("/",func(c * gin.Context){
		fmt.Println("process get request")
		c.JSON(http.StatusOK,"hello")
	})
	router.Run()
}