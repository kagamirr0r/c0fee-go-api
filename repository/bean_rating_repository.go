package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type beanRatingRepository struct {
	db *gorm.DB
}

func (brr *beanRatingRepository) Create(domainBeanRating *entity.BeanRating) error {
	modelBeanRating := entity_model.EntityBeanRatingToModel(domainBeanRating)
	if err := brr.db.Create(modelBeanRating).Error; err != nil {
		return err
	}

	// Update domain entity with database-generated fields
	domainBeanRating.ID = modelBeanRating.ID
	domainBeanRating.CreatedAt = modelBeanRating.CreatedAt
	domainBeanRating.UpdatedAt = modelBeanRating.UpdatedAt

	return nil
}

func (brr *beanRatingRepository) GetByBeanID(domainBeanRatings *[]entity.BeanRating, beanID uint) error {
	var modelBeanRatings []model.BeanRating
	if err := brr.db.
		Preload("User").
		Where("bean_id = ?", beanID).
		Find(&modelBeanRatings).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainBeanRatings = entity_model.ModelBeanRatingsToEntities(modelBeanRatings)
	return nil
}

func (brr *beanRatingRepository) UpdateByID(domainBeanRating *entity.BeanRating) error {
	modelBeanRating := entity_model.EntityBeanRatingToModel(domainBeanRating)
	if err := brr.db.Save(modelBeanRating).Error; err != nil {
		return err
	}

	// Update domain entity with database fields
	domainBeanRating.UpdatedAt = modelBeanRating.UpdatedAt

	return nil
}

func NewBeanRatingRepository(db *gorm.DB) domainRepo.IBeanRatingRepository {
	return &beanRatingRepository{db}
}
