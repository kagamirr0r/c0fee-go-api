package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upAddRoastLevelIdToBeans, downAddRoastLevelIdToBeans)
}

func upAddRoastLevelIdToBeans(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// roast_level_id カラムを追加
	if err := db.Exec("ALTER TABLE beans ADD COLUMN roast_level_id INTEGER NOT NULL DEFAULT 3").Error; err != nil {
		return err
	}

	// 外部キー制約を追加
	if err := db.Exec("ALTER TABLE beans ADD CONSTRAINT fk_beans_roast_level FOREIGN KEY (roast_level_id) REFERENCES roast_levels(id)").Error; err != nil {
		return err
	}

	// 既存のroast_levelカラムからデータを移行
	migrations := []struct {
		oldValue string
		newID    int
	}{
		{"Light", 1},
		{"Medium-Light", 2},
		{"Medium", 3},
		{"Medium-Dark", 4},
		{"Dark", 5},
	}

	for _, migration := range migrations {
		if err := db.Exec("UPDATE beans SET roast_level_id = ? WHERE roast_level = ?", migration.newID, migration.oldValue).Error; err != nil {
			return err
		}
	}

	// 古いroast_levelカラムを削除
	if err := db.Exec("ALTER TABLE beans DROP COLUMN roast_level").Error; err != nil {
		return err
	}

	fmt.Println("Successfully migrated beans table to use roast_level_id")
	return nil
}

func downAddRoastLevelIdToBeans(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to create GORM instance: %w", err)
	}

	// roast_levelカラムを再作成
	if err := db.Exec("ALTER TABLE beans ADD COLUMN roast_level VARCHAR(255) NOT NULL DEFAULT 'Medium'").Error; err != nil {
		return err
	}

	// データを逆変換
	migrations := []struct {
		id       int
		newValue string
	}{
		{1, "Light"},
		{2, "Medium-Light"},
		{3, "Medium"},
		{4, "Medium-Dark"},
		{5, "Dark"},
	}

	for _, migration := range migrations {
		if err := db.Exec("UPDATE beans SET roast_level = ? WHERE roast_level_id = ?", migration.newValue, migration.id).Error; err != nil {
			return err
		}
	}

	// 外部キー制約を削除
	if err := db.Exec("ALTER TABLE beans DROP CONSTRAINT fk_beans_roast_level").Error; err != nil {
		return err
	}

	// roast_level_idカラムを削除
	if err := db.Exec("ALTER TABLE beans DROP COLUMN roast_level_id").Error; err != nil {
		return err
	}

	fmt.Println("Successfully rolled back beans table roast_level_id migration")
	return nil
}
