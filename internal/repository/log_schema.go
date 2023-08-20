package repository

import (
	"context"
	"fmt"
	"moonlogs/ent"
	"moonlogs/ent/logschema"
	"moonlogs/internal/config"
)

type LogSchemaRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewLogSchemaRepository(ctx context.Context) *LogSchemaRepository {
	return &LogSchemaRepository{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *LogSchemaRepository) Create(logSchema ent.LogSchema, formattedSchemaName string) (*ent.LogSchema, error) {
	ls, err := r.client.LogSchema.
		Create().
		SetName(formattedSchemaName).
		SetDescription(logSchema.Description).
		SetTitle(logSchema.Title).
		SetFields(logSchema.Fields).
		Save(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating log_schema: %w", err)
	}

	return ls, nil
}

func (r *LogSchemaRepository) GetById(id int) (*ent.LogSchema, error) {
	ls, err := r.client.LogSchema.
		Query().
		Where(logschema.ID(id)).First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_schema: %w", err)
	}

	return ls, nil
}

func (r *LogSchemaRepository) GetByName(name string) (*ent.LogSchema, error) {
	ls, err := r.client.LogSchema.
		Query().
		Where(logschema.Name(name)).
		First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_schema: %w", err)
	}

	return ls, nil
}

func (r *LogSchemaRepository) GetAll() ([]*ent.LogSchema, error) {
	ls, err := r.client.LogSchema.
		Query().All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_schemas: %w", err)
	}

	return ls, nil
}
