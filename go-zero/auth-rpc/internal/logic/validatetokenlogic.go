package logic

import (
	"context"
	"fmt"

	"github.com/haoyiwen/go-examples/go-zero/auth-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/auth-rpc/pb"
	"github.com/haoyiwen/go-examples/go-zero/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证Token
func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	// 解析Token
	claims, err := utils.ParseToken(in.Token, l.svcCtx.Config.Jwt.AccessSecret)
	if err != nil {
		l.Logger.Infof("Token解析失败: %v", err)
		return &pb.ValidateTokenResp{
			Valid:    false,
			UserId:   0,
			Username: "",
		}, nil
	}

	// 检查Token是否在黑名单中（已注销）
	blacklistKey := fmt.Sprintf("token:blacklist:%s", in.Token)
	exists, _ := l.svcCtx.Redis.ExistsCtx(l.ctx, blacklistKey)
	if exists {
		return &pb.ValidateTokenResp{
			Valid:    false,
			UserId:   0,
			Username: "",
		}, nil
	}

	return &pb.ValidateTokenResp{
		Valid:    true,
		UserId:   claims.UserId,
		Username: claims.Username,
	}, nil
}
