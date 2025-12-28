package logic

import (
	"context"

	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"
	"go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserReq) (resp *types.ListUserResp, err error) {
	// 调用 user-rpc 获取用户列表
	result, err := l.svcCtx.UserRpc.ListUser(l.ctx, &pb.ListUserReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var users []types.UserInfo
	for _, u := range result.Users {
		users = append(users, types.UserInfo{
			Id:        u.Id,
			Username:  u.Username,
			Email:     u.Email,
			Phone:     u.Phone,
			Avatar:    u.Avatar,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
		})
	}

	return &types.ListUserResp{
		Users: users,
		Total: result.Total,
	}, nil
}
