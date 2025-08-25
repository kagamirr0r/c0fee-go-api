package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IFarmRepository interface {
	GetById(farm *model.Farm, id uint) error
	List(countries *[]model.Farm) error
}

type farmRepository struct {
	db *gorm.DB
}

func (ar *farmRepository) GetById(farm *model.Farm, id uint) error {
	if err := ar.db.Preload("Farmers").Where("id = ?", id).First(farm).Error; err != nil {
		return err
	}
	return nil
}

func (ar *farmRepository) List(farms *[]model.Farm) error {
	if err := ar.db.Find(farms).Error; err != nil {
		return err
	}
	return nil
}

func NewFarmRepository(db *gorm.DB) IFarmRepository {
	return &farmRepository{db}
}
