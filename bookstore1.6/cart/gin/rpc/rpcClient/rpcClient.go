package rpcClient

import (
	"fmt"
	
	"gin/rpc/cart"
	"gin/rpc/user"
	"gin/rpc/book"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func NewBookClient() book.BookServiceClient {
	conn, err := grpc.Dial(bookRpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc dial %s error: %v", bookRpc, err)
		return nil
	}
	return book.NewBookServiceClient(conn)
}

func NewUserClient() user.UserServiceClient {
	conn, err := grpc.Dial(userRpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc dial %s error: %v", userRpc, err)
		return nil
	}
	return user.NewUserServiceClient(conn)
}