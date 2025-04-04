package logic

import (
	"context"

	"go-zero-test/service/user/rpc/internal/svc"
	"go-zero-test/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserLoginResp{}, nil
}
