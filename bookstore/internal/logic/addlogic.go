package logic

import (
	"context"

	"go-zero-demo/bookstore/internal/svc"
	"go-zero-demo/bookstore/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
