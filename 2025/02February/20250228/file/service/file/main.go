package main

import (
	"Golang/2025/02February/20250228/file/service/file/Registry"
	Handle "Golang/2025/02February/20250228/file/service/file/handle"
	pb "Golang/2025/02February/20250228/file/service/file/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	pb.RegisterFileServer(grpcServer, &Handle.FileService{})

	msg := fmt.Sprintf("grpc server listening at %v", listener.Addr())
	Registry.EtcdRegistry.ServiceRegister("File", "127.0.0.1:10001")
	fmt.Println(msg)

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	defer grpcServer.GracefulStop()
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)
}
