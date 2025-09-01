package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type areaRepository struct {
	db *gorm.DB
}

func (ar *areaRepository) GetById(domainArea *entity.Area, id uint) error {
	var modelArea model.Area
	if err := ar.db.
		Preload("Country").
		Preload("Farms").
		Where("id = ?", id).
		First(&modelArea).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityArea := entity_model.ModelAreaToEntity(&modelArea)
	*domainArea = *entityArea
	return nil
}

func (ar *areaRepository) List(domainAreas *[]entity.Area) error {
	var modelAreas []model.Area
	if err := ar.db.
		Preload("Country").
		Preload("Farms").
		Find(&modelAreas).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainAreas = entity_model.ModelAreasToEntities(modelAreas)
	return nil
}

func NewAreaRepository(db *gorm.DB) domainRepo.IAreaRepository {
	return &areaRepository{db}
}
