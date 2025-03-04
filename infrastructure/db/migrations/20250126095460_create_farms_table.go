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
	goose.AddMigrationContext(upCreateFarmsTable, downCreateFarmsTable)
}

func upCreateFarmsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Farm{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated farms")
	return nil
}

func downCreateFarmsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("farms"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked farms")
	return nil
}
