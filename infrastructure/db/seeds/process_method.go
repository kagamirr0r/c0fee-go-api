package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateProcessMethodSeeds(db *gorm.DB) error {
	methods := []model.ProcessMethod{
		{ID: 1, Name: "Natural"},
		{ID: 2, Name: "Anaerobic Natural"},
		{ID: 3, Name: "Washed"},
		{ID: 4, Name: "Smatra"},
		{ID: 5, Name: "Honey"},
		{ID: 6, Name: "Wet-Hulled"},
		{ID: 7, Name: "Pulped Natural"},
	}

	for _, method := range methods {
		if err := db.Create(&method).Error; err != nil {
			return err
		}
	}

	return nil
}
