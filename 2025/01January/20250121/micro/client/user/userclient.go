package rpc

import (
	"Golang/2025/01January/20250121/micro/client/nacos"
	pb "Golang/2025/01January/20250121/micro/client/user/proto"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
)

var (
	UserClient pb.UserClient
	UserConn   *grpc.ClientConn
)

func init() {
	var err error

	nacos.RegisterServiceInstance(nacos.Client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.10",                          // 服务实例的 IP 地址
		Port:        8848,                                // 服务实例的端口号
		ServiceName: "test",                               // 服务名称
		GroupName:   "group-a",                            // 分组名称
		ClusterName: "cluster-a",                          // 集群名称
		Weight:      10,                                   // 权重
		Enable:      true,                                 // 是否启用
		Healthy:     true,                                 // 是否健康
		Ephemeral:   true,                                 // 是否为临时实例
		Metadata:    map[string]string{"idc": "shanghai"}, // 元数据信息
	})

	UserConn, err = grpc.Dial("127.0.0.1:10001",
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024), grpc.MaxCallSendMsgSize(10*1024*1024)))

	if err != nil {
		panic(err)
	}
	UserClient = pb.NewUserClient(UserConn)
}
