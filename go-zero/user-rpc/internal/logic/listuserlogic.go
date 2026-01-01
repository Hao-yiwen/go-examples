package logic

import (
	"context"

	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户列表
func (l *ListUserLogic) ListUser(in *pb.ListUserReq) (*pb.ListUserResp, error) {
	page := in.Page
	pageSize := in.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	users, err := l.svcCtx.UserModel.FindList(l.ctx, page, pageSize)
	if err != nil {
		l.Logger.Errorf("查询用户列表失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	total, err := l.svcCtx.UserModel.Count(l.ctx)
	if err != nil {
		l.Logger.Errorf("查询用户总数失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	var userInfos []*pb.UserInfo
	for _, user := range users {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		})
	}

	return &pb.ListUserResp{
		Users: userInfos,
		Total: total,
	}, nil
}
