package logic

import (
	"context"

	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 角色列表
func (l *ListRoleLogic) ListRole(in *pb.ListRoleReq) (*pb.ListRoleResp, error) {
	page := in.Page
	pageSize := in.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	roles, err := l.svcCtx.RoleModel.FindList(l.ctx, page, pageSize)
	if err != nil {
		l.Logger.Errorf("查询角色列表失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	total, err := l.svcCtx.RoleModel.Count(l.ctx)
	if err != nil {
		l.Logger.Errorf("查询角色总数失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	var roleInfos []*pb.RoleInfo
	for _, role := range roles {
		roleInfos = append(roleInfos, &pb.RoleInfo{
			Id:          role.Id,
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt.Unix(),
		})
	}

	return &pb.ListRoleResp{
		Roles: roleInfos,
		Total: total,
	}, nil
}
