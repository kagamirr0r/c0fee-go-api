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
	goose.AddMigrationContext(upCreateBeansTable, downCreateBeansTable)
}

func upCreateBeansTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.Bean{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated beans")
	return nil
}

func downCreateBeansTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("beans"); err != nil {
		return err
	}
	if err := dbConn.Migrator().DropTable("bean_countries"); err != nil {
		return err
	}
	if err := dbConn.Migrator().DropTable("bean_varieties"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked beans")
	return nil
}
