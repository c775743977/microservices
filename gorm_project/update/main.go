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
	db.Model(&User{}).Where("id = ?", "6").Update("name", "hello") //UPDATE users SET name='hello' WHERE id=6;
	users := []User{}
	db.Find(&users)
	fmt.Println(users)

	user := User{
		ID : 2,
	}
	db.Model(&user).Update("email", "333@163.com") //UPDATE users SET email='333@163.com' WHERE id=2;
	users = []User{}
	db.Find(&users)
	fmt.Println(users)

	user = User{
		Name : "aaa",
	}
	db.Model(&user).Where("email", "aaa@11.com").Update("name", "zzz") //UPDATE users SET name='zzz' where name="aaa" and email="aaa@11.com";
	users = []User{}
	db.Find(&users)
	fmt.Println(users)
}