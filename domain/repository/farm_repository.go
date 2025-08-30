package repository

import "c0fee-api/domain/entity"

type IFarmRepository interface {
	GetById(farm *entity.Farm, id uint) error
	List(farms *[]entity.Farm) error
}