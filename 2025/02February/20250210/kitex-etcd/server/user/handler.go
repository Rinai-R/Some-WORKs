package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/base"
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *base.Response, err error) {
	tracer := otel.Tracer("user-tracer")

	// 在Kitex服务中继续上下文传递
	_, span := tracer.Start(ctx, "user-span")
	defer span.End()

	// 设置一些属性
	span.SetAttributes(attribute.String("method", "MyMethod"))

	//注册的业务逻辑
	return &base.Response{
		Code: 200,
		Msg:  "11111修改了",
	}, nil
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, request *user.LoginRequest) (resp *base.Response, err error) {
	tracer := otel.Tracer("user-tracer")

	// 在Kitex服务中继续上下文传递
	_, span := tracer.Start(ctx, "user-span")
	defer span.End()

	// 设置一些属性
	span.SetAttributes(attribute.String("method", "MyMethod"))

	//登陆的业务逻辑
	return &base.Response{
		Code: 200,
		Msg:  "原神?启动!",
	}, nil
}
