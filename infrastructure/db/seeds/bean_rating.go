package seeds

import (
	"c0fee-api/model"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBeanRatingSeeds(db *gorm.DB) error {
	var user model.User
	var userId uuid.UUID

	if testID := os.Getenv("TEST_USER_ID"); testID == "" {
		if err := db.First(&user).Error; err == nil {
			userId = user.ID
		}
	} else {
		userId = uuid.MustParse(testID)
	}
	ratings := []model.BeanRating{
		{
			BeanID:     1,
			UserID:     userId,
			Bitterness: 1,
			Acidity:    1,
			Body:       1,
			FlavorNote: "Berries, Chocolate, Nuts",
		},
		{
			BeanID:     2,
			UserID:     userId,
			Bitterness: 2,
			Acidity:    2,
			Body:       2,
			FlavorNote: "Citrus, Caramel, Floral",
		},
		{
			BeanID:     3,
			UserID:     userId,
			Bitterness: 3,
			Acidity:    3,
			Body:       3,
			FlavorNote: "Tropical Fruit, Honey, Spices",
		},
		{
			BeanID:     4,
			UserID:     userId,
			Bitterness: 4,
			Acidity:    4,
			Body:       4,
			FlavorNote: "Earthy, Smoky, Woody",
		},
		{
			BeanID:     5,
			UserID:     userId,
			Bitterness: 5,
			Acidity:    5,
			Body:       5,
			FlavorNote: "Dark Chocolate, Dried Fruits, Malty",
		},
	}

	for _, rating := range ratings {
		if err := db.Create(&rating).Error; err != nil {
			return err
		}
	}

	return nil
}
