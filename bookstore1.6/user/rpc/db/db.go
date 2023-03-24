package db

import (
	"gorm.io/driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB
var RDB *redis.ClusterClient
var err error

func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@tcp(localhost)/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}

	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs : []string{"192.168.108.165:6381", "192.168.108.165:6382", "192.168.108.165:6383"},
	})
}