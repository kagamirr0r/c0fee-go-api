package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateVarietySeeds(db *gorm.DB) error {
	varieties := []model.Variety{
		{ID: 1, Name: "Heirloom"},
		{ID: 2, Name: "Typica"},
		{ID: 3, Name: "Bourbon"},
		{ID: 4, Name: "Catuai"},
		{ID: 5, Name: "SL34"},
		{ID: 6, Name: "SL28"},
		{ID: 7, Name: "Ruiru11"},
		{ID: 8, Name: "Pacamara"},
		{ID: 9, Name: "Mundo Novo"},
		{ID: 10, Name: "Chiroso"},
		{ID: 11, Name: "Geisha"},
		{ID: 12, Name: "Caturra"},
	}

	for _, variety := range varieties {
		if err := db.Create(&variety).Error; err != nil {
			return err
		}
	}

	return nil
}
