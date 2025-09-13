package roaster

import "c0fee-api/common"

type IRoasterRepository interface {
	GetById(roaster *Entity, id uint) error
	List(roasters *[]Entity) error
	Search(roasters *[]Entity, params common.QueryParams) error
	Create(roaster *Entity) error
	Update(roaster *Entity) error
}
