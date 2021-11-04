package main

import (
	"fmt"
	pb "github.com/go-script/gRPCTest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)


type helloService struct {}


func (h helloService) SayHello (ctx context.Context, input *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", input.Name)
	return resp, nil
}


func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServer(s, &helloService{})

	fmt.Println("listen: 8080")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
