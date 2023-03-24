package controller

import (
	"context"
	"strconv"
	"fmt"
	_"net/http"

	"github.com/gin-gonic/gin"

	"gin/rpc/rpcClient"
	"gin/rpc/book"
	"gin/rpc/user"
)

func IndexHandler(c *gin.Context) {
	var data book.PageReq
	min := c.PostForm("min")
	max := c.PostForm("max")
	if min == "" && max == "" {
		min = c.Query("min")
		max = c.Query("max")
	}
	data.PageNo, _ = strconv.ParseInt(c.DefaultQuery("PageNo", "1"), 10, 64)
	data.MaxPrice, _ = strconv.ParseFloat(max, 64)
	data.MinPrice, _ = strconv.ParseFloat(min, 64)
	userClient := rpcClient.NewUserClient()
	bookClient := rpcClient.NewBookClient()
	books, err := bookClient.GetPage(context.Background(), &data)
	if err != nil {
		fmt.Println("call rpc-GetPage error:", err)
		c.String(500, "index/index.html", "访问页面崩溃了:(  努力抢修中~")
		return
	}
	cookie, _ := c.Cookie("uuid")
	if cookie != "" {
		res, err := userClient.GetSession(context.Background(), &user.Cookie{
			Cookie : cookie,
		})
		if err != nil {
			fmt.Println("call rpc-GetSession error:", err)
			c.HTML(400, "index/index.html", books)
			return
		}
		books.UserName = res.UserName
		c.HTML(200, "index/index.html", books)
	} else {
		c.HTML(200, "index/index.html", books)
	}
}

func LogoutHandler(c *gin.Context) {
	userClient := rpcClient.NewUserClient()
	uuid, _ := c.Cookie("uuid")
	_, err := userClient.LogOut(context.Background(), &user.Cookie{
		Cookie : uuid,
	})
	if err != nil {
		fmt.Println("call rpc-Logout error:", err)
		return
	}
	c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	IndexHandler(c)
}
