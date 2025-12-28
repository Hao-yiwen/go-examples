package logic

import (
	"context"

	"go-zero-user/auth-rpc/pb"
	"go-zero-user/user-api/internal/svc"
	"go-zero-user/user-api/internal/types"
	userpb "go-zero-user/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. 调用 user-rpc 登录验证
	userResult, err := l.svcCtx.UserRpc.Login(l.ctx, &userpb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 2. 调用 auth-rpc 生成 Token
	tokenResult, err := l.svcCtx.AuthRpc.GenerateToken(l.ctx, &pb.GenerateTokenReq{
		UserId:   userResult.Id,
		Username: userResult.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:           userResult.Id,
		Username:     userResult.Username,
		AccessToken:  tokenResult.AccessToken,
		RefreshToken: tokenResult.RefreshToken,
		ExpiresAt:    tokenResult.ExpiresAt,
	}, nil
}
