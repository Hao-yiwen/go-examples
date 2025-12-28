package logic

import (
	"context"
	"database/sql"

	"go-zero-user/common/errorx"
	"go-zero-user/user-rpc/internal/model"
	"go-zero-user/user-rpc/internal/svc"
	"go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	// 检查用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(errorx.CodeUserNotFound)
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 更新用户信息
	err = l.svcCtx.UserModel.Update(l.ctx, &model.User{
		Id:     in.Id,
		Email:  in.Email,
		Phone:  in.Phone,
		Avatar: in.Avatar,
	})
	if err != nil {
		l.Logger.Errorf("更新用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.UpdateUserResp{
		Success: true,
	}, nil
}
