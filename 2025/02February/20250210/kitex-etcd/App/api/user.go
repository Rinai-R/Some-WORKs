package api

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Client/UserClient"
	"Golang/2025/02February/20250210/kitex-etcd/Logger"
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Register(c context.Context, ctx *app.RequestContext) {
	Logger.Logger.Debug("Register")
	rsp, _ := UserClient.UserClient.Register(c, &user.RegisterRequest{
		Username: "rinai",
		Password: "123456",
	})
	if rsp.Code == 200 {
		ctx.JSON(http.StatusOK, rsp)
	}
}

func Login(c context.Context, ctx *app.RequestContext) {
	rsp, _ := UserClient.UserClient.Login(c, &user.LoginRequest{
		Username: "rinai",
		Password: "123456",
	})
	if rsp.Code == 200 {
		ctx.JSON(http.StatusOK, rsp)
	}
}
