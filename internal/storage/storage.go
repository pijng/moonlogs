package storage

import (
	"context"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage/sqlite_adapter"
)

type UserStorage interface {
	CreateInitialAdmin(admin entities.User) (*entities.User, error)
	CreateUser(user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error
	GetAllUsers() ([]*entities.User, error)
	GetSystemUser() (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	GetUsersByTagID(id int) ([]*entities.User, error)
	GetUserByToken(token string) (*entities.User, error)
	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	UpdateUserTokenByID(id int, token string) error
}

func NewUserStorage(ctx context.Context, storageType string) UserStorage {
	var storageInstance UserStorage

	switch storageType {
	default:
		storageInstance = sqlite_adapter.NewUserStorage(ctx)
	}

	return storageInstance
}

type TagStorage interface {
	CreateTag(tag entities.Tag) (*entities.Tag, error)
	DeleteTagByID(id int) error
	GetAllTags() ([]*entities.Tag, error)
	GetTagByID(id int) (*entities.Tag, error)
	UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error)
}

func NewTagStorage(ctx context.Context, storageType string) TagStorage {
	var storageInstance TagStorage

	switch storageType {
	default:
		storageInstance = sqlite_adapter.NewTagStorage(ctx)
	}

	return storageInstance
}

type SchemaStorage interface {
	CreateSchema(schema entities.Schema) (*entities.Schema, error)
	DeleteSchemaByID(id int) error
	GetAllSchemas() ([]*entities.Schema, error)
	GetById(id int) (*entities.Schema, error)
	GetByTagID(id int) ([]*entities.Schema, error)
	GetByName(name string) (*entities.Schema, error)
	GetSchemasByTitleOrDescription(title string, description string) ([]*entities.Schema, error)
	UpdateSchemaByID(id int, schema entities.Schema) (*entities.Schema, error)
}

func NewSchemaStorage(ctx context.Context, storageType string) SchemaStorage {
	var storageInstance SchemaStorage

	switch storageType {
	default:
		storageInstance = sqlite_adapter.NewSchemaStorage(ctx)
	}

	return storageInstance
}

type RecordStorage interface {
	CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error)
	DeleteByIDs(ids []int) error
	FindStale(schemaID int, threshold int64) ([]*entities.Record, error)
	GetAllRecords(limit int, offset int) ([]*entities.Record, error)
	GetAllRecordsCount() (int, error)
	GetRecordByID(id int) (*entities.Record, error)
	GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error)
	GetRecordsByQuery(record entities.Record, limit int, offset int) ([]*entities.Record, error)
	GetRecordsCountByQuery(record entities.Record) (int, error)
}

func NewRecordStorage(ctx context.Context, storageType string) RecordStorage {
	var storageInstance RecordStorage

	switch storageType {
	default:
		storageInstance = sqlite_adapter.NewRecordStorage(ctx)
	}

	return storageInstance
}

type ApiTokenStorage interface {
	CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error)
	DeleteApiTokenByID(id int) error
	GetAllApiTokens() ([]*entities.ApiToken, error)
	GetApiTokenByDigest(digest string) (*entities.ApiToken, error)
	GetApiTokenByID(id int) (*entities.ApiToken, error)
	UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error)
}

func NewApiTokenStorage(ctx context.Context, storageType string) ApiTokenStorage {
	var storageInstance ApiTokenStorage

	switch storageType {
	default:
		storageInstance = sqlite_adapter.NewApiTokenStorage(ctx)
	}

	return storageInstance
}
