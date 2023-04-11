package main

import (
	"StudyGrpc/stream/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const port = "127.0.0.1:40041"

var wg sync.WaitGroup

type server struct{}

// 服务端流模式
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamRespData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端流模式
func (s *server) PostStream(res proto.Greeter_PostStreamServer) error {
	for true {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
	return nil
}

// 双向流模式
func (s *server) AllStream(res proto.Greeter_AllStreamServer) error {
	wg.Add(2)
	go func() {
		for {
			_ = res.Send(&proto.StreamRespData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
		defer wg.Done()
	}()

	go func() {
		for true {
			a, err := res.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(a.Data)
		}
		defer wg.Done()
	}()
	wg.Wait()
	return nil
}

func main() {
	NetConn, err := net.Listen("tcp", port)
	if err != nil {
		panic("网络连接错误" + err.Error())
	}
	sGrpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(sGrpcServer, new(server))
	sGrpcServer.Serve(NetConn)

}
