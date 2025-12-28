package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"yiwen/go-ddd/internal/domain/entity"
	"yiwen/go-ddd/internal/domain/repository"
	"yiwen/go-ddd/internal/infrastructure/persistence/model"
)

// UserRepository MySQL用户仓储实现
// 这是仓储接口的具体实现
// 基础设施层实现领域层定义的接口
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

// Save 保存用户（创建或更新）
func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	userModel := model.FromEntity(user)

	if user.ID == 0 {
		// 创建
		if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
			return err
		}
		user.ID = userModel.ID
	} else {
		// 更新
		if err := r.db.WithContext(ctx).Save(userModel).Error; err != nil {
			return err
		}
	}

	return nil
}

// FindByID 根据ID查找用户
func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel.ToEntity(), nil
}

// FindByUUID 根据UUID查找用户
func (r *UserRepository) FindByUUID(ctx context.Context, uuid string) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel.ToEntity(), nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel.ToEntity(), nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var userModel model.UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel.ToEntity(), nil
}

// Delete 删除用户（软删除）
func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&model.UserModel{}, id).Error
}

// List 分页查询用户列表
func (r *UserRepository) List(ctx context.Context, offset, limit int) ([]*entity.User, int64, error) {
	var userModels []model.UserModel
	var total int64

	// 查询总数
	if err := r.db.WithContext(ctx).Model(&model.UserModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("id DESC").
		Find(&userModels).Error; err != nil {
		return nil, 0, err
	}

	// 转换为实体
	users := make([]*entity.User, len(userModels))
	for i := range userModels {
		users[i] = userModels[i].ToEntity()
	}

	return users, total, nil
}

// ExistsByUsername 检查用户名是否存在
func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&model.UserModel{}).
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByEmail 检查邮箱是否存在
func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&model.UserModel{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
