package migrations

import (
	"c0fee-api/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreateRoastLevelsTable, downCreateRoastLevelsTable)
}

func upCreateRoastLevelsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}
	if err := db.AutoMigrate(&model.RoastLevel{}); err != nil {
		return err
	}

	// デフォルトデータの挿入
	roastLevels := []model.RoastLevel{
		{ID: 1, Name: "Light", Level: 1},
		{ID: 2, Name: "Medium-Light", Level: 2},
		{ID: 3, Name: "Medium", Level: 3},
		{ID: 4, Name: "Medium-Dark", Level: 4},
		{ID: 5, Name: "Dark", Level: 5},
	}

	for _, rl := range roastLevels {
		if err := db.Create(&rl).Error; err != nil {
			return err
		}
	}

	fmt.Println("Successfully migrated roast_levels table")
	return nil
}

func downCreateRoastLevelsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}
	if err := db.Migrator().DropTable("roast_levels"); err != nil {
		return err
	}

	fmt.Println("Successfully rolled back roast_levels table")
	return nil
}
