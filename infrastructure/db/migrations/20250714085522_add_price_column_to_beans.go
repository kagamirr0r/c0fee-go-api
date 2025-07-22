package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"c0fee-api/model"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upAddPriceColumnToBeans, downAddPriceColumnToBeans)
}

func upAddPriceColumnToBeans(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// Add price column to beans table
	if err := db.Exec("ALTER TABLE beans ADD COLUMN price NUMERIC(10,2) AFTER roast_level"); err != nil {
		return fmt.Errorf("failed to add price column to beans: %w", err)
	}

	// Add currency column to beans table
	if err := db.Exec("ALTER TABLE beans ADD COLUMN currency VARCHAR(3) AFTER price"); err != nil {
		return fmt.Errorf("failed to add currency column to beans: %w", err)
	}
	fmt.Println("Successfully added price and currency column to beans")

	return nil
}

func downAddPriceColumnToBeans(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// Drop price column from beans table
	if err := db.Migrator().DropColumn(&model.Bean{}, "Price"); err != nil {
		return fmt.Errorf("failed to drop price column from beans: %w", err)
	}

	// Drop currency column from beans table
	if err := db.Migrator().DropColumn(&model.Bean{}, "Currency"); err != nil {
		return fmt.Errorf("failed to drop currency column from beans: %w", err)
	}
	fmt.Println("Successfully dropped price and currency column from beans")

	return nil
}
