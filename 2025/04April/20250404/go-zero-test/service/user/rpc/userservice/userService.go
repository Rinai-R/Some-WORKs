// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user.proto

package userservice

import (
	"context"

	"go-zero-test/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserLoginReq     = pb.UserLoginReq
	UserLoginResp    = pb.UserLoginResp
	UserRegisterReq  = pb.UserRegisterReq
	UserRegisterResp = pb.UserRegisterResp

	UserService interface {
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUserService) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}
