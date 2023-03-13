package main

import (
	"go.etcd.io/etcd/client/v3"
	"time"
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379","192.168.108.171:2479","192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3.New error:", err)
		return
	}
	kv := clientv3.NewKV(client)
	putres, err := kv.Put(ctx, "ccc", "cdl")
	if err != nil {
		fmt.Println("kv put error:", err)
		return
	}
	fmt.Println("res:", putres)
}