package repository

import (
	"c0fee-api/common"
	"c0fee-api/domain/entity"
)

type IRoasterRepository interface {
	List(roasters *[]entity.Roaster) error
	Search(roasters *[]entity.Roaster, params common.QueryParams) error
	GetById(roaster *entity.Roaster, id uint) error
	Create(roaster *entity.Roaster) error
	Update(roaster *entity.Roaster) error
}
