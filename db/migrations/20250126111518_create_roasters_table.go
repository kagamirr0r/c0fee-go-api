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
	goose.AddMigrationContext(upCreateRoastersTable, downCreateRoastersTable)
}

func upCreateRoastersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Roaster{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated roasters")
	return nil
}

func downCreateRoastersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("roasters"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked roasters")
	return nil
}
