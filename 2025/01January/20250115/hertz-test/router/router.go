package router

import (
	"Golang/2025/01January/20250115/hertz-test/MiddleWare"
	"Golang/2025/01January/20250115/hertz-test/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	r := server.Default()

	r.POST("/Register", api.Register)

	r.POST("/Login", api.Login)

	r.Use(MiddleWare.Token())

	r.GET("/GetUserInfo", api.GetUserInfo)

	r.Spin()
}
