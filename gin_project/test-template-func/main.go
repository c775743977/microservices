package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"fmt"
	"html/template"
	"net/http"
)

func CreateOrderID() (orderid string) {
	t := time.Now()
	rand.Seed(time.Now().UnixNano())
	d := t.Day() + rand.Intn(99999)
	h := t.Hour() + rand.Intn(99999)
	m := t.Minute() + rand.Intn(99999)
	s := t.Second() + rand.Intn(99999)
	ns := t.Nanosecond() + rand.Intn(99999)
	return fmt.Sprint(d)+fmt.Sprint(s)+fmt.Sprint(ns)+fmt.Sprint(m)+fmt.Sprint(h)
}

func main() {
	
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"orderID" : CreateOrderID,
	}) //"orderID"要与前端对应
	r.LoadHTMLFiles("./test-template-func/index.html")	
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "cdl")
	})
	r.Run()
}