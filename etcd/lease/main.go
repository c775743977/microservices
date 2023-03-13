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
	//新建一个lease
	lease := clientv3.NewLease(client)
	//批准这个lease，10为TTL(单位秒)
	leaseRes, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Println("lease grant error:", err)
		return
	}
	//自动续约，每次续约都会往返回的管道中推送最新的续约数据
	keepRespChan, err := lease.KeepAlive(ctx, leaseRes.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效了")
					return
				} else { // 每秒会续租一次, 所以就会受到一次应答
					fmt.Println("收到自动续租应答:", keepResp.ID, "TTL:", keepResp.TTL)
				}
			}
		}
	}()
	kv := clientv3.NewKV(client)
	putres, err := kv.Put(ctx, "ccc", "CDL", clientv3.WithLease(leaseRes.ID))
	if err != nil {
		fmt.Println("kv put error:", err)
		return
	}
	fmt.Println("写入成功:", putres)

	// 定时的看一下key过期了没有
	for {
		getResp, err := kv.Get(context.TODO(), "ccc")
		if err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
		fmt.Println("当前租约剩余时间:", leaseRes.TTL)
	}
}