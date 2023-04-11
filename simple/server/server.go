package main

import (
	"StudyGrpc/simple/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(in.Name)
	return &proto.HelloReply{
		Message: "hahahahaha",
	}, nil
}

func main() {
	netconn, err := net.Listen("tcp", "127.0.0.1:9876")
	if err != nil {
		fmt.Println("net conn failed")
		return
	}
	gNewServer := grpc.NewServer()
	proto.RegisterGreeterServer(gNewServer, new(GreeterServer))
	err = gNewServer.Serve(netconn)
	if err != nil {
		fmt.Println("server start failed")
		return
	}
}
