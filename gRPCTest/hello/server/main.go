package main

import (
	"fmt"
	pb "github.com/go-script/gRPCTest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const Address = "127.0.0.1:50052"

type helloService struct {}

var HelloService = helloService{}


func (h helloService) SayHello (ctx context.Context, input *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", input.Name)
	return resp, nil
}


func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Error(err.Error())
	}

	s := grpc.NewServer()

	pb.RegisterHelloServer(s, HelloService)

	grpclog.Infoln("Listen on %s" +  Address)

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
