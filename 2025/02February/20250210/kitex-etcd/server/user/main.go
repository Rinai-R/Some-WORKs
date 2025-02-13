package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user/user"
	"Golang/2025/02February/20250210/kitex-etcd/server/Registry"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"net"
)

func main() {
	// 初始化 OpenTelemetry Provider
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("user"),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	// 解析服务地址
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	// 注册服务到 ETCD
	Registry.EtcdRegistry.ServiceRegister("user", "127.0.0.1:10001")

	// 创建 Kitex 服务
	svr := user.NewServer(
		&UserImpl{}, // 替换为您的服务实现
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr), // 服务地址
		//server.WithMiddleware(Middleware.OpenTelemetryMiddleware()),
	)

	// 启动服务
	if err := svr.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// 服务停止时注销服务
	defer Registry.EtcdRegistry.ServiceUnRegister("user")
	defer fmt.Println("Service stopped gracefully")
}
