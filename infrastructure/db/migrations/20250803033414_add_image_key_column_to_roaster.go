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
	goose.AddMigrationContext(upAddImageKeyColumnToRoaster, downAddImageKeyColumnToRoaster)
}

func upAddImageKeyColumnToRoaster(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// Add image key column to roasters table
	if err := db.Migrator().AddColumn(&model.Roaster{}, "image_key"); err != nil {
		return fmt.Errorf("failed to add image key column to roasters: %w", err)
	}

	fmt.Println("Successfully added image key column to roasters")

	return nil
}

func downAddImageKeyColumnToRoaster(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// Drop image key column from roasters table
	if err := db.Migrator().DropColumn(&model.Roaster{}, "image_key"); err != nil {
		return fmt.Errorf("failed to drop image key column from roasters: %w", err)
	}

	fmt.Println("Successfully dropped image key column from roasters")

	return nil
}
