package seeds

import (
	"c0fee-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUserSeeds(db *gorm.DB) error {
	users := []model.User{{ID: uuid.New(), Name: "John Doe"}, {ID: uuid.New(), Name: "Jane Doe"}}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
