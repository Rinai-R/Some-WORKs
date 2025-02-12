package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client/UserClient"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Logger"
	"Golang/2025/02February/20250210/kitex-etcd/App/router"
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func main() {
	Logger.InitLogger()
	Client.InitETCD()
	UserClient.InitUserClient()
	Initialize.InitSentinel()
	// 初始化 OpenTelemetry Provider
	p := provider.NewOpenTelemetryProvider(
		provider.WithInsecure(),         // 如果使用 HTTPS，可以去掉此选项
		provider.WithServiceName("api"), // 替换为你的服务名称
		provider.WithExportEndpoint("http://192.168.195.129:14268/api/traces"), // Jaeger 的地址
	)
	defer p.Shutdown(context.Background())
	h := server.Default()
	_, cfg := hertztracing.NewServerTracer()
	h.Use(hertztracing.ServerMiddleware(cfg))
	h.Use()
	router.InitRouter(h)
}
