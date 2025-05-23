package repository

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

type IRoasterRepository interface {
	List(countries *[]model.Roaster) error
}

type roasterRepository struct {
	db *gorm.DB
}

func (cr *roasterRepository) List(roaster *[]model.Roaster) error {
	if err := cr.db.Find(&roaster).Error; err != nil {
		return err
	}
	return nil
}

func NewRoasterRepository(db *gorm.DB) IRoasterRepository {
	return &roasterRepository{db}
}
