package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/roast_level"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type roastLevelRepository struct {
	db *gorm.DB
}

func (rlr *roastLevelRepository) GetAll(domainRoastLevels *[]roast_level.Entity) error {
	var modelRoastLevels []model.RoastLevel
	if err := rlr.db.Order("level ASC").Find(&modelRoastLevels).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainRoastLevels = entity_model.ModelsToRoastLevelEntities(modelRoastLevels)
	return nil
}

func (rlr *roastLevelRepository) GetById(domainRoastLevel *roast_level.Entity, id uint) error {
	var modelRoastLevel model.RoastLevel
	if err := rlr.db.First(&modelRoastLevel, id).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityRoastLevel := entity_model.ModelToRoastLevelEntity(&modelRoastLevel)
	*domainRoastLevel = *entityRoastLevel
	return nil
}

func NewRoastLevelRepository(db *gorm.DB) roast_level.IRoastLevelRepository {
	return &roastLevelRepository{db}
}
