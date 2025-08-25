package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateFarmersSeeds(db *gorm.DB) error {
	farmers := []model.Farmer{
		{ID: 1, FarmID: 1, Name: "Arsenio Muñoz"},
		{ID: 2, FarmID: 2, Name: "Kenya man"},
		{ID: 3, FarmID: 3, Name: "Indonesia man"},
		{ID: 4, FarmID: 4, Name: "Guatemala man"},
	}

	for _, farmer := range farmers {
		if err := db.Create(&farmer).Error; err != nil {
			return err
		}
	}

	return nil
}
