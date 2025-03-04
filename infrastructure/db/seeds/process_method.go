package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateProcessMethodSeeds(db *gorm.DB) error {
	methods := []model.ProcessMethod{
		{ID: 1, Method: "Natural"},
		{ID: 2, Method: "Anaerobic Natural"},
		{ID: 3, Method: "Washed"},
		{ID: 4, Method: "Smatra"},
		{ID: 5, Method: "Honey"},
		{ID: 6, Method: "Wet-Hulled"},
		{ID: 7, Method: "Pulped Natural"},
	}

	for _, method := range methods {
		if err := db.Create(&method).Error; err != nil {
			return err
		}
	}

	return nil
}
