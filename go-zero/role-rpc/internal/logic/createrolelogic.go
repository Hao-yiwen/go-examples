package logic

import (
	"context"
	"database/sql"

	"go-zero-user/common/errorx"
	"go-zero-user/role-rpc/internal/model"
	"go-zero-user/role-rpc/internal/svc"
	"go-zero-user/role-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建角色
func (l *CreateRoleLogic) CreateRole(in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	// 检查角色编码是否已存在
	_, err := l.svcCtx.RoleModel.FindOneByCode(l.ctx, in.Code)
	if err == nil {
		return nil, errorx.NewCodeError(errorx.CodeRoleAlreadyExists)
	}
	if err != sql.ErrNoRows {
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 创建角色
	role := &model.Role{
		Name:        in.Name,
		Code:        in.Code,
		Description: in.Description,
		Status:      1,
	}

	result, err := l.svcCtx.RoleModel.Insert(l.ctx, role)
	if err != nil {
		l.Logger.Errorf("创建角色失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	id, _ := result.LastInsertId()

	return &pb.CreateRoleResp{
		Id: id,
	}, nil
}
