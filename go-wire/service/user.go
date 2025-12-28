package service

import (
	"yiwen/go-wire/model"
	"yiwen/go-wire/repository"
)

// UserService 用户业务逻辑接口
type UserService interface {
	GetAllUsers() ([]*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	CreateUser(req *model.CreateUserRequest) (*model.User, error)
	UpdateUser(id int64, req *model.UpdateUserRequest) (*model.User, error)
	DeleteUser(id int64) error
}

// userService 用户业务逻辑实现
type userService struct {
	repo repository.UserRepository
}

// NewUserService 创建用户服务（这是一个 Provider）
// 注意：它依赖 UserRepository，Wire 会自动注入
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id int64, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
