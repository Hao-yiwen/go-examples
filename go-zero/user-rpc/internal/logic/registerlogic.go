package logic

import (
	"context"
	"database/sql"

	"github.com/haoyiwen/go-examples/go-zero/common/errorx"
	"github.com/haoyiwen/go-examples/go-zero/common/utils"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/internal/model"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 检查用户是否已存在
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, errorx.NewCodeError(errorx.CodeUserAlreadyExists)
	}
	if err != sql.ErrNoRows {
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(in.Password)
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	// 创建用户
	user := &model.User{
		Username: in.Username,
		Password: hashedPassword,
		Email:    in.Email,
		Phone:    in.Phone,
		Status:   1,
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("创建用户失败: %v", err)
		return nil, errorx.NewCodeError(errorx.CodeInternal)
	}

	id, _ := result.LastInsertId()

	return &pb.RegisterResp{
		Id:       id,
		Username: in.Username,
	}, nil
}
