package repository

import "c0fee-api/domain/entity"

type IBeanRatingRepository interface {
	Create(beanRating *entity.BeanRating) error
	GetByBeanID(beanRatings *[]entity.BeanRating, beanID uint) error
	UpdateByID(beanRating *entity.BeanRating) error
}