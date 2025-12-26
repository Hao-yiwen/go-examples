package service

import (
	"context"
	"errors"
	"log/slog"

	"example/simple-gin/internal/model"
	"example/simple-gin/pkg/validator"
)

// UserService 用户服务接口定义
type UserService interface {
	// GetUsers 获取所有用户
	GetUsers(ctx context.Context) ([]*model.User, error)
	// GetUserByID 根据ID获取用户
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	// UpdateUser 更新用户
	UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) (*model.User, error)
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id int) error
}

// userService 用户服务实现
type userService struct {
	db Database
}

// NewUserService 创建用户服务实例
func NewUserService(db Database) UserService {
	return &userService{
		db: db,
	}
}

// GetUsers 实现获取所有用户
func (s *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	select {
	case <-ctx.Done():
		slog.Warn("GetUsers request cancelled", "error", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	slog.Debug("fetching all users")
	users := s.db.GetAllUsers()

	if users == nil {
		users = make([]*model.User, 0)
	}

	return users, nil
}

// GetUserByID 实现根据ID获取用户
func (s *userService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	select {
	case <-ctx.Done():
		slog.Warn("GetUserByID request cancelled", "error", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	slog.Debug("fetching user by id", "id", id)
	user := s.db.GetUser(id)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateUser 实现创建用户
func (s *userService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	select {
	case <-ctx.Done():
		slog.Warn("CreateUser request cancelled", "error", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if req == nil {
		return nil, errors.New("invalid request")
	}

	// 使用 pkg/validator 进行验证
	if !validator.IsNotEmpty(req.Name) {
		return nil, errors.New("name is required")
	}

	if !validator.IsValidEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	if !validator.IsValidPhone(req.Phone) {
		return nil, errors.New("invalid phone format")
	}

	slog.Info("creating user", "email", req.Email)
	user := s.db.CreateUser(req)

	return user, nil
}

// UpdateUser 实现更新用户
func (s *userService) UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) (*model.User, error) {
	select {
	case <-ctx.Done():
		slog.Warn("UpdateUser request cancelled", "error", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	// 检查用户是否存在
	existingUser := s.db.GetUser(id)
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	// 使用 pkg/validator 验证更新字段
	if req.Email != "" && !validator.IsValidEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	if req.Phone != "" && !validator.IsValidPhone(req.Phone) {
		return nil, errors.New("invalid phone format")
	}

	slog.Info("updating user", "id", id)
	user := s.db.UpdateUser(id, req)

	return user, nil
}

// DeleteUser 实现删除用户
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		slog.Warn("DeleteUser request cancelled", "error", ctx.Err())
		return ctx.Err()
	default:
	}

	if id <= 0 {
		return errors.New("invalid user id")
	}

	slog.Info("deleting user", "id", id)
	if !s.db.DeleteUser(id) {
		return errors.New("user not found")
	}

	return nil
}
