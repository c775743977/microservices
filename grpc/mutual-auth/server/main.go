package main

import (
	"grpc/single-auth/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"tls"
	"net"
	"fmt"
	"log"
)

func main() {
	// creds, err := credentials.NewServerTLSFromFile("E:/golang/go_code/src/grpc/credentials/server.pem", "E:/golang/go_code/src/grpc/credentials/server.key")
	// if err != nil {
	// 	fmt.Println("credentials.NewServerTLSFromFile error:", err)
	// 	return
	// }
	cert, err := tls.LoadX509KeyPair("keys/mszlu.pem", "keys/mszlu.key")
    if err != nil {
        log.Fatal("证书读取错误",err)
    }
    // 创建一个新的、空的 CertPool
    certPool := x509.NewCertPool()
    ca, err := ioutil.ReadFile("keys/ca.crt")
    if err != nil {
        log.Fatal("ca证书读取错误",err)
    }
    // 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
    certPool.AppendCertsFromPEM(ca)
    // 构建基于 TLS 的 TransportCredentials 选项
    creds := credentials.NewTLS(&tls.Config{
        // 设置证书链，允许包含一个或多个
        Certificates: []tls.Certificate{cert},
        // 要求必须校验客户端的证书。可以根据实际情况选用以下参数
        ClientAuth: tls.RequireAndVerifyClientCert,
        // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
        ClientCAs: certPool,
    })

	server := grpc.NewServer(grpc.Creds(creds))
	service.RegisterHelloServiceServer(server, service.HS)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}