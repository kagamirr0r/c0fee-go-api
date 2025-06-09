package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IAreaRepository interface {
	GetById(area *model.Area, id uint) error
	List(countries *[]model.Area) error
}

type areaRepository struct {
	db *gorm.DB
}

func (ar *areaRepository) GetById(area *model.Area, id uint) error {
	if err := ar.db.Preload("Farms").Where("id = ?", id).First(area).Error; err != nil {
		return err
	}
	return nil
}

func (ar *areaRepository) List(areas *[]model.Area) error {
	if err := ar.db.Find(areas).Error; err != nil {
		return err
	}
	return nil
}

func NewAreaRepository(db *gorm.DB) IAreaRepository {
	return &areaRepository{db}
}
