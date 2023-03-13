package main

import (
	"grpc/pbfile/service"
	"google.golang.org/protobuf/proto"
	"fmt"
)

func main() {
	user := &service.User{
		Username : "cdl",
		Age : 25,
	}

	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("marshal:", marshal)
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	fmt.Println(newUser.String())
}