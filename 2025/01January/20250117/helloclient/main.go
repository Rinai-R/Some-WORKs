package main

import (
	service "Golang/2025/01January/20250117/helloclient/client/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata")
	}
	var name string
	var password string
	if v, ok := md["name"]; ok {
		name = v[0]

	}
	if v, ok := md["password"]; ok {
		password = v[0]
	}
	if name != "admin" || password != "admin" {
		return nil, errors.New("invalid metadata")
	}
	return &service.HelloResponse{ResponseMsg: "Hello " + req.RequestName}, nil
}

func main() {
	//cred, _ := credentials.NewServerTLSFromFile("2025/01January/20250117/helloclient/key/test.pem",
	//	"2025/01January/20250117/helloclient/key/test.key")

	litsen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	grpcserver := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	service.RegisterSayHelloServer(grpcserver, &server{})
	err = grpcserver.Serve(litsen)
	if err != nil {
		panic(err)
	}
}
