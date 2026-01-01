package logic

import (
	"context"
	"strings"

	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/pb"
	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.LogoutResp, err error) {
	// 从 context 中获取 token
	authHeader, ok := l.ctx.Value("Authorization").(string)
	if !ok {
		return &types.LogoutResp{Success: true}, nil
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// 调用 auth-rpc 注销 Token
	_, err = l.svcCtx.AuthRpc.RevokeToken(l.ctx, &pb.RevokeTokenReq{
		Token: token,
	})
	if err != nil {
		l.Logger.Errorf("注销Token失败: %v", err)
	}

	return &types.LogoutResp{
		Success: true,
	}, nil
}
