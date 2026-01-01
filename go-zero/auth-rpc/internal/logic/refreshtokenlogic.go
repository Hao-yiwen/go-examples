package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/haoyiwen/go-examples/go-zero/auth-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/auth-rpc/pb"
	"github.com/haoyiwen/go-examples/go-zero/common/errorx"
	"github.com/haoyiwen/go-examples/go-zero/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 刷新Token
func (l *RefreshTokenLogic) RefreshToken(in *pb.RefreshTokenReq) (*pb.RefreshTokenResp, error) {
	// 解析Refresh Token
	claims, err := utils.ParseToken(in.RefreshToken, l.svcCtx.Config.Jwt.RefreshSecret)
	if err != nil {
		l.Logger.Infof("RefreshToken解析失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeTokenInvalid)
	}

	// 生成新的 Access Token
	accessToken, err := utils.GenerateToken(
		claims.UserId,
		claims.Username,
		l.svcCtx.Config.Jwt.AccessSecret,
		l.svcCtx.Config.Jwt.AccessExpire,
	)
	if err != nil {
		l.Logger.Errorf("生成AccessToken失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeTokenGenerate)
	}

	// 生成新的 Refresh Token
	refreshToken, err := utils.GenerateToken(
		claims.UserId,
		claims.Username,
		l.svcCtx.Config.Jwt.RefreshSecret,
		l.svcCtx.Config.Jwt.RefreshExpire,
	)
	if err != nil {
		l.Logger.Errorf("生成RefreshToken失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeTokenGenerate)
	}

	// 更新Redis中的token
	tokenKey := fmt.Sprintf("token:%d", claims.UserId)
	err = l.svcCtx.Redis.SetexCtx(l.ctx, tokenKey, accessToken, int(l.svcCtx.Config.Jwt.AccessExpire))
	if err != nil {
		l.Logger.Errorf("更新Token到Redis失败: %v", err)
	}

	expiresAt := time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.AccessExpire) * time.Second).Unix()

	return &pb.RefreshTokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
