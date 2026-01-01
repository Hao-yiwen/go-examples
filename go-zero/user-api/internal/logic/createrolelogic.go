package logic

import (
	"context"

	"github.com/haoyiwen/go-examples/go-zero/role-rpc/pb"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.CreateRoleResp, err error) {
	// 调用 role-rpc 创建角色
	result, err := l.svcCtx.RoleRpc.CreateRole(l.ctx, &pb.CreateRoleReq{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateRoleResp{
		Id: result.Id,
	}, nil
}
