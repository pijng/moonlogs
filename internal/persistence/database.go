package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/lib/qrx"

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
	created_at INTEGER,
	schema_name TEXT,
	schema_id INTEGER,
	query JSON,
	kind string,
	group_hash TEXT,
	level TEXT
);
CREATE TABLE IF NOT EXISTS schemas (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	description TEXT,
	name TEXT,
	fields JSON,
	kinds JSON,
	tag_id INTEGER,
	retention_days INTEGER
);
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	password TEXT,
	password_digest TEXT,
	role TEXT,
	tag_ids TEXT,
	token TEXT
);
CREATE TABLE IF NOT EXISTS api_tokens (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	token TEXT,
	token_digest TEXT,
	name TEXT,
	is_revoked INTEGER
);
CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT
);

CREATE INDEX IF NOT EXISTS idx_schema_id ON records(schema_id);
CREATE INDEX IF NOT EXISTS idx_schema_name ON records(schema_name);
CREATE INDEX IF NOT EXISTS idx_created_at ON records(created_at);
CREATE INDEX IF NOT EXISTS idx_created_at ON records(kind);
CREATE INDEX IF NOT EXISTS idx_created_at ON records(group_hash);`

func InitDB(dataSourceName string) error {
	if dbInstance != nil {
		return nil
	}

	db, err := sql.Open("sqlite", fmt.Sprintf("%s:%s?%s", "file", dataSourceName, "cache=shared&_fk=1"))
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
