package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IVarietyRepository interface {
	List(varieties *[]model.Variety) error
}

type varietyRepository struct {
	db *gorm.DB
}

func (vr *varietyRepository) List(varieties *[]model.Variety) error {
	if err := vr.db.Find(&varieties).Error; err != nil {
		return err
	}
	return nil
}

func NewVarietyRepository(db *gorm.DB) IVarietyRepository {
	return &varietyRepository{db}
}