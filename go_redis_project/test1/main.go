package main

import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "192.168.108.165:6381",
		Password: "", // 没有密码，默认值
		DB:		  0,  // 默认DB 0
	})
	fmt.Println(rdb)
	ctx := context.Background()
	err := rdb.HSet(ctx , "cookie", "ID", "2").Err()
	if err != nil {
		fmt.Println("error:", err)
	} 
	err = rdb.HSet(ctx , "cookie", "User", "root").Err()
	if err != nil {
		fmt.Println("error:", err)
	} 
	user, err := rdb.HGet(ctx, "cookie", "User").Result()
	if err != nil {
		fmt.Println("error:", err)
	} 
	fmt.Println(user)
	err = rdb.Set(ctx, "k1", "v1", -1).Err()
	if err != nil {
		fmt.Println("error:", err)
	} 
	v, err := rdb.Get(ctx, "k1").Result()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("v=", v)
	}
}