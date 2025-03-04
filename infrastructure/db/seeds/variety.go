package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateVarietySeeds(db *gorm.DB) error {
	varieties := []model.Variety{
		{ID: 1, Variety: "Heirloom"},
		{ID: 2, Variety: "Typica"},
		{ID: 3, Variety: "Bourbon"},
		{ID: 4, Variety: "Catuai"},
		{ID: 5, Variety: "SL34"},
		{ID: 6, Variety: "SL28"},
		{ID: 7, Variety: "Ruiru11"},
		{ID: 8, Variety: "Pacamara"},
		{ID: 9, Variety: "Mundo Novo"},
		{ID: 10, Variety: "Chiroso"},
		{ID: 11, Variety: "Geisha"},
		{ID: 12, Variety: "Caturra"},
	}

	for _, variety := range varieties {
		if err := db.Create(&variety).Error; err != nil {
			return err
		}
	}

	return nil
}
