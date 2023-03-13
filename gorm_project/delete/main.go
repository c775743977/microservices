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

func main() {
	db, err := gorm.Open(mysql.Open("root:Chen@123@/testdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
	//如果不指定主键或者任何字段则会删除所有数据
	// user := User{}
	// db.Delete(&user)

	user := User{ID:4,}
	db.Delete(&user) //delete from users where id=4;
	users := []User{}
	db.Find(&users)
	fmt.Println(users)
}