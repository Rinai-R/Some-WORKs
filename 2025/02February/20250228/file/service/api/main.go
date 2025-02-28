package main

import (
	"Golang/2025/02February/20250228/file/service/api/fileclient"
	"Golang/2025/02February/20250228/file/service/api/router"
)

func main() {
	fileclient.InitETCD()
	fileclient.InitClient()
	router.InitRouter()
}
