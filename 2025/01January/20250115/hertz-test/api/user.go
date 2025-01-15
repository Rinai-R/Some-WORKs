package api

import (
	"Golang/2025/01January/20250115/hertz-test/model"
	"Golang/2025/01January/20250115/hertz-test/response"
	"Golang/2025/01January/20250115/hertz-test/service"
	"Golang/2025/01January/20250115/hertz-test/utils"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
)

func Register(_ context.Context, ctx *app.RequestContext) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, response.Bind(err))
		return
	}
	if err = service.Register(user); err != nil {
		if errors.Is(err, response.ErrNameLength) {
			ctx.JSON(401, response.Register(err))
		} else if errors.Is(err, response.ErrPasswordLength) {
			ctx.JSON(401, response.Register(err))
		} else {
			ctx.JSON(402, response.Internal(err))
		}
		return
	}
	ctx.JSON(200, response.OK())
	return
}

func Login(_ context.Context, ctx *app.RequestContext) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, response.Bind(err))
		return
	}
	if err = service.Login(user); err != nil {
		if errors.Is(err, response.PasswordError) {
			ctx.JSON(403, response.Password())
			return
		}
	}
	Token, _ := utils.GenerateJWT(user.Name)
	ctx.JSON(200, response.OkWithData(Token))
}

func GetUserInfo(_ context.Context, ctx *app.RequestContext) {
	var user model.User
	GetName, exists := ctx.Get("GetName")
	if !exists {
		ctx.JSON(405, response.TokenError())
		return
	}
	user.Name = GetName.(string)
	if err := service.GetUserInfo(&user); err != nil {
		ctx.JSON(402, response.Internal(err))
		return
	}
	ctx.JSON(200, response.OkWithData(user))
	return

}
