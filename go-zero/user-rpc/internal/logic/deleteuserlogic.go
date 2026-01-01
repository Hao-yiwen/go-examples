package logic

import (
	"context"

	"github.com/haoyiwen/go-examples/go-zero/common/errorx"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteUserLogic) DeleteUser(in *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	err := l.svcCtx.UserModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("删除用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.DeleteUserResp{
		Success: true,
	}, nil
}
