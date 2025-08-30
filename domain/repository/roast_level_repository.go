package repository

import "c0fee-api/domain/entity"

type IRoastLevelRepository interface {
	GetAll(roastLevels *[]entity.RoastLevel) error
	GetById(roastLevel *entity.RoastLevel, id uint) error
}