package main

import (
	"StudyGrpc/stream/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:40041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc Dial failed")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	//服务端流模式
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	for true {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
	//客户端流模式
	res1, _ := c.PostStream(context.Background())
	i := 0
	for true {
		i++
		res1.Send(&proto.StreamReqData{Data: fmt.Sprintf("慕课网%d", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式
	var wg sync.WaitGroup
	res2, _ := c.AllStream(context.Background())
	wg.Add(2)
	go func() {
		for {
			_ = res2.Send(&proto.StreamReqData{Data: string(time.Now().Unix())})
			time.Sleep(time.Second)
		}
		defer wg.Done()
	}()

	go func() {
		for true {
			a, err := res2.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(a.Data)
		}
		defer wg.Done()
	}()
	wg.Wait()

}
