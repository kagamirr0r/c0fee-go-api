package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type ICountryRepository interface {
	GetById(country *model.Country, id uint) error
	List(countries *[]model.Country) error
}

type countryRepository struct {
	db *gorm.DB
}

func (cr *countryRepository) GetById(country *model.Country, id uint) error {
	if err := cr.db.Preload("Areas").Where("id = ?", id).First(country).Error; err != nil {
		return err
	}
	return nil
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
