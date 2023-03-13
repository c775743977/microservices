package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Login(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
}

func Register(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("Hello new %s", name))
}

func main() {
	r := gin.Default()
	user := r.Group("/user")
	user.GET("/login", Login)
	user.GET("/register", Register)
	book := r.Group("/book")
	book.GET("/login", Login)
	book.GET("/register", Register)
	r.Run()
}