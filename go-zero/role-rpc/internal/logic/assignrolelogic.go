package logic

import (
	"context"

	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleLogic {
	return &AssignRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分配角色
func (l *AssignRoleLogic) AssignRole(in *pb.AssignRoleReq) (*pb.AssignRoleResp, error) {
	err := l.svcCtx.RoleModel.AssignRoles(l.ctx, in.UserId, in.RoleIds)
	if err != nil {
		l.Logger.Errorf("分配角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.AssignRoleResp{
		Success: true,
	}, nil
}
