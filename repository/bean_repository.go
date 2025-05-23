package repository

import (
	"c0fee-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBeanRepository interface {
	GetBeanById(bean *model.Bean, id uint) error
	GetBeansByUserId(beans []model.Bean, userID uuid.UUID) error
}

type beanRepository struct {
	db *gorm.DB
}

func (br *beanRepository) GetBeanById(bean *model.Bean, id uint) error {
	if err := br.db.
		Preload("User").
		Preload("Roaster").
		Preload("Country").
		Preload("ProcessMethod").
		Preload("Varieties").
		Preload("Area").
		Preload("Farm").
		Preload("Farmer").
		Preload("BeanRatings").
		Preload("BeanRatings.User").
		Where("id = ?", id).
		First(bean).Error; err != nil {
		return err
	}
	return nil
}

func (br *beanRepository) GetBeansByUserId(beans []model.Bean, userID uuid.UUID) error {
	if err := br.db.
		Preload("User").
		Preload("Roaster").
		Preload("Country").
		Preload("ProcessMethod").
		Preload("Varieties").
		Preload("Area").
		Preload("Farm").
		Preload("Farmer").
		Where("user_id = ?", userID).
		Find(&beans).Error; err != nil {
		return err
	}
	return nil
}

func NewBeanRepository(db *gorm.DB) IBeanRepository {
	return &beanRepository{db}
}
