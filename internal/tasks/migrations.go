package tasks

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB, dbAdapter string, embedMigrations embed.FS) error {
	goose.SetBaseFS(embedMigrations)

	var dialect string
	switch dbAdapter {
	default:
		dialect = "sqlite"
	}

	if err := goose.SetDialect(dialect); err != nil {
		return fmt.Errorf("failed setting up goose dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed applying migrations: %w", err)
	}

	return nil
}
