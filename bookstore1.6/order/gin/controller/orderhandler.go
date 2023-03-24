package controller

import (
	"github.com/gin-gonic/gin"

	"gin/rpc/rpcClient"
	"gin/rpc/user"
	"gin/rpc/cart"
	"gin/rpc/order"
	"gin/rpc/book"
	"gin/model"

	"context"
	"strconv"
	"fmt"
)

func CreateOrderHandler(c *gin.Context) {
	ctx := context.Background()
	userClient := rpcClient.NewUserClient()
	cartClient := rpcClient.NewCartClient()
	bookClient := rpcClient.NewBookClient()
	orderClient := rpcClient.NewOrderClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	sess, err := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, err := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	orderid, err := orderClient.CreateOrder(ctx, &order.OrderRes{
		TotalCount : cartData.Num,
		TotalAmount : cartData.Amount,
		UserID : cartData.UserID,
	})
	fmt.Println("orderid:", orderid)
	for _, k := range cartData.Items {
		bookData, _ := bookClient.GetBook(ctx, &book.BookReq{
			ID : k.BookID,
		})
		_, err = orderClient.CreateOrderItem(ctx, &order.OrderItem{
			OrderID : orderid.ID,
			Num : k.Num,
			Amount : k.Amount,
			Title : bookData.Title,
			Author : bookData.Author,
			Price : bookData.Price,
		})
	}
	cartClient.CleanCartItem(ctx, &cart.CartItem{
		CartID : cartData.ID,
	})
	cartClient.CleanCart(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	c.HTML(200, "cart/checkout.html", model.OrderData{
		UserName : sess.UserName,
		ID : orderid.ID,
	})
}

func MyOrderHandler(c *gin.Context) {
	ctx := context.Background()
	userClient := rpcClient.NewUserClient()
	orderClient := rpcClient.NewOrderClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	sess, err := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	orderData, err := orderClient.GetAllOrder(ctx, &order.OrderReqByUserID{
		UserID : sess.UserID,
	})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	c.HTML(200, "order/order.html", model.Orders{
		UserName : sess.UserName,
		Order : orderData.Orders,
	})
}

func MyOrderItemHandler(c *gin.Context) {
	ctx := context.Background()
	orderClient := rpcClient.NewOrderClient()
	orderid := c.Query("orderID")
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	orderitems, err := orderClient.GetOrderItem(ctx, &order.OrderItemReq{
		OrderID : orderid,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	fmt.Println("items:", orderitems.Items)
	c.HTML(200, "order/order_info.html", orderitems.Items)
}

func SignOrderHandler(c *gin.Context) {
	ctx := context.Background()
	orderClient := rpcClient.NewOrderClient()
	orderid := c.Query("orderID")
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	_, err := orderClient.AlterOrderStatus(ctx, &order.OrderRes{
		ID : orderid,
		Status : 4,
	})
	if err != nil {
		c.String(500, "服务器内部出错")
		return
	}
	MyOrderHandler(c)
}

func PayOrderHandler(c *gin.Context) {
	ctx := context.Background()
	orderClient := rpcClient.NewOrderClient()
	orderid := c.Query("orderID")
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	_, err := orderClient.AlterOrderStatus(ctx, &order.OrderRes{
		ID : orderid,
		Status : 1,
	})
	if err != nil {
		c.String(500, "服务器内部出错")
		return
	}
	MyOrderHandler(c)
}

func CancelOrderHandler(c *gin.Context) {
	ctx := context.Background()
	orderClient := rpcClient.NewOrderClient()
	orderid := c.Query("orderID")
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	_, err := orderClient.DelOrder(ctx, &order.OrderReqByID{
		ID : orderid,
	})
	if err != nil {
		c.String(500, "服务器内部出错")
		return
	}
	MyOrderHandler(c)
}

func ShowAllOrdersHandler(c *gin.Context) {
	ctx := context.Background()
	userClient := rpcClient.NewUserClient()
	orderClient := rpcClient.NewOrderClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	sess, err := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	userData, _ := userClient.GetUser(ctx, &user.UserReq{
		UserID : userid,
	})
	if userData.UserPrivilege != "Y" {
		c.String(400, "无权访问")
		return
	}
	orderData, err := orderClient.GetAllOrder(ctx, &order.OrderReqByUserID{})
	if err != nil {
		c.String(500, "服务区内部错误")
		return
	}
	c.HTML(200, "order/order_manager.html", orderData.Orders)
}

func DeliverOrderHandler(c *gin.Context) {
	ctx := context.Background()
	orderClient := rpcClient.NewOrderClient()
	orderid := c.Query("orderID")
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	_, err := orderClient.AlterOrderStatus(ctx, &order.OrderRes{
		ID : orderid,
		Status : 3,
	})
	if err != nil {
		c.String(500, "服务器内部出错")
		return
	}
	ShowAllOrdersHandler(c)
}