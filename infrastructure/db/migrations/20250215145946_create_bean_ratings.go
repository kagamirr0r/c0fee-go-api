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
	goose.AddMigrationContext(upCreateBeanRatings, downCreateBeanRatings)
}

func upCreateBeanRatings(ctx context.Context, tx *sql.Tx) error {
	dbConn := db.NewDB()
	if err := dbConn.AutoMigrate(&model.BeanRating{}); err != nil {
		return err
	}

	fmt.Println("SuccessFully migrated bean_ratings")
	return nil
}

func downCreateBeanRatings(ctx context.Context, tx *sql.Tx) error {
	dbConn := db.NewDB()
	if err := dbConn.Migrator().DropTable("bean_ratings"); err != nil {
		return err
	}

	fmt.Println("SuccessFully Rollbacked bean_ratings")
	return nil
}
