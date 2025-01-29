package api

import (
	"Golang/2025/01January/20250121/micro/app/client/user"
	pb "Golang/2025/01January/20250121/micro/app/client/user/proto"
	"Golang/2025/01January/20250121/micro/app/model"
	"Golang/2025/01January/20250121/micro/app/utils"
	"Golang/2025/01January/20250121/micro/response"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

func Register(ctx context.Context, r *app.RequestContext) {
	var user model.User
	err := r.BindJSON(&user)
	if err != nil {
		r.JSON(400, response.Bind(err))
		return
	}
	if res, err := rpc.UserClient.Register(ctx, &pb.RegisterRequest{
		Username: user.Name,
		Password: user.Password,
	}); err != nil {
		if res == nil {
			log.Println(err)
		}
		if errors.Is(err, response.ErrNameLength) {
			r.JSON(401, response.Register(err))
		} else if errors.Is(err, response.ErrPasswordLength) {
			r.JSON(401, response.Register(err))
		} else {
			r.JSON(402, response.Internal(err))
		}
		return
	}
	r.JSON(200, response.OK())
	return
}

func Login(ctx context.Context, r *app.RequestContext) {
	var user model.User
	err := r.BindJSON(&user)
	if err != nil {
		r.JSON(400, response.Bind(err))
		return
	}
	if _, err = rpc.UserClient.Login(ctx, &pb.LoginRequest{
		Username: user.Name,
		Password: user.Password,
	}); err != nil {
		if errors.Is(err, response.PasswordError) {
			r.JSON(403, response.Password())
		} else {
			r.JSON(402, response.Internal(err))
		}
		return
	}
	Token, _ := utils.GenerateJWT(user.Name)
	r.JSON(200, response.OkWithData(Token))
}
