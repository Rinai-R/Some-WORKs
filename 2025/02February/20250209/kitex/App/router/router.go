package router

import (
	"Golang/2025/02February/20250209/kitex/App/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	r := server.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", api.Register)
		userGroup.POST("/login", api.Login)

	}

	r.Spin()
}
