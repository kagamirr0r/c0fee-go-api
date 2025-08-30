package repository

import "c0fee-api/domain/entity"

type IAreaRepository interface {
	GetById(area *entity.Area, id uint) error
	List(areas *[]entity.Area) error
}