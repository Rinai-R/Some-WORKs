package rpc

import (
	"Golang/2025/01January/20250121/micro/client/nacos"
	pb "Golang/2025/01January/20250121/micro/client/user/proto"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"log"
)

var (
	UserClient pb.UserClient
	UserConn   *grpc.ClientConn
)

func init() {
	var err error

	param := vo.GetServiceParam{
		ServiceName: "UserTest",            // 替换为你的服务名称
		GroupName:   "GroupTest",           // 根据需要设置
		Clusters:    []string{"cluster-a"}, // 集群名称
	}

	service, err := nacos.Client.GetService(param)
	if err != nil {
		log.Fatalf("can't discover the service: %v", err)
		return
	}

	// 获取第一个服务实例的地址
	if len(service.Hosts) == 0 {
		log.Fatal("no healthy instance found for service 'user'")
		return
	}

	addr := fmt.Sprintf("%s:%d", service.Hosts[0].Ip, service.Hosts[0].Port)

	// 连接到 gRPC 服务
	UserConn, err = grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024), grpc.MaxCallSendMsgSize(10*1024*1024)))

	if err != nil {
		log.Fatalf("failed to connect to gRPC service: %v", err)
		return
	}

	UserClient = pb.NewUserClient(UserConn)
}
