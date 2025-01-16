package main

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// 1. 创建 Nacos 客户端
	client := createNacosClient()

	// 2. 定义服务实例的元数据信息
	serviceName := "servicetest"   // 服务名称
	groupName := "group1"          // 分组名称
	clusterName := "cluster-a"     // 集群名称
	ip := "127.0.0.1"              // 服务实例的 IP 地址（当前虚拟机的 IP 地址）
	port := 8080                   // 服务实例的端口号
	metadata := map[string]string{ // 元数据信息
		"idc": "shanghai",
	}

	// 3. 注册服务实例到 Nacos
	registerServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
		GroupName:   groupName,
		ClusterName: clusterName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    metadata,
	})

	// 4. 启动 HTTP 服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is demo.go service at %s:%d!", ip, port)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: nil,
	}

	go func() {
		fmt.Printf("Starting HTTP server at http://%s:%d\n", ip, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server failed: %v\n", err)
		}
	}()

	// 5. 监听系统信号，实现优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// 6. 取消注册服务实例
	deRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
		GroupName:   groupName,
		Cluster:     clusterName,
		Ephemeral:   true,
	})

	// 7. 关闭 HTTP 服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("HTTP server shutdown failed: %v\n", err)
	}

	fmt.Println("Server exited gracefully")
}

// createNacosClient 创建 Nacos 客户端
func createNacosClient() naming_client.INamingClient {
	// 创建 ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"192.168.195.129",                  // Nacos 服务器的 IP 地址
			8848,                               // Nacos 服务器的端口号
			constant.WithContextPath("/nacos"), // Nacos 服务器的上下文路径
		),
	}

	// 创建 ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),              // 命名空间 ID
		constant.WithTimeoutMs(5000),              // 超时时间
		constant.WithNotLoadCacheAtStart(true),    // 启动时不加载缓存
		constant.WithLogDir("/tmp/nacos/log"),     // 日志目录
		constant.WithCacheDir("/tmp/nacos/cache"), // 缓存目录
		constant.WithLogLevel("debug"),            // 日志级别
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
	)

	// 创建命名服务客户端
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic("Failed to create Nacos client: " + err.Error())
	}

	return client
}

// registerServiceInstance 向 Nacos 注册一个服务实例
func registerServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		panic("Failed to register service instance: " + err.Error())
	}
	fmt.Printf("Registered service instance: %+v\n", param)
}

// deRegisterServiceInstance 从 Nacos 取消注册一个服务实例
func deRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, err := client.DeregisterInstance(param)
	if !success || err != nil {
		panic("Failed to deregister service instance: " + err.Error())
	}
	fmt.Printf("Deregistered service instance: %+v\n", param)
}
