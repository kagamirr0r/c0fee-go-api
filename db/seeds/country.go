package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateCountriesSeeds(db *gorm.DB) error {
	countries := []model.Country{{Name: "Japan", Code: "JP"}}

	for _, country := range countries {
		if err := db.Create(country).Error; err != nil {
			return err
		}
	}

	return nil
}
