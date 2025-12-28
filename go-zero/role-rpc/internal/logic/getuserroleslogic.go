package logic

import (
	"context"

	"go-zero-user/common/errorx"
	"go-zero-user/role-rpc/internal/svc"
	"go-zero-user/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户角色
func (l *GetUserRolesLogic) GetUserRoles(in *pb.GetUserRolesReq) (*pb.GetUserRolesResp, error) {
	roles, err := l.svcCtx.RoleModel.GetUserRoles(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("获取用户角色失败: %v", err)
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

	return &pb.GetUserRolesResp{
		Roles: roleInfos,
	}, nil
}
