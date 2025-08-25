package seeds

import (
	"c0fee-api/model"

	"gorm.io/gorm"
)

func CreateRoasterSeeds(db *gorm.DB) error {
	roasters := []model.Roaster{
		{ID: 1, Name: "Starbucks", Address: "Seattle's Pike Place Market", WebURL: "https://archive.starbucks.com"},
		{ID: 2, Name: "Blue Bottle Coffee", Address: "Oakland, California", WebURL: "https://bluebottlecoffee.com"},
		{ID: 3, Name: "Stumptown Coffee Roasters", Address: "Portland, Oregon", WebURL: "https://www.stumptowncoffee.com"},
		{ID: 4, Name: "Intelligentsia Coffee", Address: "Chicago, Illinois", WebURL: "https://www.intelligentsiacoffee.com"},
		{ID: 5, Name: "Counter Culture Coffee", Address: "Durham, North Carolina", WebURL: "https://counterculturecoffee.com"},
	}

	for _, roaster := range roasters {
		if err := db.Create(&roaster).Error; err != nil {
			return err
		}
	}

	return nil
}
