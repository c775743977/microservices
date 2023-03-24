package main

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/**/*")
	r.Static("/static", "./static")
	r.Static("/pages", "./pages")
	r.GET("/index", controller.IndexHandler)
	r.GET("/cart", controller.MyCartHandler)
	r.POST("/AddBookToCart", controller.AddItemHandler)
	r.GET("/delete_item", controller.DelItemHandler)
	r.POST("/alterNum", controller.AlterItemNumHandler)
	r.GET("/CleanUp", controller.CleanCartHandler)
	r.Run(":8082")
}