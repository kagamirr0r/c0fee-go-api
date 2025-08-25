package migrations

import (
	"c0fee-api/infrastructure/db"
	"c0fee-api/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateFarmersTable, downCreateFarmersTable)
}

func upCreateFarmersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Farmer{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated farms")
	return nil
}

func downCreateFarmersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("farmers"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked farms")
	return nil
}
