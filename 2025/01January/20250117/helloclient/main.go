package main

import (
	service "Golang/2025/01January/20250117/helloclient/client/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{ResponseMsg: "Hello " + req.RequestName}, nil
}

func main() {
	litsen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	grpcserver := grpc.NewServer()
	service.RegisterSayHelloServer(grpcserver, &server{})
	err = grpcserver.Serve(litsen)
	if err != nil {
		panic(err)
	}
}
