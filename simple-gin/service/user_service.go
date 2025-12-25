package service

import (
	"context"
	"errors"
	"example/simple-gin/models"
	"log"
)

// UserService 用户服务接口定义
type UserService interface {
	// GetUsers 获取所有用户
	GetUsers(ctx context.Context) ([]*models.User, error)
	// GetUserByID 根据ID获取用户
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	// CreateUser 创建用户
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
	// UpdateUser 更新用户
	UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error)
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
func (s *userService) GetUsers(ctx context.Context) ([]*models.User, error) {
	// 监听context取消信号
	select {
	case <-ctx.Done():
		log.Printf("GetUsers request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	log.Println("Service: fetching all users")
	users := s.db.GetAllUsers()

	if users == nil {
		users = make([]*models.User, 0)
	}

	return users, nil
}

// GetUserByID 实现根据ID获取用户
func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	select {
	case <-ctx.Done():
		log.Printf("GetUserByID request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	log.Printf("Service: fetching user by id: %d", id)
	user := s.db.GetUser(id)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateUser 实现创建用户
func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	select {
	case <-ctx.Done():
		log.Printf("CreateUser request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if req == nil {
		return nil, errors.New("invalid request")
	}

	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Phone == "" {
		return nil, errors.New("phone is required")
	}

	log.Printf("Service: creating user with email: %s", req.Email)
	user := s.db.CreateUser(req)

	return user, nil
}

// UpdateUser 实现更新用户
func (s *userService) UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error) {
	select {
	case <-ctx.Done():
		log.Printf("UpdateUser request cancelled: %v", ctx.Err())
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

	log.Printf("Service: updating user with id: %d", id)
	user := s.db.UpdateUser(id, req)

	return user, nil
}

// DeleteUser 实现删除用户
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		log.Printf("DeleteUser request cancelled: %v", ctx.Err())
		return ctx.Err()
	default:
	}

	if id <= 0 {
		return errors.New("invalid user id")
	}

	log.Printf("Service: deleting user with id: %d", id)
	if !s.db.DeleteUser(id) {
		return errors.New("user not found")
	}

	return nil
}
