package service

import (
	"context"
	"errors"

	"yiwen/go-ddd/internal/domain/entity"
	"yiwen/go-ddd/internal/domain/repository"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserNotActive         = errors.New("user is not active")
	ErrInvalidCredentials    = errors.New("invalid credentials")
)

// UserDomainService 用户领域服务
// 领域服务用于处理不属于单个实体的业务逻辑
// 例如：涉及多个实体的操作、需要访问仓储的验证逻辑等
type UserDomainService struct {
	userRepo repository.UserRepository
}

// NewUserDomainService 创建用户领域服务
func NewUserDomainService(userRepo repository.UserRepository) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
	}
}

// ValidateUniqueUsername 验证用户名唯一性
func (s *UserDomainService) ValidateUniqueUsername(ctx context.Context, username string) error {
	exists, err := s.userRepo.ExistsByUsername(ctx, username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUsernameAlreadyExists
	}
	return nil
}

// ValidateUniqueEmail 验证邮箱唯一性
func (s *UserDomainService) ValidateUniqueEmail(ctx context.Context, email string) error {
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailAlreadyExists
	}
	return nil
}

// ValidateUserCredentials 验证用户凭证（登录）
func (s *UserDomainService) ValidateUserCredentials(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := user.Password.Verify(password); err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.IsActive() {
		return nil, ErrUserNotActive
	}

	return user, nil
}

// CanUserPerformAction 检查用户是否可以执行操作
func (s *UserDomainService) CanUserPerformAction(user *entity.User, action string) bool {
	// 管理员可以执行所有操作
	if user.IsAdmin() {
		return true
	}

	// 普通用户只能执行部分操作
	allowedActions := map[string]bool{
		"view_profile":   true,
		"update_profile": true,
		"change_password": true,
	}

	return allowedActions[action]
}

// TransferAdmin 转移管理员权限
func (s *UserDomainService) TransferAdmin(ctx context.Context, fromUser, toUser *entity.User) error {
	if !fromUser.IsAdmin() {
		return errors.New("source user is not an admin")
	}

	// 这里可以添加更多的业务规则
	// 例如：检查目标用户是否满足成为管理员的条件

	return nil
}
