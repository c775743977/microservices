package main

import (
	"google.golang.org/grpc"
	proto "etcd/example/pbfile"
	"go.etcd.io/etcd/client/v3"
	"net"
	"fmt"
	"flag"
	"context"
	"time"
	"os"
)

const schema = "ms"
 
var host = "127.0.0.1" //服务器主机
var (
    Port        = flag.String("port", "30020", "listening port")                           //服务器监听端口
    ServiceName = flag.String("serviceName", "greet_service", "service name")        //服务名称
    EtcdAddr    = flag.String("etcdAddr", "192.168.108.171:2379", "register etcd address") //etcd的地址
)

type Service struct {
	Name string
	Port string
	Addr string
	Protocol string
}

type GreetServer struct {}
var gs = &GreetServer{}

func(this *GreetServer) Morning(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Println("Morning 调用:", req.Req)
	return &proto.GreetResponse{
		Res : "Good Morning " + req.Req,
		From : fmt.Sprintf("127.0.0.1:%s", *Port),
	}, nil
}

func(this *GreetServer) Night(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Println("Night 调用:", req.Req)
	return &proto.GreetResponse{
		Res : "Good Night " + req.Req,
		From : fmt.Sprintf("127.0.0.1:%s", *Port),
	}, nil
}

func(s1 *Service) KeepAlive(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3 new error:", err)
		return
	}
	defer client.Close()
	lease := clientv3.NewLease(client)
	kv := clientv3.NewKV(client)
	txn := kv.Txn(ctx)
	var leaseID clientv3.LeaseID
	fmt.Println("service:", s1)
	getRes, err := kv.Get(ctx, s1.Name, clientv3.WithCountOnly())
	if err != nil {
		fmt.Println("kv get error:", err)
		return
	}
	if getRes.Count == 0 {
		fmt.Println("service:", s1, "开始注册")
		leaseRes, err := lease.Grant(ctx, 10)
		fmt.Println("注册的leaseID:", leaseRes.ID)
		if err != nil {
			fmt.Println("lease grant error:", err)
			return
		}
		leaseID = leaseRes.ID
	}
	txn.If(clientv3.Compare(clientv3.CreateRevision(s1.Name), "=", 0)).Then(
		clientv3.OpPut(s1.Name, s1.Name, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s1.Name + ".port", s1.Port, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s1.Name + ".IP", s1.Addr, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s1.Name + ".protocol", s1.Protocol, clientv3.WithLease(leaseID)),
	).Else(
		clientv3.OpPut(s1.Name, s1.Name, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s1.Name + ".port", s1.Port, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s1.Name + ".IP", s1.Addr, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s1.Name + ".protocol", s1.Protocol, clientv3.WithIgnoreLease()),
	).Commit()
	ResChan, err := lease.KeepAlive(ctx, leaseID)
	if err != nil {
		fmt.Println("lease keepalive error:", err)
		return
	} else {
		for {
			select {
			case lease := <-ResChan :
				if lease == nil {
					fmt.Printf("租约:%v 已过期\n", leaseID)
					delres, err := kv.Delete(ctx, "Greet", clientv3.WithPrefix())
					if err != nil {
						fmt.Println("kv delete error:", err)
						return
					}
					fmt.Println("已清除过期数据:", delres)
					fmt.Println("请重新运行")
					return
				} else {
					fmt.Printf("租约:%v 续约成功\n", leaseID)
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	s1 := &Service{
		Name : *ServiceName,
		Port : *Port,
		Addr :  GetIP(),
		Protocol : "grpc",
	}
	go s1.KeepAlive(context.Background())
	server := grpc.NewServer()
	proto.RegisterGreetServer(server, gs)
	listener, err := net.Listen("tcp", ":" + s1.Port)
	if err != nil {
		fmt.Println("net listen to 30020 error:", err)
		return
	}
	server.Serve(listener)
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ip := ""
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}

		}
	}
	return ip
}