package repository

import (
	"context"
	"fmt"
	"moonlogs/ent"
	"moonlogs/ent/logschema"
	"moonlogs/ent/schema"
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
	var formattedFields []schema.Field

	for _, field := range logSchema.Fields {
		var formattedField schema.Field

		if field.Type == "" {
			formattedField.Type = "string"
		} else {
			formattedField.Type = field.Type
		}

		if field.Title == "" || field.Name == "" {
			return nil, fmt.Errorf("failed creating log_schema: %w", fmt.Errorf("`title` and `name` fields must be present for each `fields` object"))
		}

		formattedField.Title = field.Title
		formattedField.Name = field.Name

		formattedFields = append(formattedFields, formattedField)
	}

	ls, err := r.client.LogSchema.
		Create().
		SetName(formattedSchemaName).
		SetDescription(logSchema.Description).
		SetTitle(logSchema.Title).
		SetFields(formattedFields).
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

func (r *LogSchemaRepository) GetByTitleOrDescriptionAll(title string, description string) ([]*ent.LogSchema, error) {
	ls, err := r.client.Debug().LogSchema.
		Query().
		Where(logschema.Or(logschema.TitleContains(title), logschema.DescriptionContains(description))).
		Order(ent.Desc(logschema.FieldID)).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_schemas: %w", err)
	}

	return ls, nil
}

func (r *LogSchemaRepository) GetAll() ([]*ent.LogSchema, error) {
	ls, err := r.client.Debug().LogSchema.
		Query().Order(ent.Desc(logschema.FieldID)).All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_schemas: %w", err)
	}

	return ls, nil
}
