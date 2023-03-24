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
	r.GET("/checkout", controller.CreateOrderHandler)
	r.GET("/MyOrders", controller.MyOrderHandler)
	r.GET("/showOrderItems", controller.MyOrderItemHandler)
	r.GET("/payOrder", controller.PayOrderHandler)
	r.GET("/signOrder", controller.SignOrderHandler)
	r.GET("/cancelOrder", controller.CancelOrderHandler)
	r.GET("/order_manager", controller.ShowAllOrdersHandler)
	r.GET("/deliverOrder", controller.DeliverOrderHandler)
	r.Run(":8083")
}