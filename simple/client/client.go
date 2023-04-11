package main

import (
	"StudyGrpc/simple/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ClientConn, err := grpc.Dial("127.0.0.1:9876", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer ClientConn.Close()
	c := proto.NewGreeterClient(ClientConn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "你好"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}
