package RpcClient

import (
	"gin/rpc/user"
	"gin/rpc/cart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
)

var bookRpc = "192.168.0.143:8002"
var userRpc = "192.168.0.143:8001"
var cartRpc = "192.168.108.157:8003"

func NewCartClient() cart.CartServiceClient {
	conn, err := grpc.Dial(cartRpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("rpc dial %s error:%v", cartRpc, err)
	}
	return cart.NewCartServiceClient(conn)
}


func NewUserClient() user.UserServiceClient {
	conn, err := grpc.Dial(userRpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc dial %s error: %v", userRpc, err)
		return nil
	}
	return user.NewUserServiceClient(conn)
}