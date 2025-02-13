package persistence

import (
	"database/sql"
	"fmt"
	"moonlogs/internal/storage"
	"moonlogs/internal/storage/mongodb_adapter"
	"moonlogs/internal/storage/sqlite_adapter"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SQLITE_ADAPTER  = "sqlite"
	MONGODB_ADAPTER = "mongodb"
)

const (
	MONGODB_DATABASE_NAME = "moonlogs"
)

type Databases struct {
	SqliteReadInstance  *sql.DB
	SqliteWriteInstance *sql.DB
	MongoInstance       *mongo.Client
}

func InitDB(DBAdapter string, DBPath string) (*Databases, error) {
	var databases Databases
	var err error

	switch DBAdapter {
	case MONGODB_ADAPTER:
		mongoInstance, err := initMongoDB(DBPath)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
		}

		databases = Databases{MongoInstance: mongoInstance}
	default:
		sqliteWriteInstance, sqliteReadInstance, err := initSqliteDB(DBPath)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to Sqlite: %w", err)
		}

		databases = Databases{SqliteReadInstance: sqliteReadInstance, SqliteWriteInstance: sqliteWriteInstance}
	}

	return &databases, err
}

type Storages struct {
	ActionStorage              storage.ActionStorage
	ApiTokenStorage            storage.ApiTokenStorage
	RecordStorage              storage.RecordStorage
	SchemaStorage              storage.SchemaStorage
	TagStorage                 storage.TagStorage
	UserStorage                storage.UserStorage
	AlertingRuleStorage        storage.AlertingRuleStorage
	IncidentStorage            storage.IncidentStorage
	NotificationProfileStorage storage.NotificationProfileStorage
}

func InitStorages(storageType string, databases *Databases) Storages {
	switch storageType {
	case MONGODB_ADAPTER:
		mongoDB := databases.MongoInstance.Database(MONGODB_DATABASE_NAME)

		return Storages{
			ActionStorage:              mongodb_adapter.NewActionStorage(mongoDB),
			ApiTokenStorage:            mongodb_adapter.NewApiTokenStorage(mongoDB),
			RecordStorage:              mongodb_adapter.NewRecordStorage(mongoDB),
			SchemaStorage:              mongodb_adapter.NewSchemaStorage(mongoDB),
			TagStorage:                 mongodb_adapter.NewTagStorage(mongoDB),
			UserStorage:                mongodb_adapter.NewUserStorage(mongoDB),
			AlertingRuleStorage:        mongodb_adapter.NewAlertingRuleStorage(mongoDB),
			IncidentStorage:            mongodb_adapter.NewIncidentStorage(mongoDB),
			NotificationProfileStorage: mongodb_adapter.NewNotificationProfileStorage(mongoDB),
		}
	default:
		return Storages{
			ActionStorage:              sqlite_adapter.NewActionStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			ApiTokenStorage:            sqlite_adapter.NewApiTokenStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			RecordStorage:              sqlite_adapter.NewRecordStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			SchemaStorage:              sqlite_adapter.NewSchemaStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			TagStorage:                 sqlite_adapter.NewTagStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			UserStorage:                sqlite_adapter.NewUserStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			AlertingRuleStorage:        sqlite_adapter.NewAlertingRuleStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			IncidentStorage:            sqlite_adapter.NewIncidentStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
			NotificationProfileStorage: sqlite_adapter.NewNotificationProfileStorage(databases.SqliteReadInstance, databases.SqliteWriteInstance),
		}
	}
}
