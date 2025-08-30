package repository

import "c0fee-api/domain/entity"

type IVarietyRepository interface {
	List(varieties *[]entity.Variety) error
}