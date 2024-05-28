package persistence

import (
	"database/sql"
	"moonlogs/internal/config"
	"moonlogs/internal/storage"
	"moonlogs/internal/storage/mongodb_adapter"
	"moonlogs/internal/storage/sqlite_adapter"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SQLITE_ADAPTER        = "sqlite"
	MONGODB_ADAPTER       = "mongodb"
	MONGODB_DATABASE_NAME = "moonlogs"
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

type Storages struct {
	ActionStorage   storage.ActionStorage
	ApiTokenStorage storage.ApiTokenStorage
	RecordStorage   storage.RecordStorage
	SchemaStorage   storage.SchemaStorage
	TagStorage      storage.TagStorage
	UserStorage     storage.UserStorage
}

func InitStorages(storageType string) Storages {
	switch storageType {
	case MONGODB_ADAPTER:
		mongoDB := mongoInstance.Database(MONGODB_DATABASE_NAME)

		return Storages{
			ActionStorage:   mongodb_adapter.NewActionStorage(mongoDB),
			ApiTokenStorage: mongodb_adapter.NewApiTokenStorage(mongoDB),
			RecordStorage:   mongodb_adapter.NewRecordStorage(mongoDB),
			SchemaStorage:   mongodb_adapter.NewSchemaStorage(mongoDB),
			TagStorage:      mongodb_adapter.NewTagStorage(mongoDB),
			UserStorage:     mongodb_adapter.NewUserStorage(mongoDB),
		}
	default:
		return Storages{
			ActionStorage:   sqlite_adapter.NewActionStorage(sqliteReadInstance, sqliteWriteInstance),
			ApiTokenStorage: sqlite_adapter.NewApiTokenStorage(sqliteReadInstance, sqliteWriteInstance),
			RecordStorage:   sqlite_adapter.NewRecordStorage(sqliteReadInstance, sqliteWriteInstance),
			SchemaStorage:   sqlite_adapter.NewSchemaStorage(sqliteReadInstance, sqliteWriteInstance),
			TagStorage:      sqlite_adapter.NewTagStorage(sqliteReadInstance, sqliteWriteInstance),
			UserStorage:     sqlite_adapter.NewUserStorage(sqliteReadInstance, sqliteWriteInstance),
		}
	}
}
