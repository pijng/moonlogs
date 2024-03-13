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

var sqliteReadInstance *sql.DB
var sqliteWriteInstance *sql.DB

func SqliteReadDB() *sql.DB {
	return sqliteReadInstance
}

func SqliteWriteDB() *sql.DB {
	return sqliteWriteInstance
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
		sqliteWriteInstance, sqliteReadInstance, err = initSqliteDB(cfg.DBPath)
	}

	return err
}
