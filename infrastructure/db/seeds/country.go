package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateCountriesSeeds(db *gorm.DB) error {
	countries := []model.Country{
		{ID: 1, Name: "Ethiopia", Code: "ET"},
		{ID: 2, Name: "Colombia", Code: "CO"},
		{ID: 3, Name: "Kenya", Code: "KE"},
		{ID: 4, Name: "Indonesia", Code: "ID"},
		{ID: 5, Name: "Guatemala", Code: "GT"},
	}

	for _, country := range countries {
		if err := db.Create(&country).Error; err != nil {
			return err
		}
	}

	return nil
}
