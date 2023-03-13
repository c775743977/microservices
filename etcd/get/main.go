package main

import (
	"go.etcd.io/etcd/client/v3"
	"time"
	"fmt"
	"context"
)

type Service struct {
	Name string
	Port string
	Addr string
	Protocol string
}

func main() {
	var s = &Service{
		Name : "Greet",
	}
	s.ServiceDiscover()
	s.WatchService()
}

func(this *Service) ServiceDiscover() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("new client error:", err)
		return
	}
	kv := clientv3.NewKV(client)
	getRes, err := kv.Get(context.Background(), this.Name+".IP")
	if err != nil {
		fmt.Println("kv get IP error:", err)
		return
	}
	this.Addr = string(getRes.Kvs[0].Value)
	getRes, err = kv.Get(context.Background(), this.Name+".port")
	if err != nil {
		fmt.Println("kv get port error:", err)
		return
	}
	this.Port = string(getRes.Kvs[0].Value)
	getRes, err = kv.Get(context.Background(), this.Name+".protocol")
	if err != nil {
		fmt.Println("kv get protocol error:", err)
		return
	}
	this.Protocol = string(getRes.Kvs[0].Value)
}

func(this *Service) WatchService() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("create new client error:", err)
		return
	}
	watcher := clientv3.NewWatcher(client)
	watcherChan := watcher.Watch(context.Background(), this.Name, clientv3.WithPrefix())
	for watchRes := range watcherChan {
		for _, event := range watchRes.Events {
			if event.Type == clientv3.EventTypeDelete {
				this = nil
			}
			if event.Type == clientv3.EventTypePut {
				switch string(event.Kv.Key) {
				case this.Name:
					this.Name = string(event.Kv.Value)
					fmt.Println("serviceName 发生更改， 最新值为:", this.Name)
				case this.Name + ".IP":
					this.Addr = string(event.Kv.Value)
					fmt.Println("serviceIP 发生更改， 最新值为:", this.Addr)
				case this.Name + ".port":
					this.Port = string(event.Kv.Value)
					fmt.Println("servicePort 发生更改， 最新值为:", this.Port)
				case this.Name + ".protocol":
					this.Protocol = string(event.Kv.Value)
					fmt.Println("serviceProtocol 发生更改， 最新值为:", this.Protocol)
				}
			}
		}
	}
} 