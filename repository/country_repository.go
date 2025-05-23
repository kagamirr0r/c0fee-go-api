package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type ICountryRepository interface {
	List(countries *[]model.Country) error
}

type countryRepository struct {
	db *gorm.DB
}

func (cr *countryRepository) List(countries *[]model.Country) error {
	if err := cr.db.Find(&countries).Error; err != nil {
		return err
	}
	return nil
}

func NewCountryRepository(db *gorm.DB) ICountryRepository {
	return &countryRepository{db}
}
