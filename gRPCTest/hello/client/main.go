package main

import (
	"fmt"
	pb "github.com/go-script/gRPCTest/proto" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8080"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "wfc"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		panic(err)
	}


	fmt.Println(res.Message)
}
