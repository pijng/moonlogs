package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type SchemaStorage struct {
	ctx     context.Context
	schemas *qrx.TableQuerier[entities.Schema]
	db      *sql.DB
}

func NewSchemaStorage(ctx context.Context) *SchemaStorage {
	return &SchemaStorage{
		ctx:     ctx,
		schemas: qrx.Scan(entities.Schema{}).With(persistence.DB()).From("schemas"),
		db:      persistence.DB(),
	}
}

func (s *SchemaStorage) CreateSchema(schema entities.Schema) (*entities.Schema, error) {
	sm, err := s.schemas.Create(s.ctx, map[string]interface{}{
		"name":           schema.Name,
		"description":    schema.Description,
		"retention_days": schema.RetentionDays,
		"title":          schema.Title,
		"fields":         schema.Fields,
		"kinds":          schema.Kinds,
		"tag_id":         schema.TagID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) UpdateSchemaByID(id int, schema entities.Schema) (*entities.Schema, error) {
	sm, err := s.schemas.Where("id = ?", id).UpdateOne(s.ctx, map[string]interface{}{
		"description":    schema.Description,
		"title":          schema.Title,
		"fields":         schema.Fields,
		"kinds":          schema.Kinds,
		"retention_days": schema.RetentionDays,
		"tag_id":         schema.TagID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed updating schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) GetById(id int) (*entities.Schema, error) {
	sm, err := s.schemas.Where("id = ?", id).First(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) GetByTagID(id int) ([]*entities.Schema, error) {
	sm, err := s.schemas.Where("tag_id = ?", id).All(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) GetByName(name string) (*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE name=? LIMIT 1;"
	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return &entities.Schema{}, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	row := stmt.QueryRowContext(s.ctx, name)

	var sm entities.Schema
	err = row.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
	if errors.Is(err, sql.ErrNoRows) {
		return &entities.Schema{}, nil
	}

	if err != nil {
		return &entities.Schema{}, fmt.Errorf("failed scanning schema: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetSchemasByTitleOrDescription(title string, description string) ([]*entities.Schema, error) {
	sm, err := s.schemas.Where("title LIKE ? OR description LIKE ? ORDER BY id DESC", qrx.Contains(title), qrx.Contains(description)).All(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) GetAllSchemas() ([]*entities.Schema, error) {
	sm, err := s.schemas.All(s.ctx, "ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) DeleteSchemaByID(id int) error {
	_, err := s.schemas.DeleteOne(s.ctx, "id = ?", id)

	return err
}
