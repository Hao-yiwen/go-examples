package logic

import (
	"context"
	"encoding/json"

	"go-zero-user/role-rpc/pb"
	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles() (resp *types.GetUserRolesResp, err error) {
	// 从 context 中获取用户ID
	userId, _ := l.ctx.Value("userId").(json.Number)
	userIdInt, _ := userId.Int64()

	// 调用 role-rpc 获取用户角色
	result, err := l.svcCtx.RoleRpc.GetUserRoles(l.ctx, &pb.GetUserRolesReq{
		UserId: userIdInt,
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

	return &types.GetUserRolesResp{
		Roles: roles,
	}, nil
}
