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
	goose.AddMigrationContext(upCreateCountriesTable, downCreateCountriesTable)
}

func upCreateCountriesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Country{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated countries")
	return nil
}

func downCreateCountriesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("countries"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked countries")
	return nil
}
