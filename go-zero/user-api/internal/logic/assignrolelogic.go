package logic

import (
	"context"

	"github.com/haoyiwen/go-examples/go-zero/role-rpc/pb"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleLogic {
	return &AssignRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRoleLogic) AssignRole(req *types.AssignRoleReq) (resp *types.AssignRoleResp, err error) {
	// 调用 role-rpc 分配角色
	result, err := l.svcCtx.RoleRpc.AssignRole(l.ctx, &pb.AssignRoleReq{
		UserId:  req.UserId,
		RoleIds: req.RoleIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.AssignRoleResp{
		Success: result.Success,
	}, nil
}
