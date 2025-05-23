// Code generated by Kitex v0.12.1. DO NOT EDIT.

package user

import (
	base "Golang/2025/02February/20250209/kitex/kitex_gen/base"
	user "Golang/2025/02February/20250209/kitex/kitex_gen/user"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserRegisterArgs,
		newUserRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserLoginArgs,
		newUserLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceInfo                = NewServiceInfo()
	userServiceInfoForClient       = NewServiceInfoForClient()
	userServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "User"
	handlerType := (*user.User)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserRegisterArgs)
	realResult := result.(*user.UserRegisterResult)
	success, err := handler.(user.User).Register(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserRegisterArgs() interface{} {
	return user.NewUserRegisterArgs()
}

func newUserRegisterResult() interface{} {
	return user.NewUserRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserLoginArgs)
	realResult := result.(*user.UserLoginResult)
	success, err := handler.(user.User).Login(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserLoginArgs() interface{} {
	return user.NewUserLoginArgs()
}

func newUserLoginResult() interface{} {
	return user.NewUserLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, request *user.RegisterRequest) (r *base.Response, err error) {
	var _args user.UserRegisterArgs
	_args.Request = request
	var _result user.UserRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, request *user.LoginRequest) (r *base.Response, err error) {
	var _args user.UserLoginArgs
	_args.Request = request
	var _result user.UserLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
