package logic

import (
	"context"

	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"
	"go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 调用 user-rpc 注册接口
	result, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Id:       result.Id,
		Username: result.Username,
	}, nil
}
