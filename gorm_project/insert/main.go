package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_"time"
	"fmt"
  )

  type User struct {
	ID int
	Name string
	Email string
  }
//插入时会根据结构体的名字来选择表，比如结构体名为“User”，选择表就为“users”, “Student”->“students”
//如果结构体名称为复数形式，那么选择的表就不会再加s，仅仅将首字母小写 “Users”->“users”
  
func main() {
	db ,err := gorm.Open(mysql.Open("root:Chen@123@/testdb"))
	if err != nil {
		fmt.Println("connect database error:", err)
		return
	}
	user := User{
		Name : "cdl",
		Email : "775743977@.com",
	}
	result := db.Create(&user)
	fmt.Println("result:", result)
	fmt.Println("userID:", user.ID)

	user = User{
		Name : "root",
	}
	result = db.Create(&user)
	fmt.Println("result:", result)
	fmt.Println("userID:", user.ID)

	var users = []User{{Name : "aaa"}, {Name : "bbb"}, {Name : "ccc"}}  //批量插入
	result = db.Create(&users)
	fmt.Println("result:", result)
}