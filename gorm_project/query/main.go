package main

import (
	"fmt"
	"gorm.io/gorm"
	_"time"
	"gorm.io/driver/mysql"
	_"errors"
)

type User struct {
	ID int
	Name string
	Email string
  }

func main() {
	db, err := gorm.Open(mysql.Open("root:Chen@123@/testdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
	user := User{}
	db.First(&user) //SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println("first:", user) 

	user = User{} //同一个变量再次使用需要先置空
	db.Take(&user) //SELECT * FROM users LIMIT 1;
	fmt.Println("take:", user)

	user = User{}
	db.Last(&user) //SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("last:", user)

	users := []User{}
	db.Find(&users) //SELECT * FROM users;
	fmt.Println("Find:", users)

	user = User{}
	db.First(&user, 5) //SELECT * FROM users WHERE id = 10;
	fmt.Println("result:", user)

	user = User{}
	db.Where("name = ?", "cdl").First(&user) //SELECT * FROM users WHERE name = 'cdl' limit 1;
	fmt.Println("result1:", user)

	users = []User{}
	db.Where("name = ?", "cdl").Find(&users) //SELECT * FROM users WHERE name = 'cdl';
	fmt.Println("result2:", users)
}