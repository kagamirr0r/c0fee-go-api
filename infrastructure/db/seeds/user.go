package seeds

import (
	"c0fee-api/model"

	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUserSeeds(db *gorm.DB) error {
	var users []model.User
	if testID := os.Getenv("TEST_USER_ID"); testID != "" {
		users = append(users, model.User{
			ID:        uuid.MustParse(testID),
			Name:      "Test User",
			AvatarKey: testID + "/avatar.png",
		})
	}
	users = append(users, model.User{ID: uuid.New(), Name: "John Doe", AvatarKey: ""})

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
