package user

import (
	"context"

	"go-zero-test/service/user/api/internal/svc"
	"go-zero-test/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.UserPageReq) (resp *types.UserPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
