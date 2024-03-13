package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

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
`

type ReadDB *sql.DB
type WriteDB *sql.DB

func initSqliteDB(dataSourceName string) (WriteDB, ReadDB, error) {
	dir := filepath.Dir(dataSourceName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating db dir: %w", err)
	}

	readDB, err := initReadDB(dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("initializing read db: %w", err)
	}

	writeDB, err := initWriteDB(dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("initializing write db: %w", err)
	}

	_, err = writeDB.ExecContext(context.Background(), schema)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return writeDB, readDB, nil

}

func initReadDB(dataSourceName string) (*sql.DB, error) {
	params := url.Values{}
	params.Add("_fk", "1")
	params.Add("mode", "ro")
	params.Add("_journal_mode", "wal")
	params.Add("_pragma=analysis_limit", "400")
	params.Add("_pragma=synchronous", "normal")
	params.Add("_pragma=temp_store", "memory")
	params.Add("_pragma=mmap_size", "536870912")
	params.Add("_pragma=busy_timeout", "5000")
	params.Add("_pragma=cache_size", "-512000")

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?%s", dataSourceName, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	db.SetMaxOpenConns(runtime.GOMAXPROCS(0) * 2)
	db.SetMaxIdleConns(runtime.GOMAXPROCS(0) * 2)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed pinging sqlite: %w", err)
	}

	return db, nil
}

func initWriteDB(dataSourceName string) (*sql.DB, error) {
	params := url.Values{}
	params.Add("_fk", "1")
	params.Add("mode", "rw")
	params.Add("_journal_mode", "wal")
	params.Add("_pragma=analysis_limit", "400")
	params.Add("_pragma=synchronous", "normal")
	params.Add("_pragma=temp_store", "memory")
	params.Add("_pragma=mmap_size", "536870912")
	params.Add("_pragma=busy_timeout", "5000")
	params.Add("_pragma=cache_size", "-512000")

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?%s", dataSourceName, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed pinging sqlite: %w", err)
	}

	return db, nil
}
