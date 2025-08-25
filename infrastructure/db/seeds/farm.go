package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateFarmsSeeds(db *gorm.DB) error {
	farms := []model.Farm{
		{ID: 1, AreaID: 2, Name: "Los Cauchos"},
		{ID: 2, AreaID: 3, Name: "ROI Farm"},
		{ID: 3, AreaID: 4, Name: "Small farmers"},
		{ID: 4, AreaID: 5, Name: "San Juan"},
	}

	for _, farm := range farms {
		if err := db.Create(&farm).Error; err != nil {
			return err
		}
	}

	return nil
}
