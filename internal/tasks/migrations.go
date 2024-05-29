package tasks

import (
	"database/sql"
	"embed"
	"fmt"
	"moonlogs/internal/persistence"

	"github.com/pressly/goose/v3"
)

func Migrate(dbAdapter string, databases *persistence.Databases, embedMigrations embed.FS) error {
	var err error

	switch dbAdapter {
	case persistence.MONGODB_ADAPTER:
	default:
		err = runSqliteMigrations(databases.SqliteWriteInstance, embedMigrations)
	}

	return err
}

func runSqliteMigrations(writeInstance *sql.DB, embedMigrations embed.FS) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(persistence.SQLITE_ADAPTER); err != nil {
		return fmt.Errorf("failed setting up goose dialect: %w", err)
	}

	if err := goose.Up(writeInstance, "migrations"); err != nil {
		return fmt.Errorf("failed applying migrations: %w", err)
	}

	return nil
}
