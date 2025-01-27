package migrations

import (
	"c0fee-api/db"
	"c0fee-api/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateProcessMethodsTable, downCreateProcessMethodsTable)
}

func upCreateProcessMethodsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.ProcessMethod{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated process_methods")
	return nil
}

func downCreateProcessMethodsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("process_methods"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked process_methods")
	return nil
}
