package main

import (
	"fmt"
	"context"
	"time"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : time.Second * 5,
	})
	if err != nil {
		fmt.Println("connect to etcd error:", err)
		return
	}

	kv := clientv3.NewKV(client)

	go func(){
		for {
			kv.Put(context.Background(), "/demo/A/B1", "I am B1")
			kv.Delete(context.Background(), "/demo/A/B1")
			time.Sleep(3 * time.Second)
		}
	}()
	getRes, err := kv.Get(context.Background(), "/demo/A/B1")
	if err != nil {
		fmt.Println("kv Get error:", err)
		return
	}

	if len(getRes.Kvs) != 0 {
		fmt.Println("当前值:", string(getRes.Kvs[0].Value))
	}
	watchStartRevision := getRes.Header.Revision + 1

	// 创建一个watcher
	watcher := clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从该版本向后监听:", watchStartRevision)

	// 创建一个 5s 后取消的上下文
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	// 该监听动作在 5s 后取消
	watchRespChan := watcher.Watch(ctx, "/demo/A/B1", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", 
								event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}