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
	goose.AddMigrationContext(upCreateAreasTable, downCreateAreasTable)
}

func upCreateAreasTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Area{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated areas")
	return nil
}

func downCreateAreasTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("areas"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked areas")
	return nil
}
