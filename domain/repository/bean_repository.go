package repository

import (
	"c0fee-api/common"
	"c0fee-api/domain/entity"

	"github.com/google/uuid"
)

type IBeanRepository interface {
	GetById(bean *entity.Bean, id uint) error
	GetBeansByUserId(beans *[]entity.Bean, userID uuid.UUID, params common.QueryParams) error
	SearchBeansByUserId(beans *[]entity.Bean, userID uuid.UUID, params common.QueryParams) error
	Create(bean *entity.Bean) error
	Update(bean *entity.Bean) error
	SetVarieties(beanID uint, varietyIDs []uint) error
}