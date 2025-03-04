package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateAreasSeeds(db *gorm.DB) error {
	areas := []model.Area{
		{ID: 1, CountryID: 1, Name: "Yirgacheffe"},
		{ID: 2, CountryID: 2, Name: "San Adolfo, Huila"},
		{ID: 3, CountryID: 3, Name: "Kiambu"},
		{ID: 4, CountryID: 4, Name: "Lumban Julu,Toba Regency,North Sumatra"},
		{ID: 5, CountryID: 5, Name: "Antigua"},
	}

	for _, area := range areas {
		if err := db.Create(&area).Error; err != nil {
			return err
		}
	}

	return nil
}
