package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	proto "grpc/demo1/service/pbfile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go.etcd.io/etcd/client/v3"
	"fmt"
	"context"
	"time"
)

type Service struct {
	Name string
	Port string
	IP string
	Protocol string
}


func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	service := &Service{
		Name : "Login",
	}
	service.ServiceDiscover()
	go service.WatchService()
	conn, err := grpc.Dial(":" + service.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc dial 8081 error:", err)
		return
	}
	defer conn.Close()
	client := proto.NewLoginServiceClient(conn)
	r.GET("/index", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/login", func(c *gin.Context){
		name := c.PostForm("username")
		pass := c.PostForm("password")
		res, err := client.Login(context.Background(), &proto.LoginRequest{
			Name : name,
			Password : pass,
		})
		if err != nil {
			fmt.Println("client func login error:", err)
			return
		}
		c.String(http.StatusOK, res.Res)
	})
	r.Run(":8080")
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
	this.IP = string(getRes.Kvs[0].Value)
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
	svcName := this.Name
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
				case svcName:
					this.Name = string(event.Kv.Value)
					fmt.Println("serviceName 发生更改， 最新值为:", this.Name)
				case svcName + ".IP":
					this.IP = string(event.Kv.Value)
					fmt.Println("serviceIP 发生更改， 最新值为:", this.IP)
				case svcName + ".port":
					this.Port = string(event.Kv.Value)
					fmt.Println("servicePort 发生更改， 最新值为:", this.Port)
				case svcName + ".protocol":
					this.Protocol = string(event.Kv.Value)
					fmt.Println("serviceProtocol 发生更改， 最新值为:", this.Protocol)
				}
			}
		}
	}
} 