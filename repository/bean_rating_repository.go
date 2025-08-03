package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IBeanRatingRepository interface {
	Create(beanRating *model.BeanRating) error
	GetByBeanID(beanRatings *[]model.BeanRating, beanID uint) error
}

type beanRatingRepository struct {
	db *gorm.DB
}

func (brr *beanRatingRepository) Create(beanRating *model.BeanRating) error {
	if err := brr.db.Create(beanRating).Error; err != nil {
		return err
	}
	return nil
}

func (brr *beanRatingRepository) GetByBeanID(beanRatings *[]model.BeanRating, beanID uint) error {
	if err := brr.db.
		Preload("User").
		Where("bean_id = ?", beanID).
		Find(beanRatings).Error; err != nil {
		return err
	}
	return nil
}

func NewBeanRatingRepository(db *gorm.DB) IBeanRatingRepository {
	return &beanRatingRepository{db}
}
