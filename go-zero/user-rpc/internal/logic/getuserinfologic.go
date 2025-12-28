package logic

import (
	"context"
	"database/sql"

	"go-zero-user/common/errorx"
	"go-zero-user/user-rpc/internal/svc"
	"go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(errorx.CodeUserNotFound)
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	return &pb.GetUserInfoResp{
		User: &pb.UserInfo{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		},
	}, nil
}
