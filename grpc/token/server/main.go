package main

import (
	"grpc/single-auth/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"net"
	"fmt"
	"context"
)

func Auth(ctx context.Context) error {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return fmt.Errorf("missing credentials")
    }
    var user string
    var password string

    if val, ok := md["user"]; ok {
        user = val[0]
    }
    if val, ok := md["password"]; ok {
        password = val[0]
    }

    if user != "cdl" || password != "chilang16" {
        return status.Errorf(codes.Unauthenticated, "token验证不通过!")
    }
    return nil
}

func main() {
	var authInterceptor grpc.UnaryServerInterceptor
    authInterceptor = func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (resp interface{}, err error) {
        //拦截普通方法请求，验证 Token
        err = Auth(ctx)
        if err != nil {
            return
        }
        // 继续处理请求
        return handler(ctx, req)
    }
    server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
    service.RegisterHelloServiceServer(server,service.HS)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error:", err)
		return
	}
	server.Serve(listener)
}