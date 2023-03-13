package service

import (
	"time"
	"go.etcd.io/etcd/client/v3"
	"fmt"
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"sync"
)

type Service struct {
	Name string
	IP string
	Port string
	Protocol string
}

func ServiceRegist(s *Service) {
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
	var grantLease bool
	var leaseID clientv3.LeaseID
	ctx := context.Background()
	//clientv3.WithCountOnly()会使结果只返回查询到的数量
	getRes, err := cli.Get(ctx, s.Name, clientv3.WithCountOnly())
	if err != nil {
		fmt.Println("cli get error:", err)
		return
	}
	//如果数量为0表示该服务还没有注册，也说明还没有创建租约
	if getRes.Count == 0 {
		grantLease = true
	}
	//如果租约没有创建那么就创建一个租约
	if grantLease {
		leaseRes, err := cli.Grant(ctx, 10)
		if err != nil {
			fmt.Println("cli grant error:", err)
			return
		}
		leaseID = leaseRes.ID
	}
	kv := clientv3.NewKV(cli)
	txn := kv.Txn(ctx)
	// Revision作用域为集群，逻辑时间戳，全局单调递增，任何 key 的增删改都会使其自增
	// CreateRevision作用域为 key, 等于创建这个 key 时集群的 Revision, 直到删除前都保持不变
	//如果CreateRevision(s.Name)返回0，那么说明该服务还没有被注册
	txn = txn.If(clientv3.Compare(clientv3.CreateRevision(s.Name), "=", 0))
	//如果该服务还没有注册就进行注册并附加租约
	txn = txn.Then(
		clientv3.OpPut(s.Name, s.Name, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s.Name + ".ip", s.IP, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s.Name + ".port", s.Port, clientv3.WithLease(leaseID)),
		clientv3.OpPut(s.Name + ".protocol", s.Protocol, clientv3.WithLease(leaseID)),
	)
	//clientv3.WithIgnoreLease()表示沿用之前的租约，如果服务已经注册，则更新服务的信息
	txn = txn.Else(
		clientv3.OpPut(s.Name, s.Name, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s.Name + ".ip", s.IP, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s.Name + ".port", s.Port, clientv3.WithIgnoreLease()),
		clientv3.OpPut(s.Name + ".protocol", s.Protocol, clientv3.WithIgnoreLease()),
	)
	_, err = txn.Commit()
	if err != nil {
		fmt.Println("txn commit error:", err)
		return
	}
	//如果是第一次创建租约，则对租约进行自动租约
	if grantLease {
		leaseKeepAlive, err := cli.KeepAlive(ctx, leaseID)
		if err != nil {
			fmt.Println("cli.KeepAlive error:", err)
			return
		}
		for lease := range leaseKeepAlive {
			fmt.Printf("leaseID:%x, ttl:%d\n", lease.ID, lease.TTL)
		}
	}
}

type Services struct {
	services map[string]*Service
	sync.RWMutex
}

var myServices = &Services{
	services : map[string]*Service{},
}

var ms = &Services{}

func ServiceDiscover(svcName string) *Service {
	// var s *Service = nil
	myServices.RLock()
	s, _ := myServices.services[svcName]
	myServices.RUnlock()
	return s
}

func WatchServiceName(svcName string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints : []string{"192.168.108.171:2379", "192.168.108.171:2479", "192.168.108.171:2579"},
		DialTimeout : 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3.New error:", err)
		return
	}
	defer cli.Close()

	getRes, err := cli.Get(context.Background(), svcName, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("cli get error:", err)
		return
	}
	if getRes.Count > 0 {
		mp := SliceToMap(getRes.Kvs)
		s := &Service{}
		if kv, ok := mp[svcName]; ok {
			s.Name = string(kv.Value)
		}
		if kv, ok := mp[svcName+".ip"]; ok {
			s.IP = string(kv.Value)
		}
		if kv, ok := mp[svcName+".port"]; ok {
			s.Port = string(kv.Value)
		}
		if kv, ok := mp[svcName+".protocol"]; ok {
			s.Protocol = string(kv.Value)
		}
		myServices.Lock()
		myServices.services[svcName] = s
		myServices.Unlock()
	}

	wch := cli.Watch(context.Background(), svcName, clientv3.WithPrefix())
	for wres := range wch {
		for _, ev := range wres.Events {
			if ev.Type == clientv3.EventTypeDelete{
				myServices.Lock()
				delete(myServices.services, svcName)
				myServices.Unlock()
			}
			if ev.Type == clientv3.EventTypePut{
				myServices.Lock()
				switch string(ev.Kv.Key) {
				case svcName:
					myServices.services[svcName].Name = string(ev.Kv.Value)
				case svcName+".ip":
					myServices.services[svcName].IP = string(ev.Kv.Value)
				case svcName+".port":
					myServices.services[svcName].Port = string(ev.Kv.Value)
				case svcName+".protocol":
					myServices.services[svcName].Protocol = string(ev.Kv.Value)
				}
				myServices.Unlock()
			}
		}
	}
} 

func SliceToMap(list []*mvccpb.KeyValue) map[string]*mvccpb.KeyValue {
	mp := make(map[string]*mvccpb.KeyValue)
	for _, item := range list {
		mp[string(item.Key)] = item
	}
	return mp
}