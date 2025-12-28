package service

import (
	"context"

	"github.com/google/uuid"

	"yiwen/go-ddd/internal/application/command"
	"yiwen/go-ddd/internal/application/dto"
	"yiwen/go-ddd/internal/application/query"
	"yiwen/go-ddd/internal/domain/aggregate"
	"yiwen/go-ddd/internal/domain/repository"
	domainservice "yiwen/go-ddd/internal/domain/service"
	"yiwen/go-ddd/internal/domain/valueobject"
	"yiwen/go-ddd/pkg/errors"
)

// UserApplicationService 用户应用服务
// 应用服务是应用层的核心，负责：
// 1. 协调领域对象完成用例
// 2. 处理事务
// 3. 调用领域服务
// 4. 不包含业务逻辑（业务逻辑在领域层）
type UserApplicationService struct {
	userRepo          repository.UserRepository
	userDomainService *domainservice.UserDomainService
}

// NewUserApplicationService 创建用户应用服务
func NewUserApplicationService(
	userRepo repository.UserRepository,
	userDomainService *domainservice.UserDomainService,
) *UserApplicationService {
	return &UserApplicationService{
		userRepo:          userRepo,
		userDomainService: userDomainService,
	}
}

// Register 注册用户
func (s *UserApplicationService) Register(ctx context.Context, cmd *command.RegisterUserCommand) (*dto.UserDTO, error) {
	// 验证用户名唯一性
	if err := s.userDomainService.ValidateUniqueUsername(ctx, cmd.Username); err != nil {
		return nil, err
	}

	// 验证邮箱唯一性
	if err := s.userDomainService.ValidateUniqueEmail(ctx, cmd.Email); err != nil {
		return nil, err
	}

	// 创建邮箱值对象
	email, err := valueobject.NewEmail(cmd.Email)
	if err != nil {
		return nil, errors.Wrap(err, "invalid email")
	}

	// 创建密码值对象
	password, err := valueobject.NewPassword(cmd.Password)
	if err != nil {
		return nil, errors.Wrap(err, "invalid password")
	}

	// 使用聚合根创建用户
	userAggregate := aggregate.Register(uuid.New().String(), cmd.Username, email, password)
	userAggregate.User.Nickname = cmd.Nickname

	// 保存用户
	if err := s.userRepo.Save(ctx, userAggregate.User); err != nil {
		return nil, errors.Wrap(err, "failed to save user")
	}

	// TODO: 发布领域事件
	// for _, event := range userAggregate.GetUncommittedEvents() {
	//     eventPublisher.Publish(event)
	// }

	result := dto.ToUserDTO(userAggregate.User)
	return &result, nil
}

// Login 用户登录
func (s *UserApplicationService) Login(ctx context.Context, q *query.LoginQuery) (*dto.UserDTO, error) {
	user, err := s.userDomainService.ValidateUserCredentials(ctx, q.Username, q.Password)
	if err != nil {
		return nil, err
	}

	result := dto.ToUserDTO(user)
	return &result, nil
}

// GetUserByID 根据ID获取用户
func (s *UserApplicationService) GetUserByID(ctx context.Context, q *query.GetUserByIDQuery) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(ctx, q.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	result := dto.ToUserDTO(user)
	return &result, nil
}

// GetUserByUUID 根据UUID获取用户
func (s *UserApplicationService) GetUserByUUID(ctx context.Context, q *query.GetUserByUUIDQuery) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByUUID(ctx, q.UUID)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	result := dto.ToUserDTO(user)
	return &result, nil
}

// ListUsers 获取用户列表
func (s *UserApplicationService) ListUsers(ctx context.Context, q *query.ListUsersQuery) (*dto.UserListDTO, error) {
	users, total, err := s.userRepo.List(ctx, q.Offset, q.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	return &dto.UserListDTO{
		Total: total,
		Items: dto.ToUserDTOList(users),
	}, nil
}

// UpdateProfile 更新用户资料
func (s *UserApplicationService) UpdateProfile(ctx context.Context, cmd *command.UpdateProfileCommand) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(ctx, cmd.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	// 使用聚合根更新资料
	userAggregate := aggregate.NewUserAggregate(user)
	userAggregate.UpdateProfile(cmd.Nickname, cmd.Avatar)

	if err := s.userRepo.Save(ctx, userAggregate.User); err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	result := dto.ToUserDTO(userAggregate.User)
	return &result, nil
}

// ChangePassword 修改密码
func (s *UserApplicationService) ChangePassword(ctx context.Context, cmd *command.ChangePasswordCommand) error {
	user, err := s.userRepo.FindByID(ctx, cmd.UserID)
	if err != nil {
		return errors.Wrap(err, "user not found")
	}

	// 验证旧密码
	if err := user.Password.Verify(cmd.OldPassword); err != nil {
		return domainservice.ErrInvalidCredentials
	}

	// 创建新密码
	newPassword, err := valueobject.NewPassword(cmd.NewPassword)
	if err != nil {
		return errors.Wrap(err, "invalid new password")
	}

	// 使用聚合根修改密码
	userAggregate := aggregate.NewUserAggregate(user)
	userAggregate.ChangePassword(newPassword)

	if err := s.userRepo.Save(ctx, userAggregate.User); err != nil {
		return errors.Wrap(err, "failed to update password")
	}

	return nil
}

// DeleteUser 删除用户
func (s *UserApplicationService) DeleteUser(ctx context.Context, cmd *command.DeleteUserCommand) error {
	if err := s.userRepo.Delete(ctx, cmd.UserID); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}
