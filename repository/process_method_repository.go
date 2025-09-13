package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/process_method"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type processMethodRepository struct {
	db *gorm.DB
}

func (pmr *processMethodRepository) List(domainProcessMethods *[]process_method.Entity) error {
	var modelProcessMethods []model.ProcessMethod
	if err := pmr.db.Find(&modelProcessMethods).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainProcessMethods = entity_model.ModelsToProcessMethodEntities(modelProcessMethods)
	return nil
}

func NewProcessMethodRepository(db *gorm.DB) process_method.IProcessMethodRepository {
	return &processMethodRepository{db}
}
