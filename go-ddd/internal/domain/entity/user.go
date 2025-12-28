package entity

import (
	"time"

	"yiwen/go-ddd/internal/domain/valueobject"
)

// UserStatus 用户状态
type UserStatus int

const (
	UserStatusActive   UserStatus = 1 // 激活
	UserStatusInactive UserStatus = 2 // 未激活
	UserStatusBanned   UserStatus = 3 // 禁用
)

// UserRole 用户角色
type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

// User 用户实体
// 实体是DDD中的核心概念，具有唯一标识（ID）
// 实体的相等性由ID决定，而不是属性
type User struct {
	ID        uint64                  // 数据库自增ID
	UUID      string                  // 业务唯一标识
	Username  string                  // 用户名
	Email     valueobject.Email       // 邮箱（值对象）
	Password  valueobject.Password    // 密码（值对象）
	Nickname  string                  // 昵称
	Avatar    string                  // 头像URL
	Status    UserStatus              // 状态
	Role      UserRole                // 角色
	CreatedAt time.Time               // 创建时间
	UpdatedAt time.Time               // 更新时间
}

// NewUser 创建新用户
func NewUser(uuid, username string, email valueobject.Email, password valueobject.Password) *User {
	now := time.Now()
	return &User{
		UUID:      uuid,
		Username:  username,
		Email:     email,
		Password:  password,
		Status:    UserStatusActive,
		Role:      UserRoleUser,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// IsActive 检查用户是否激活
func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

// IsAdmin 检查用户是否是管理员
func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

// UpdateProfile 更新用户资料
func (u *User) UpdateProfile(nickname, avatar string) {
	u.Nickname = nickname
	u.Avatar = avatar
	u.UpdatedAt = time.Now()
}

// ChangePassword 修改密码
func (u *User) ChangePassword(newPassword valueobject.Password) {
	u.Password = newPassword
	u.UpdatedAt = time.Now()
}

// Activate 激活用户
func (u *User) Activate() {
	u.Status = UserStatusActive
	u.UpdatedAt = time.Now()
}

// Deactivate 停用用户
func (u *User) Deactivate() {
	u.Status = UserStatusInactive
	u.UpdatedAt = time.Now()
}

// Ban 禁用用户
func (u *User) Ban() {
	u.Status = UserStatusBanned
	u.UpdatedAt = time.Now()
}

// PromoteToAdmin 提升为管理员
func (u *User) PromoteToAdmin() {
	u.Role = UserRoleAdmin
	u.UpdatedAt = time.Now()
}
