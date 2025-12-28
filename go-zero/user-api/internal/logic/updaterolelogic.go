package logic

import (
	"context"

	"go-zero-user/role-rpc/pb"
	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	// 调用 role-rpc 更新角色
	result, err := l.svcCtx.RoleRpc.UpdateRole(l.ctx, &pb.UpdateRoleReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateRoleResp{
		Success: result.Success,
	}, nil
}
