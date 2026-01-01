package logic

import (
	"context"
	"fmt"

	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/pb"
	"github.com/Hao-yiwen/go-examples/go-zero/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevokeTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevokeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevokeTokenLogic {
	return &RevokeTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销Token
func (l *RevokeTokenLogic) RevokeToken(in *pb.RevokeTokenReq) (*pb.RevokeTokenResp, error) {
	// 解析Token获取用户信息
	claims, err := utils.ParseToken(in.Token, l.svcCtx.Config.Jwt.AccessSecret)
	if err != nil {
		// Token无效也算注销成功
		return &pb.RevokeTokenResp{
			Success: true,
		}, nil
	}

	// 将Token加入黑名单
	blacklistKey := fmt.Sprintf("token:blacklist:%s", in.Token)
	err = l.svcCtx.Redis.SetexCtx(l.ctx, blacklistKey, "1", int(l.svcCtx.Config.Jwt.AccessExpire))
	if err != nil {
		l.Logger.Errorf("将Token加入黑名单失败: %v", err)
	}

	// 删除用户的token记录
	tokenKey := fmt.Sprintf("token:%d", claims.UserId)
	_, _ = l.svcCtx.Redis.DelCtx(l.ctx, tokenKey)

	return &pb.RevokeTokenResp{
		Success: true,
	}, nil
}
