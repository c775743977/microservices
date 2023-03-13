package main

import (
	"google.golang.org/grpc"
	"grpc/single-auth/service"
	"google.golang.org/grpc/credentials"
	"context"
	"fmt"
)

func main() {
	//"*.cdl.com"是openssl.cnf里设置的DNS
	// creds, err := credentials.NewClientTLSFromFile("E:/golang/go_code/src/grpc/credentials/server.pem", "*.cdl.com")
	// if err != nil {
	// 	fmt.Println("client get credentials error:", err)
	// 	return
	// }
	cert, _ := tls.LoadX509KeyPair("client/keys/test.pem", "client/keys/test.key")
    // 创建一个新的、空的 CertPool
    certPool := x509.NewCertPool()
    ca, _ := ioutil.ReadFile("client/keys/ca.crt")
    // 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
    certPool.AppendCertsFromPEM(ca)
    // 构建基于 TLS 的 TransportCredentials 选项
    creds := credentials.NewTLS(&tls.Config{
        // 设置证书链，允许包含一个或多个
        Certificates: []tls.Certificate{cert},
        // 要求必须校验客户端的证书。可以根据实际情况选用以下参数
        ServerName: "*.mszlu.com",
        RootCAs:    certPool,
    })
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("connect to 8080 error:", err)
		return
	}
	defer conn.Close()
	client := service.NewHelloServiceClient(conn)
	var helloreq = &service.HelloRequest{Req : "cdl",}
	res, err := client.SayHello(context.Background(), helloreq)
	if err != nil {
		fmt.Println("client.SayHello error:", err)
		return
	}
	fmt.Println(res)
}