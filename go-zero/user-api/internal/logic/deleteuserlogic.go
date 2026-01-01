package logic

import (
	"context"

	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/types"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// 调用 user-rpc 删除用户
	result, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &pb.DeleteUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteUserResp{
		Success: result.Success,
	}, nil
}
