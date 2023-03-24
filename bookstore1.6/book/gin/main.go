package main

import (
	"github.com/gin-gonic/gin"
	"gin/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/**/*")
	r.Static("/static", "./static")
	r.Static("/pages", "./pages")
	r.GET("/index", controller.IndexHandler)
	r.GET("/logout", controller.LogoutHandler)
	r.POST("/index", controller.IndexHandler)
	r.GET("/admin", controller.ManageHandler)
	r.GET("/book_manager", controller.BookManageHandler)
	r.GET("/book_alter_add", controller.EditBookHandler)
	r.POST("/book_alter_add", controller.AddOrAlterBook)
	r.GET("/deletebook", controller.DelBookHandler)
	r.POST("/AddBookToCart", controller.AddItemHandler)
	r.Run(":8081")
}