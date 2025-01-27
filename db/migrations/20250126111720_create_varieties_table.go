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
	goose.AddMigrationContext(upCreateVarietiesTable, downCreateVarietiesTable)
}

func upCreateVarietiesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Variety{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated varieties")
	return nil
}

func downCreateVarietiesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("varieties"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked varieties")
	return nil
}
