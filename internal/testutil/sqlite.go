package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/schemas"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/pressly/goose/v3"
)

const SQLITE_ADAPTER = "sqlite"
const SQLITE_TEST_DSN = "test.db"

type ReadDB *sql.DB
type WriteDB *sql.DB

func SetupSqlite() (WriteDB, ReadDB, error) {
	readDB, err := sql.Open(SQLITE_ADAPTER, fmt.Sprintf("file:%s", SQLITE_TEST_DSN))
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open Sqlite for read: %w", err)
	}

	writeDB, err := sql.Open(SQLITE_ADAPTER, fmt.Sprintf("file:%s", SQLITE_TEST_DSN))
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open Sqlite for write: %w", err)
	}

	_, err = writeDB.ExecContext(context.Background(), schemas.SQLITE_SCHEMA)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tables: %w", err)
	}

	if err := goose.SetDialect(SQLITE_ADAPTER); err != nil {
		return nil, nil, fmt.Errorf("failed setting up goose dialect: %w", err)
	}

	if err := goose.Up(writeDB, "../../../migrations"); err != nil {
		return nil, nil, fmt.Errorf("failed applying migrations: %w", err)
	}

	return writeDB, readDB, nil
}

func TeardownSqlite() error {
	err := os.Remove(SQLITE_TEST_DSN)
	if err != nil {
		return fmt.Errorf("failed to remove sqlite test file: %w", err)
	}

	return nil
}
