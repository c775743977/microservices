package etcd

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
	"context"
)

func KvDemo() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("handle error:", err)
		return
	}
	defer cli.Close()

	putRes, err := cli.Put(context.Background(), "k1", "v1", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println("put error:", err)
		return
	}
	fmt.Println("putres:", putRes)

	getRes, err := cli.Get(context.Background(), "greeting")
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	fmt.Println("getRes:", getRes)
}