package repository

import (
	"c0fee-api/common/converter/entity_model"
	"c0fee-api/domain/country"
	"c0fee-api/model"

	"gorm.io/gorm"
)

type countryRepository struct {
	db *gorm.DB
}

func (cr *countryRepository) GetById(domainCountry *country.Entity, id uint) error {
	var modelCountry model.Country
	if err := cr.db.
		Preload("Areas").
		Where("id = ?", id).
		First(&modelCountry).Error; err != nil {
		return err
	}

	// Convert model to domain entity
	entityCountry := entity_model.ModelToCountryEntity(&modelCountry)
	*domainCountry = *entityCountry
	return nil
}

func (cr *countryRepository) List(domainCountries *[]country.Entity) error {
	var modelCountries []model.Country
	if err := cr.db.Find(&modelCountries).Error; err != nil {
		return err
	}

	// Convert model slice to domain entity slice
	*domainCountries = entity_model.ModelsToCountryEntities(modelCountries)
	return nil
}

func NewCountryRepository(db *gorm.DB) country.ICountryRepository {
	return &countryRepository{db}
}
