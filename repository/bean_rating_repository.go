package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/bean_rating"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type beanRatingRepository struct {
	db *gorm.DB
}

func (brr *beanRatingRepository) Create(domainBeanRating *bean_rating.Entity) error {
	modelBeanRating := entity_model.BeanRatingEntityToModel(domainBeanRating)
	if err := brr.db.Create(modelBeanRating).Error; err != nil {
		return err
	}

	// Update domain entity with database-generated fields
	domainBeanRating.ID = modelBeanRating.ID
	domainBeanRating.CreatedAt = modelBeanRating.CreatedAt
	domainBeanRating.UpdatedAt = modelBeanRating.UpdatedAt

	return nil
}

func (brr *beanRatingRepository) GetByBeanID(domainBeanRatings *[]bean_rating.Entity, beanID uint) error {
	var modelBeanRatings []model.BeanRating
	if err := brr.db.
		Preload("User").
		Where("bean_id = ?", beanID).
		Find(&modelBeanRatings).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainBeanRatings = entity_model.ModelsToBeanRatingEntities(modelBeanRatings)
	return nil
}

func (brr *beanRatingRepository) UpdateByID(domainBeanRating *bean_rating.Entity) error {
	modelBeanRating := entity_model.BeanRatingEntityToModel(domainBeanRating)
	if err := brr.db.Save(modelBeanRating).Error; err != nil {
		return err
	}

	// Update domain entity with database fields
	domainBeanRating.UpdatedAt = modelBeanRating.UpdatedAt

	return nil
}

func NewBeanRatingRepository(db *gorm.DB) bean_rating.IBeanRatingRepository {
	return &beanRatingRepository{db}
}
