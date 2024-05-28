package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type SchemaStorage interface {
	CreateSchema(ctx context.Context, schema entities.Schema) (*entities.Schema, error)
	DeleteSchemaByID(ctx context.Context, id int) error
	GetAllSchemas(ctx context.Context) ([]*entities.Schema, error)
	GetById(ctx context.Context, id int) (*entities.Schema, error)
	GetByTagID(ctx context.Context, id int) ([]*entities.Schema, error)
	GetByName(ctx context.Context, name string) (*entities.Schema, error)
	GetSchemasByTitleOrDescription(ctx context.Context, title string, description string) ([]*entities.Schema, error)
	UpdateSchemaByID(ctx context.Context, id int, schema entities.Schema) (*entities.Schema, error)
}
