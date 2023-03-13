package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
  )
  
  func main() {
	_, err := gorm.Open(mysql.Open("root:Chen@123@/testdb"), &gorm.Config{}) //当前版本不再需要调用defer db.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
  }