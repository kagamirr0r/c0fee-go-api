package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/area"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type areaRepository struct {
	db *gorm.DB
}

func (ar *areaRepository) GetById(domainArea *area.Entity, id uint) error {
	var modelArea model.Area
	if err := ar.db.
		Preload("Country").
		Preload("Farms").
		Where("id = ?", id).
		First(&modelArea).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	areaEntity := entity_model.ModelToAreaEntity(&modelArea)
	*domainArea = *areaEntity
	return nil
}

func (ar *areaRepository) List(domainAreas *[]area.Entity) error {
	var modelAreas []model.Area
	if err := ar.db.
		Preload("Country").
		Preload("Farms").
		Find(&modelAreas).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainAreas = entity_model.ModelsToAreaEntities(modelAreas)
	return nil
}

func NewAreaRepository(db *gorm.DB) area.IAreaRepository {
	return &areaRepository{db}
}
