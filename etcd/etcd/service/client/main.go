package main

import (
	"google.golang.org/grpc"
	"go-zero/etcd/service/pbfile"
	"go-zero/etcd/service"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
	"context"
	"flag"
	"time"
)

func getServerAddr(svcName string) string {
	s := service.ServiceDiscover(svcName)
	if s == nil {
		return ""
	}
	return s.IP + ":" + s.Port
}

var name = flag.String("name", "root", "input your name")

func main() {
	go service.WatchServiceName("SayHello")
	for {
		time.Sleep(time.Second * 2)
		addr := getServerAddr("SayHello")
		if addr == "" {
			fmt.Println("未发现可用服务")
			continue
		}
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("grpc dial 8080 error:", err)
			continue
		}
		flag.Parse()
		client := pbfile.NewHelloServiceClient(conn)
		res, err := client.SayHello(context.Background(), &pbfile.HelloRequest{
			Req : *name,
		})
		if err != nil {
			fmt.Println("client calls SayHello error:", err)
			continue
		}
		fmt.Println(res)
	}
}

