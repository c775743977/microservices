package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

type User struct {
	Name string
	Password string
}

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@tcp(localhost)/testdb"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/login", LoginHandler)
	r.Run(":8080")
}

func LoginHandler(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")
	var user User
	DB.Where("name = ?", name).Find(&user)
	if user.Password != pass {
		c.String(400, "login failed")
	} else {
		c.String(http.StatusOK, "login success!")
	}
}