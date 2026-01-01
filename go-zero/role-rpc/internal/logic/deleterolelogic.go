package logic

import (
	"context"

	"github.com/haoyiwen/go-examples/go-zero/common/errorx"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeleteRoleLogic) DeleteRole(in *pb.DeleteRoleReq) (*pb.DeleteRoleResp, error) {
	err := l.svcCtx.RoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("删除角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.DeleteRoleResp{
		Success: true,
	}, nil
}
