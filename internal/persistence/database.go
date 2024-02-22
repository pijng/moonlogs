package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/lib/qrx"
	"os"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
)

var dbInstance *sql.DB

func DB() *sql.DB {
	return dbInstance
}

var schema = `
CREATE TABLE IF NOT EXISTS records (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT,
	created_at INTEGER NOT NULL,
	schema_name TEXT NOT NULL,
	schema_id INTEGER NOT NULL,
	query JSON,
	kind string,
	group_hash TEXT NOT NULL,
	level TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS schemas (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	name TEXT NOT NULL,
	fields JSON,
	kinds JSON,
	tag_id INTEGER,
	retention_days INTEGER
);
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT,
	password_digest TEXT NOT NULL,
	role TEXT NOT NULL,
	tag_ids TEXT,
	token TEXT,
	is_revoked INTEGER DEFAULT 0
);
CREATE TABLE IF NOT EXISTS api_tokens (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	token TEXT,
	token_digest TEXT NOT NULL,
	name TEXT NOT NULL,
	is_revoked INTEGER DEFAULT 0
);
CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_schema_id ON records(schema_id);
CREATE INDEX IF NOT EXISTS idx_schema_name ON records(schema_name);
CREATE INDEX IF NOT EXISTS idx_kind ON records(kind);
CREATE INDEX IF NOT EXISTS idx_level ON records(level);
CREATE INDEX IF NOT EXISTS idx_group_hash ON records(group_hash);`

func InitDB(dataSourceName string) error {
	if dbInstance != nil {
		return nil
	}

	dir := filepath.Dir(dataSourceName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating db dir: %w", err)
	}

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?cache=shared&_fk=1&_journal_mode=WAL", dataSourceName))
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed pinging sqlite: %w", err)
	}

	dbInstance = db

	_, err = qrx.With(dbInstance).Exec(context.Background(), schema)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}
