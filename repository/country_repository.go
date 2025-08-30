package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/entity"
	domainRepo "c0fee-api/domain/repository"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type countryRepository struct {
	db *gorm.DB
}

func (cr *countryRepository) GetById(domainCountry *entity.Country, id uint) error {
	var modelCountry model.Country
	if err := cr.db.
		Preload("Areas").
		Where("id = ?", id).
		First(&modelCountry).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityCountry := entity_model.ModelCountryToEntity(&modelCountry)
	*domainCountry = *entityCountry
	return nil
}

func (cr *countryRepository) List(domainCountries *[]entity.Country) error {
	var modelCountries []model.Country
	if err := cr.db.Find(&modelCountries).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainCountries = entity_model.ModelCountriesToEntities(modelCountries)
	return nil
}

func NewCountryRepository(db *gorm.DB) domainRepo.ICountryRepository {
	return &countryRepository{db}
}
