package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client/UserClient"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Logger"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"

	"Golang/2025/02February/20250210/kitex-etcd/App/router"
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

func main() {
	Logger.InitLogger()
	Client.InitETCD()
	UserClient.InitUserClient()
	Initialize.InitSentinel()
	// 初始化 OpenTelemetry Provider
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("api"),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := hertztracing.NewServerTracer()
	h := server.Default(tracer)
	h.Use(hertztracing.ServerMiddleware(cfg))
	router.InitRouter(h)
}
