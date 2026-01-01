package logic

import (
	"context"
	"encoding/json"

	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/types"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	// 从 context 中获取用户ID（JWT中间件已解析）
	userId, _ := l.ctx.Value("userId").(json.Number)
	userIdInt, _ := userId.Int64()

	// 调用 user-rpc 获取用户信息
	result, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: userIdInt,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		User: types.UserInfo{
			Id:        result.User.Id,
			Username:  result.User.Username,
			Email:     result.User.Email,
			Phone:     result.User.Phone,
			Avatar:    result.User.Avatar,
			Status:    result.User.Status,
			CreatedAt: result.User.CreatedAt,
		},
	}, nil
}
