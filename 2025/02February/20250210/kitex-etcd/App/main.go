package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client/UserClient"
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Logger"
	"Golang/2025/02February/20250210/kitex-etcd/App/router"
)

func main() {
	Logger.InitLogger()
	Client.InitETCD()
	UserClient.InitUserClient()
	Initialize.InitSentinel()

	router.InitRouter()
}
