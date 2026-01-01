package logic

import (
	"context"
	"database/sql"

	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/internal/model"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色
func (l *UpdateRoleLogic) UpdateRole(in *pb.UpdateRoleReq) (*pb.UpdateRoleResp, error) {
	// 检查角色是否存在
	_, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(errorx.CodeRoleNotFound)
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 更新角色
	err = l.svcCtx.RoleModel.Update(l.ctx, &model.Role{
		Id:          in.Id,
		Name:        in.Name,
		Description: in.Description,
		Status:      in.Status,
	})
	if err != nil {
		l.Logger.Errorf("更新角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.UpdateRoleResp{
		Success: true,
	}, nil
}
