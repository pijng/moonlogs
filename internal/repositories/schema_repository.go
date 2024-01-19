package repositories

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/persistence"
	"moonlogs/lib/qrx"
)

type SchemaRepository struct {
	ctx     context.Context
	schemas *qrx.TableQuerier[entities.Schema]
}

func NewSchemaRepository(ctx context.Context) *SchemaRepository {
	return &SchemaRepository{
		ctx:     ctx,
		schemas: qrx.Scan(entities.Schema{}).With(persistence.DB()).From("schemas"),
	}
}

func (r *SchemaRepository) CreateSchema(schema entities.Schema) (*entities.Schema, error) {
	s, err := r.schemas.Create(r.ctx, map[string]interface{}{
		"name":           schema.Name,
		"description":    schema.Description,
		"retention_days": schema.RetentionDays,
		"title":          schema.Title,
		"fields":         schema.Fields,
		"kinds":          schema.Kinds,
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating schema: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) UpdateSchemaByID(id int, schema entities.Schema) (*entities.Schema, error) {
	s, err := r.schemas.Where("id = ?", id).UpdateOne(r.ctx, map[string]interface{}{
		"description":    schema.Description,
		"title":          schema.Title,
		"fields":         schema.Fields,
		"retention_days": schema.RetentionDays,
	})

	if err != nil {
		return nil, fmt.Errorf("failed updating schema: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) GetById(id int) (*entities.Schema, error) {
	s, err := r.schemas.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) GetByName(name string) (*entities.Schema, error) {
	s, err := r.schemas.Where("name = ?", name).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) GetSchemasByTitleOrDescription(title string, description string) ([]*entities.Schema, error) {
	s, err := r.schemas.Where("title LIKE ? OR description LIKE ? ORDER BY id DESC", qrx.Contains(title), qrx.Contains(description)).All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) GetAllSchemas() ([]*entities.Schema, error) {
	s, err := r.schemas.All(r.ctx, "ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}

	return s, nil
}

func (r *SchemaRepository) DestroySchemaByID(id int) error {
	_, err := r.schemas.DeleteOne(r.ctx, "id = ?", id)

	return err
}
