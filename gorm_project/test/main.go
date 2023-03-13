package main

import (
	"fmt"
	"gorm.io/gorm"
	_"time"
	"gorm.io/driver/mysql"
	_"errors"
)

type Book struct {
	ID int64
	Title string
	Author string
	Price float64
	Sales int64
	Stock int64
	Img_path string
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:Chen@123@/bookstore"), &gorm.Config{})
	var books []*Book
	db.Limit(4).Offset(8).Find(&books)
	for _, k := range books {
		fmt.Println(k)
	}
}