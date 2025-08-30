package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type varietyRepository struct {
	db *gorm.DB
}

func (vr *varietyRepository) List(domainVarieties *[]entity.Variety) error {
	var modelVarieties []model.Variety
	if err := vr.db.Find(&modelVarieties).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainVarieties = entity_model.ModelVarietiesToEntities(modelVarieties)
	return nil
}

func NewVarietyRepository(db *gorm.DB) domainRepo.IVarietyRepository {
	return &varietyRepository{db}
}
