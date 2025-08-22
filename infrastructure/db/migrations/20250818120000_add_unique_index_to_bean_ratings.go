package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUniqueIndexToBeanRatings, downAddUniqueIndexToBeanRatings)
}

func upAddUniqueIndexToBeanRatings(ctx context.Context, tx *sql.Tx) error {
	// bean_id と user_id の複合ユニークインデックスを追加
	query := `
		ALTER TABLE bean_ratings 
		ADD CONSTRAINT unique_idx_bean_user UNIQUE (bean_id, user_id);
	`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("failed to add unique index: %w", err)
	}

	fmt.Println("Successfully added unique index (bean_id, user_id) to bean_ratings")
	return nil
}

func downAddUniqueIndexToBeanRatings(ctx context.Context, tx *sql.Tx) error {
	// ユニークインデックスを削除
	query := `
		ALTER TABLE bean_ratings 
		DROP CONSTRAINT IF EXISTS unique_idx_bean_user;
	`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("failed to drop unique index: %w", err)
	}

	fmt.Println("Successfully dropped unique index unique_idx_bean_user from bean_ratings")
	return nil
}
