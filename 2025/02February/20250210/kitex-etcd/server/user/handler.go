package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/base"
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user"
	"context"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *base.Response, err error) {
	//注册的业务逻辑
	return &base.Response{
		Code: 200,
		Msg:  "11111修改了",
	}, nil
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, request *user.LoginRequest) (resp *base.Response, err error) {
	//登陆的业务逻辑
	return &base.Response{
		Code: 200,
		Msg:  "原神?启动!",
	}, nil
}
