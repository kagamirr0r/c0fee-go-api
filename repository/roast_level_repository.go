package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IRoastLevelRepository interface {
	GetAll(roastLevels *[]model.RoastLevel) error
	GetById(roastLevel *model.RoastLevel, id uint) error
}

type roastLevelRepository struct {
	db *gorm.DB
}

func (rlr *roastLevelRepository) GetAll(roastLevels *[]model.RoastLevel) error {
	if err := rlr.db.Order("level ASC").Find(roastLevels).Error; err != nil {
		return err
	}
	return nil
}

func (rlr *roastLevelRepository) GetById(roastLevel *model.RoastLevel, id uint) error {
	if err := rlr.db.First(roastLevel, id).Error; err != nil {
		return err
	}
	return nil
}

func NewRoastLevelRepository(db *gorm.DB) IRoastLevelRepository {
	return &roastLevelRepository{db}
}
