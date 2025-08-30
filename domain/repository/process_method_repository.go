package repository

import "c0fee-api/domain/entity"

type IProcessMethodRepository interface {
	List(processMethods *[]entity.ProcessMethod) error
}