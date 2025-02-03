package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateProcessMethodSeeds(db *gorm.DB) error {
	methods := []model.ProcessMethod{
		{Method: "Natural"},
		{Method: "Washed"},
		{Method: "Honey"},
		{Method: "Wet-Hulled"},
		{Method: "Pulped Natural"},
	}

	for _, method := range methods {
		if err := db.Create(&method).Error; err != nil {
			return err
		}
	}

	return nil
}
