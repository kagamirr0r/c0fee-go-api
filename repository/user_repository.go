package repository

import (
	"c0fee-api/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrDuplicateUser = errors.New("duplicate user")
var ErrDuplicateId = errors.New("duplicate id")
var ErrDuplicateName = errors.New("duplicate name")

// UserRepositoryのインターフェース
type IUserRepository interface {
	GetById(user *model.User, id uuid.UUID) error
	CreateUser(user *model.User) error
}

// UserRepositoryの構造体
type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) GetById(user *model.User, id uuid.UUID) error {
	if err := ur.db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	var existUser model.User

	// 重複確認, Firstでエラーがない場合は既にユーザーが存在する
	if err := ur.db.Where("id = ?", user.ID).Or("name = ?", user.Name).First(&existUser).Error; err == nil {
		if user.ID == existUser.ID {
			return ErrDuplicateId
		} else if user.Name == existUser.Name {
			return ErrDuplicateName
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UserRepositoryのコンストラクタ(ファクトリ)関数
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
