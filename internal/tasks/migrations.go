package tasks

import (
	"embed"
	"fmt"
	"moonlogs/internal/persistence"

	"github.com/pressly/goose/v3"
)

func Migrate(dbAdapter string, embedMigrations embed.FS) error {
	var err error

	switch dbAdapter {
	case persistence.MONGODB_ADAPTER:
	default:
		err = runGooseMigrations(embedMigrations)
	}

	return err
}

func runGooseMigrations(embedMigrations embed.FS) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(persistence.SQLITE_ADAPTER); err != nil {
		return fmt.Errorf("failed setting up goose dialect: %w", err)
	}

	if err := goose.Up(persistence.SqliteDB(), "migrations"); err != nil {
		return fmt.Errorf("failed applying migrations: %w", err)
	}

	return nil
}
