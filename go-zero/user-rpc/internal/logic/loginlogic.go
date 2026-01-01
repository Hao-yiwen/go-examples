package logic

import (
	"context"
	"database/sql"

	"github.com/Hao-yiwen/go-examples/go-zero/common/errorx"
	"github.com/Hao-yiwen/go-examples/go-zero/common/utils"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/internal/svc"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 查询用户
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(errorx.CodeUserNotFound)
		}
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, errorx.NewCodeError(errorx.CodeUserDisabled)
	}

	// 验证密码
	if !utils.CheckPassword(in.Password, user.Password) {
		return nil, errorx.NewCodeError(errorx.CodePasswordError)
	}

	return &pb.LoginResp{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}
