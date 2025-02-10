package seeds

import (
	"c0fee-api/model"
	"fmt"

	"gorm.io/gorm"
)

func CreateBeanSeeds(db *gorm.DB) error {

	// ユーザーを取得
	var user model.User
	if err := db.First(&user).Error; err != nil {
		return err
	}
	fmt.Print(user)

	beans := []model.Bean{
		{
			Name:            "Ethiopian Yirgacheffe",
			UserID:          user.ID,
			RoasterID:       1,
			ProcessMethodID: 1,
			Countries:       []model.Country{{Name: "Ethiopia", Code: "ET"}},
			Varieties:       []model.Variety{{Variety: "Heirloom"}},
			Area:            "Yirgacheffe",
			RoastLevel:      model.Medium,
			ImageKey:        "1/image.jpg",
		},
		{
			Name:            "Colombian Supremo",
			UserID:          user.ID,
			RoasterID:       2,
			ProcessMethodID: 2,
			Countries:       []model.Country{{Name: "Colombia", Code: "CO"}},
			Varieties:       []model.Variety{{Variety: "Caturra"}},
			Area:            "Antioquia",
			RoastLevel:      model.Dark,
			ImageKey:        "2/image.jpg",
		},
		{
			Name:            "Kenyan AA",
			UserID:          user.ID,
			RoasterID:       3,
			ProcessMethodID: 3,
			Countries:       []model.Country{{Name: "Kenya", Code: "KE"}},
			Varieties:       []model.Variety{{Variety: "SL28"}},
			Area:            "Nyeri",
			RoastLevel:      model.Light,
			ImageKey:        "3/image.jpg",
		},
		{
			Name:            "Sumatra Mandheling",
			UserID:          user.ID,
			RoasterID:       4,
			ProcessMethodID: 4,
			Countries:       []model.Country{{Name: "Indonesia", Code: "ID"}},
			Varieties:       []model.Variety{{Variety: "Typica"}},
			Area:            "Aceh",
			RoastLevel:      model.MediumDark,
			ImageKey:        "4/image.jpg",
		},
		{
			Name:            "Guatemalan Antigua",
			UserID:          user.ID,
			RoasterID:       5,
			ProcessMethodID: 5,
			Countries:       []model.Country{{Name: "Guatemala", Code: "GT"}},
			Varieties:       []model.Variety{{Variety: "Bourbon"}},
			Area:            "Antigua",
			RoastLevel:      model.MediumLight,
			ImageKey:        "5/image.jpg",
		},
	}

	for _, bean := range beans {
		if err := db.Create(&bean).Error; err != nil {
			return err
		}
	}

	return nil
}
