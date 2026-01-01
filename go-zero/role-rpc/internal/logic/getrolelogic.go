package logic

import (
	"context"
	"database/sql"

	"github.com/haoyiwen/go-examples/go-zero/common/errorx"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取角色
func (l *GetRoleLogic) GetRole(in *pb.GetRoleReq) (*pb.GetRoleResp, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(errorx.CodeRoleNotFound)
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.GetRoleResp{
		Role: &pb.RoleInfo{
			Id:          role.Id,
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt.Unix(),
		},
	}, nil
}
