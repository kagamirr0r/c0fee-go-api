package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type farmRepository struct {
	db *gorm.DB
}

func (fr *farmRepository) GetById(domainFarm *entity.Farm, id uint) error {
	var modelFarm model.Farm
	if err := fr.db.
		Preload("Area").
		Preload("Farmers").
		Where("id = ?", id).
		First(&modelFarm).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityFarm := entity_model.ModelFarmToEntity(&modelFarm)
	*domainFarm = *entityFarm
	return nil
}

func (fr *farmRepository) List(domainFarms *[]entity.Farm) error {
	var modelFarms []model.Farm
	if err := fr.db.
		Preload("Area").
		Preload("Farmers").
		Find(&modelFarms).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainFarms = entity_model.ModelFarmsToEntities(modelFarms)
	return nil
}

func NewFarmRepository(db *gorm.DB) domainRepo.IFarmRepository {
	return &farmRepository{db}
}
