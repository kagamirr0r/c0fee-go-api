package repository

import (
	"c0fee-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepositoryのインターフェース
type IUserRepository interface {
	GetUserById(user *model.User, id uuid.UUID) error
	CreateUser(user *model.User) error
}

// UserRepositoryの構造体
type UserRepository struct {
	db *gorm.DB
}

// UserRepositoryのコンストラクタ(ファクトリ)関数
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserById(user *model.User, id uuid.UUID) error {
	if err := ur.db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
