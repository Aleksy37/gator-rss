package database

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/Aleksy37/gator-rss/sql/schema"
)

func RunMigrations(db *sql.DB) error {
	// 1. Tell goose to use your embedded filesystem
	goose.SetBaseFS(schema.FS)

	// 2. Set the database dialect
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	// 3. Run all 'up' migrations located in the root of the embedded FS
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}