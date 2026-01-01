package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/pb"
	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成Token
func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// 生成 Access Token
	accessToken, err := utils.GenerateToken(
		in.UserId,
		in.Username,
		l.svcCtx.Config.Jwt.AccessSecret,
		l.svcCtx.Config.Jwt.AccessExpire,
	)
	if err != nil {
		l.Logger.Errorf("生成AccessToken失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeTokenGenerate)
	}

	// 生成 Refresh Token
	refreshToken, err := utils.GenerateToken(
		in.UserId,
		in.Username,
		l.svcCtx.Config.Jwt.RefreshSecret,
		l.svcCtx.Config.Jwt.RefreshExpire,
	)
	if err != nil {
		l.Logger.Errorf("生成RefreshToken失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeTokenGenerate)
	}

	// 将token存入Redis（用于注销功能）
	tokenKey := fmt.Sprintf("token:%d", in.UserId)
	err = l.svcCtx.Redis.SetexCtx(l.ctx, tokenKey, accessToken, int(l.svcCtx.Config.Jwt.AccessExpire))
	if err != nil {
		l.Logger.Errorf("存储Token到Redis失败: %v", err)
	}

	expiresAt := time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.AccessExpire) * time.Second).Unix()

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
