package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IProcessMethodRepository interface {
	List(processMethods *[]model.ProcessMethod) error
}

type processMethodRepository struct {
	db *gorm.DB
}

func (pmr *processMethodRepository) List(processMethods *[]model.ProcessMethod) error {
	if err := pmr.db.Find(processMethods).Error; err != nil {
		return err
	}
	return nil
}

func NewProcessMethodRepository(db *gorm.DB) IProcessMethodRepository {
	return &processMethodRepository{db}
}
