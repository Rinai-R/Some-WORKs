package api

import (
	"Golang/2025/02February/20250209/kitex/App/Client/UserClient"
	"Golang/2025/02February/20250209/kitex/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Register(c context.Context, ctx *app.RequestContext) {
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
