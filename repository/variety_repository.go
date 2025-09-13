package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/variety"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type varietyRepository struct {
	db *gorm.DB
}

func (vr *varietyRepository) List(domainVarieties *[]variety.Entity) error {
	var modelVarieties []model.Variety
	if err := vr.db.Find(&modelVarieties).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainVarieties = entity_model.ModelsToVarietyEntities(modelVarieties)
	return nil
}

func NewVarietyRepository(db *gorm.DB) variety.IVarietyRepository {
	return &varietyRepository{db}
}
