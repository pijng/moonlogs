package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/schemas"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

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

	_, err = writeDB.ExecContext(context.Background(), schemas.SQLITE_SCHEMA)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return writeDB, readDB, nil
}

func initReadDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?%s", dataSourceName, buildQueryString()))
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
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?%s", dataSourceName, buildQueryString()))
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

func buildQueryString() string {
	var parts []string

	parts = append(parts, "_fk=1")
	parts = append(parts, "_journal_mode=wal")
	parts = append(parts, "_pragma=analysis_limit=400")
	parts = append(parts, "_pragma=synchronous=normal")
	parts = append(parts, "_pragma=temp_store=memory")
	parts = append(parts, "_pragma=mmap_size=536870912")
	parts = append(parts, "_pragma=busy_timeout=5000")
	parts = append(parts, "_pragma=cache_size=-512000")

	return strings.Join(parts, "&")
}
