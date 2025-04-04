package logic

import (
	"context"

	"go-zero-test/service/user/rpc/internal/svc"
	"go-zero-test/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserRegisterResp{}, nil
}
