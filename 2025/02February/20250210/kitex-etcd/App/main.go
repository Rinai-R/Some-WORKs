package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client/UserClient"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Logger"
	"Golang/2025/02February/20250210/kitex-etcd/App/Middleware"
	"Golang/2025/02February/20250210/kitex-etcd/App/pkg/opentel"
	"Golang/2025/02February/20250210/kitex-etcd/App/router"
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	Logger.InitLogger()
	Client.InitETCD()
	UserClient.InitUserClient()
	Initialize.InitSentinel()
	// 初始化 OpenTelemetry Provider
	sdk, err := opentel.SetupOTelSDK(context.Background(), "api", "1.0.0")
	if err != nil {
		return
	}
	defer sdk(context.Background())
	h := server.Default()

	h.Use(Middleware.OpenTelemetryMiddleware())

	h.Use()
	router.InitRouter(h)
}
