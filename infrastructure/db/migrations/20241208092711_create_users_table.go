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
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated")
	return nil
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("users"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked")
	return nil
}
