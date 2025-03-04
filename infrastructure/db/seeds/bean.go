package seeds

import (
	"c0fee-api/common"
	"c0fee-api/model"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBeanSeeds(db *gorm.DB) error {

	// ユーザーを取得
	var user model.User
	var userId uuid.UUID

	if testID := os.Getenv("TEST_USER_ID"); testID == "" {
		if err := db.First(&user).Error; err == nil {
			userId = user.ID
		}
	} else {
		userId = uuid.MustParse(testID)
	}

	var varieties []model.Variety
	if err := db.Find(&varieties).Error; err != nil {
		return err
	}

	beans := []model.Bean{
		{
			Name:            common.StoPoint("Ethiopian Yirgacheffe"),
			UserID:          userId,
			RoasterID:       1,
			CountryID:       1,
			AreaID:          common.ItoPoint(1),
			Varieties:       []model.Variety{varieties[0]},
			ProcessMethodID: common.ItoPoint(1),
			RoastLevel:      model.Medium,
			ImageKey:        common.StoPoint("1/image.png"),
		},
		{
			Name:            common.StoPoint("Colombian"),
			UserID:          userId,
			RoasterID:       2,
			CountryID:       2,
			AreaID:          common.ItoPoint(2),
			FarmID:          common.ItoPoint(1),
			FarmerID:        common.ItoPoint(1),
			Varieties:       []model.Variety{varieties[9]},
			ProcessMethodID: common.ItoPoint(2),
			RoastLevel:      model.Dark,
			ImageKey:        common.StoPoint("2/image.png"),
		},
		{
			Name:            common.StoPoint("ROI Farm Washed"),
			UserID:          userId,
			RoasterID:       3,
			CountryID:       3,
			AreaID:          common.ItoPoint(3),
			FarmID:          common.ItoPoint(2),
			FarmerID:        common.ItoPoint(2),
			Varieties:       []model.Variety{varieties[5], varieties[6]},
			ProcessMethodID: common.ItoPoint(3),
			RoastLevel:      model.Light,
			ImageKey:        common.StoPoint("3/image.png"),
		},
		{
			Name:            common.StoPoint("MandhelingG1 Toba Berkah"),
			UserID:          userId,
			RoasterID:       4,
			CountryID:       4,
			AreaID:          common.ItoPoint(4),
			FarmID:          common.ItoPoint(3),
			Varieties:       []model.Variety{varieties[2]},
			ProcessMethodID: common.ItoPoint(3),
			RoastLevel:      model.MediumDark,
			ImageKey:        common.StoPoint("4/image.png"),
		},
		{
			UserID:          userId,
			RoasterID:       5,
			CountryID:       5,
			AreaID:          common.ItoPoint(5),
			FarmID:          common.ItoPoint(4),
			FarmerID:        common.ItoPoint(4),
			Varieties:       []model.Variety{varieties[7]},
			ProcessMethodID: common.ItoPoint(5),
			RoastLevel:      model.MediumLight,
			ImageKey:        common.StoPoint("5/image.png"),
		},
	}

	for _, bean := range beans {
		if err := db.Create(&bean).Error; err != nil {
			return err
		}
	}

	return nil
}
