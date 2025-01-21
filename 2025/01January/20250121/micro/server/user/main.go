package main

import (
	"Golang/2025/01January/20250121/micro/app/model"
	"Golang/2025/01January/20250121/micro/client/nacos"
	pb "Golang/2025/01January/20250121/micro/client/user/proto"
	"Golang/2025/01January/20250121/micro/response"
	"Golang/2025/01January/20250121/micro/server/user/dao"
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (U *UserServer) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.User{
		Name:     req.Username,
		Password: req.Password,
	}
	if len(user.Password) < 5 || len(user.Password) > 20 {
		return &pb.RegisterResponse{
			Success: false,
			Message: "PasswordLength",
		}, response.ErrPasswordLength
	}
	if len(user.Name) < 5 || len(user.Name) > 20 {
		return &pb.RegisterResponse{
			Success: false,
			Message: "NameLength",
		}, response.ErrNameLength
	}
	if err := dao.Register(user); err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}
	return &pb.RegisterResponse{
		Success: true,
		Message: "ok",
	}, nil
}

func (U *UserServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := model.User{
		Name:     req.Username,
		Password: req.Password,
	}
	if err := dao.Login(user); err != nil {
		return &pb.LoginResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}
	return &pb.LoginResponse{
		Success: true,
		Message: "ok",
	}, nil
}

func main() {
	// 设置监听地址和端口
	listener, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建gRPC服务器并设置消息大小限制
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	// 注册服务
	pb.RegisterUserServer(grpcServer, &UserServer{})

	// 启动Nacos服务注册
	nacos.RegisterServiceInstance(nacos.Client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.10",                          // 根据实际情况填写
		Port:        10001,                                // gRPC服务的端口
		ServiceName: "test",                               // 服务名称
		GroupName:   "group-a",                            // 分组名称
		ClusterName: "cluster-a",                          // 集群名称
		Weight:      10,                                   // 权重
		Enable:      true,                                 // 是否启用
		Healthy:     true,                                 // 是否健康
		Ephemeral:   true,                                 // 是否为临时实例
		Metadata:    map[string]string{"idc": "shanghai"}, // 元数据信息
	})

	// 启动服务
	log.Printf("Starting gRPC server on :10001")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	select {}
}
