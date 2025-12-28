package logic

import (
	"context"
	"encoding/json"

	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"
	"go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	// 从 context 中获取用户ID
	userId, _ := l.ctx.Value("userId").(json.Number)
	userIdInt, _ := userId.Int64()

	// 调用 user-rpc 更新用户信息
	result, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &pb.UpdateUserReq{
		Id:     userIdInt,
		Email:  req.Email,
		Phone:  req.Phone,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateUserResp{
		Success: result.Success,
	}, nil
}
