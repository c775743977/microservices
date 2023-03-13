package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
)

type User struct {
	ID int
	Name string
	Email string
  }

func (u *User) AfterFind(db *gorm.DB) (err error) {
	if u.Name == "" {
		u.Name = "user"
	  }
	  return
}

func main() {
	db, err := gorm.Open(mysql.Open("root:Chen@123@/testdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
	user := User{Email : "zxc@111.com"}
	db.Create(&user)
	user = User{}
	db.Where("id=?", 8).Find(&user)
	fmt.Println("user:", user)
}