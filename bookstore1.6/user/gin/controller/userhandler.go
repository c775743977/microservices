package controller

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"fmt"
	"context"

	"gin/rpc/RpcClient"
	"gin/rpc/user"
	"gin/rpc/cart"
)

func IndexHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.HTML(http.StatusOK, "index/index.html", nil)
	} else {
		userClient := RpcClient.NewUserClient()
		sess, err := userClient.GetSession(context.Background(), &user.Cookie{
			Cookie : cookie,
		})
		fmt.Println("sess:", sess)
		if err != nil {
			fmt.Println("call rpc-GetSession error:", err)
			c.HTML(http.StatusOK, "index/index.html", nil)
			return
		}
		c.HTML(http.StatusOK, "index/index.html", sess)
	}
}

func LoginHandler(c *gin.Context) {
	userClient := RpcClient.NewUserClient()
	var data user.LoginReq
	err := c.Bind(&data)
	if err != nil {
		c.HTML(400, "user/login.html", "无法提交空白数据")
		return
	}
	res, err := userClient.Login(context.Background(), &data)
	if err != nil {
		fmt.Println("call rpc-Login error:", err)
		return
	}
	if res.Result == "登录成功" {
		c.SetCookie("uuid", res.Cookie, 0, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "user/login_success.html", data.UserName)
	} else {
		c.HTML(http.StatusOK, "user/login.html", res.Result)
	}
}

func RegisterHandler(c *gin.Context) {
	userClient := RpcClient.NewUserClient()
	var data user.RegisterReq
	err := c.Bind(&data)
	if err != nil {
		c.HTML(400, "user/regist.html", "无法提交空白数据")
		return
	}
	var repwd = c.PostForm("repwd")
	fmt.Println("data:", data)
	fmt.Println("repwd:", repwd)
	if repwd != data.UserPassword {
		c.HTML(http.StatusOK, "user/regist.html", "两次输入的密码不一致")
		return
	}
	res, err := userClient.Register(context.Background(), &data)
	if err != nil {
		fmt.Println("call rpc-Register error:", err)
		return
	}
	if res.Result == "注册成功" {
		userData, _ := userClient.GetUserByName(context.Background(), &user.UserName{
			Name : data.UserName,
		})
		CreateCart(userData.UserID)
		c.HTML(http.StatusOK, "user/regist_success.html", data.UserName)
	} else {
		c.HTML(http.StatusOK, "user/regist.html", res.Result)
	}
}

func LogoutHandler(c *gin.Context) {
	userClient := RpcClient.NewUserClient()
	uuid, _ := c.Cookie("uuid")
	_, err := userClient.LogOut(context.Background(), &user.Cookie{
		Cookie : uuid,
	})
	if err != nil {
		fmt.Println("call rpc-Logout error:", err)
		return
	}
	c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	c.HTML(http.StatusOK, "index/index.html", nil)
}

func CreateCart(userid int64) {
	cartClient := RpcClient.NewCartClient()
	cartClient.CreateCart(context.Background(), &cart.CartReqByUserID{
		UserID : userid,
	})
}
