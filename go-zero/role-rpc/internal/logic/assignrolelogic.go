package logic

import (
	"context"

	"go-zero-user/common/errorx"
	"go-zero-user/role-rpc/internal/svc"
	"go-zero-user/role-rpc/pb"

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
