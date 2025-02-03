package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateVarietySeeds(db *gorm.DB) error {
	varieties := []model.Variety{
		{Variety: "Bourbon"},
		{Variety: "Typica"},
		{Variety: "Geisha"},
		{Variety: "SL28"},
		{Variety: "Caturra"},
	}

	for _, variety := range varieties {
		if err := db.Create(&variety).Error; err != nil {
			return err
		}
	}

	return nil
}
