package model

import (
	"time"

	"gorm.io/gorm"

	"yiwen/go-ddd/internal/domain/entity"
	"yiwen/go-ddd/internal/domain/valueobject"
)

// UserModel 用户数据库模型
// 数据库模型与领域实体分离的好处：
// 1. 领域实体不受数据库结构影响
// 2. 可以自由添加数据库特有的字段（如软删除）
// 3. 便于处理ORM特有的标签和钩子
type UserModel struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement"`
	UUID         string         `gorm:"type:varchar(36);uniqueIndex;not null"`
	Username     string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email        string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	Nickname     string         `gorm:"type:varchar(50)"`
	Avatar       string         `gorm:"type:varchar(255)"`
	Status       int            `gorm:"type:tinyint;not null;default:1"`
	Role         string         `gorm:"type:varchar(20);not null;default:user"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (UserModel) TableName() string {
	return "users"
}

// ToEntity 将数据库模型转换为领域实体
func (m *UserModel) ToEntity() *entity.User {
	email, _ := valueobject.NewEmail(m.Email)
	password := valueobject.NewPasswordFromHash(m.PasswordHash)

	return &entity.User{
		ID:        m.ID,
		UUID:      m.UUID,
		Username:  m.Username,
		Email:     email,
		Password:  password,
		Nickname:  m.Nickname,
		Avatar:    m.Avatar,
		Status:    entity.UserStatus(m.Status),
		Role:      entity.UserRole(m.Role),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromEntity 从领域实体创建数据库模型
func FromEntity(user *entity.User) *UserModel {
	return &UserModel{
		ID:           user.ID,
		UUID:         user.UUID,
		Username:     user.Username,
		Email:        user.Email.String(),
		PasswordHash: user.Password.Hash(),
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Status:       int(user.Status),
		Role:         string(user.Role),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
