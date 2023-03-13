package main

import (
	"google.golang.org/grpc"
	proto "grpc/demo1/service/pbfile"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"go.etcd.io/etcd/client/v3"
	"context"
	"fmt"
	"net"
	"os"
	"flag"
	"time"
)

var DB *gorm.DB
var err error
var ls = &LoginService{}
var (
    Port        = flag.String("port", "8081", "listening port")                           //服务器监听端口
    ServiceName = flag.String("serviceName", "Login", "service name")        //服务名称
    EtcdAddr    = flag.String("etcdAddr", "192.168.108.171:2379", "register etcd address") //etcd的地址
)

type LoginService struct {

}

type User struct {
	Name string
	Password string
}

type Service struct {
	Name string
	IP string
	Port string
	Protocol string
}


func(this *LoginService) Login(ctx context.Context, lr *proto.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("收到客户端数据:", lr)
	var user = User{}
	DB.Where("name = ?", lr.Name).Find(&user)
	if user.Password != lr.Password {
		return &proto.LoginResponse{
			Res : "login failed!",
		}, nil
	} else {
		return &proto.LoginResponse{
			Res : "login succeeded!",
		}, nil
	}
}

func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@tcp(localhost)/testdb"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
	}
}

func main() {
	flag.Parse()
	service := &Service{
		Name : *ServiceName,
		Port : *Port,
		IP :  GetIP(),
		Protocol : "grpc",
	}
	go service.RegisterInEtcd(context.Background())
	server := grpc.NewServer()
	proto.RegisterLoginServiceServer(server, ls)
	listener, err := net.Listen("tcp", ":"+*Port)
	if err != nil {
		fmt.Println("listen to :8081 error:", err)
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

func(service *Service) RegisterInEtcd(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("v3 new client error:", err)
		return
	}
	defer client.Close()
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	txn := kv.Txn(ctx)
	var leaseID clientv3.LeaseID
	getRes, err := kv.Get(ctx, service.Name, clientv3.WithCountOnly())
	if err != nil {
		fmt.Println("kv Get error:", err)
		return
	}
	if getRes.Count == 0 {
		fmt.Println("service:", service, "开始注册")
		leaseRes, err := lease.Grant(ctx, 10)
		fmt.Println("注册的leaseID:", leaseRes.ID)
		if err != nil {
			fmt.Println("lease grant error:", err)
			return
		}
		leaseID = leaseRes.ID
	}
	txn.If(clientv3.Compare(clientv3.CreateRevision(service.Name), "=", 0)).Then(
		clientv3.OpPut(service.Name, service.Name, clientv3.WithLease(leaseID)),
		clientv3.OpPut(service.Name + ".port", service.Port, clientv3.WithLease(leaseID)),
		clientv3.OpPut(service.Name + ".IP", service.IP, clientv3.WithLease(leaseID)),
		clientv3.OpPut(service.Name + ".protocol", service.Protocol, clientv3.WithLease(leaseID)),
	).Else(
		clientv3.OpPut(service.Name, service.Name, clientv3.WithIgnoreLease()),
		clientv3.OpPut(service.Name + ".port", service.Port, clientv3.WithIgnoreLease()),
		clientv3.OpPut(service.Name + ".IP", service.IP, clientv3.WithIgnoreLease()),
		clientv3.OpPut(service.Name + ".protocol", service.Protocol, clientv3.WithIgnoreLease()),
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