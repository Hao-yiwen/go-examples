package repository

import (
	"errors"
	"sync"

	"yiwen/go-wire/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id int64) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int64) error
}

// userRepository 用户数据访问实现（内存存储）
type userRepository struct {
	mu     sync.RWMutex
	users  map[int64]*model.User
	nextID int64
}

// NewUserRepository 创建用户仓库（这是一个 Provider）
func NewUserRepository() UserRepository {
	return &userRepository{
		users:  make(map[int64]*model.User),
		nextID: 1,
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *userRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return ErrUserNotFound
	}
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}
