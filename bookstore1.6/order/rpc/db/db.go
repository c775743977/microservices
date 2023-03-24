package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@tcp(localhost)/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
}