package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrDuplicateUser = errors.New("duplicate user")
var ErrDuplicateId = errors.New("duplicate id")
var ErrDuplicateName = errors.New("duplicate name")


// UserRepositoryの構造体
type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) GetById(user *entity.User, id uuid.UUID) error {
	var modelUser model.User
	if err := ur.db.Where("id = ?", id).First(&modelUser).Error; err != nil {
		return err
	}
	
	// Use converter to convert model.User to entity.User
	entityUser := entity_model.ModelUserToEntity(&modelUser)
	*user = *entityUser
	return nil
}

func (ur *userRepository) CreateUser(user *entity.User) error {
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

	// Use converter to convert entity.User to model.User for database operations
	modelUser := entity_model.EntityUserToModel(user)

	if err := ur.db.Create(modelUser).Error; err != nil {
		return err
	}
	
	// Update entity with database-generated fields
	user.CreatedAt = modelUser.CreatedAt
	user.UpdatedAt = modelUser.UpdatedAt
	
	return nil
}

// UserRepositoryのコンストラクタ(ファクトリ)関数
func NewUserRepository(db *gorm.DB) domainRepo.IUserRepository {
	return &userRepository{db}
}
