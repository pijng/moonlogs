package persistence

import (
	"database/sql"
	"moonlogs/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SQLITE_ADAPTER  = "sqlite"
	MONGODB_ADAPTER = "mongodb"
)

var sqliteInstance *sql.DB

func SqliteDB() *sql.DB {
	return sqliteInstance
}

var mongoInstance *mongo.Client

func MongoDB() *mongo.Client {
	return mongoInstance
}

func InitDB(cfg *config.Config) error {
	var err error

	switch cfg.DBAdapter {
	case MONGODB_ADAPTER:
		mongoInstance, err = initMongoDB(cfg.DBPath)
	default:
		sqliteInstance, err = initSqliteDB(cfg.DBPath)
	}

	return err
}
