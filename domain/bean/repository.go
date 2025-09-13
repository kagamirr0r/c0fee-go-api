package bean

import (
	"c0fee-api/common"
	"c0fee-api/domain/bean_rating"

	"github.com/google/uuid"
)

type IBeanRepository interface {
	GetById(bean *Entity, id uint) error
	GetBeansByUserId(beans *[]Entity, userID uuid.UUID, params common.QueryParams) error
	SearchBeansByUserId(beans *[]Entity, userID uuid.UUID, params common.QueryParams) error
	Create(bean *Entity) error
	Update(bean *Entity) error
	SetVarieties(beanID uint, varietyIDs []uint) error
}

type IBeanRatingRepository interface {
	Create(beanRating *bean_rating.Entity) error
	GetByBeanID(beanRatings *[]bean_rating.Entity, beanID uint) error
	UpdateByID(beanRating *bean_rating.Entity) error
}
