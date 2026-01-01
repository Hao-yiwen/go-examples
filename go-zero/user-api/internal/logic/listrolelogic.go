package logic

import (
	"context"

	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/pb"
	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoleLogic) ListRole(req *types.ListRoleReq) (resp *types.ListRoleResp, err error) {
	// 调用 role-rpc 获取角色列表
	result, err := l.svcCtx.RoleRpc.ListRole(l.ctx, &pb.ListRoleReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var roles []types.RoleInfo
	for _, r := range result.Roles {
		roles = append(roles, types.RoleInfo{
			Id:          r.Id,
			Name:        r.Name,
			Code:        r.Code,
			Description: r.Description,
			Status:      r.Status,
			CreatedAt:   r.CreatedAt,
		})
	}

	return &types.ListRoleResp{
		Roles: roles,
		Total: result.Total,
	}, nil
}
