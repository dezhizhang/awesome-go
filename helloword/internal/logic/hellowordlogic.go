package logic

import (
	"context"

	"helloword/internal/svc"
	"helloword/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HellowordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHellowordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HellowordLogic {
	return &HellowordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HellowordLogic) Helloword(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
