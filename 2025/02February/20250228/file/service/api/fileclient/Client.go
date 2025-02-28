package fileclient

import (
	pb "Golang/2025/02February/20250228/file/service/api/fileclient/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	FileClient pb.FileClient
	FileConn   *grpc.ClientConn
)

func InitClient() {

	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	err := ETCD.DiscoverService("File")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(ETCD.Services)
	addr := ETCD.GetService("File")
	FileConn, err = grpc.Dial(addr, opt)
	if err != nil {
		panic(err)
	}
	FileClient = pb.NewFileClient(FileConn)
	fmt.Println(FileClient)
}
