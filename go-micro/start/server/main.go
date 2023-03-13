package main

import (
    "context"
    "fmt"

    micro "github.com/asim/go-micro/v3"
    proto "github.com/asim/go-micro/service/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
    rsp.Greeting = "Hello " + req.Name
    return nil
}

func main() {
    // Create a new service. Optionally include some options here.
    service := micro.NewService(
        micro.Name("greeter"),
    )

    // Init will parse the command line flags.
    service.Init()

    // Register handler
    proto.RegisterGreeterHandler(service.Server(), new(Greeter))

    // Run the server
    if err := service.Run(); err != nil {
        fmt.Println(err)
    }
}