package main

import (
	"grpc/test/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"fmt"
	"context"
)

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial 8002 error:", err)
		return
	}
	client := service.NewProductServiceClient(conn)
	request := &service.ProductRequest{
		ProdId : 12306,
	}
	stock, err := client.GetStock(context.Background(), request)
	fmt.Println("stock:", stock.ProdStock)
	res, err := client.CheckStock(context.Background(), stock)
	fmt.Println("check result:", res.Res)
}