package main

import (
	"time"
	"go.etcd.io/etcd/client/v3"
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2379", "192.168.108.171:2379"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3.New error:", err)
		return
	}
	kv := clientv3.NewKV(client)
	//clientv3.WithPrevKV()能够返回执行本次操作之前的会涉及到的值，如果不添加该选项则delres中看不到删除的数据
	delres, err := kv.Delete(ctx, "Greet", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("kv delete error:", err)
		return
	}
	fmt.Println("delres:", delres)
}