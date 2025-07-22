package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upProcessMethodsRenameColumnMethodToName, downProcessMethodsRenameColumnMethodToName)
}

func upProcessMethodsRenameColumnMethodToName(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// process_methodsテーブルのmethodカラムをnameカラムにリネーム
	if err := db.Migrator().RenameColumn("process_methods", "method", "name"); err != nil {
		return fmt.Errorf("failed to rename process_methods.method to name: %w", err)
	}
	fmt.Println("Successfully renamed process_methods.method to name")
	return nil
}

func downProcessMethodsRenameColumnMethodToName(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// process_methodsテーブルのmethodカラムをnameカラムにリネーム
	if err := db.Migrator().RenameColumn("process_methods", "name", "method"); err != nil {
		return fmt.Errorf("failed to rename process_methods.name to method: %w", err)
	}
	return nil
}
