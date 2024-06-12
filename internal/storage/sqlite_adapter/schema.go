package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
)

type SchemaStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewSchemaStorage(readDB *sql.DB, writeDB *sql.DB) *SchemaStorage {
	return &SchemaStorage{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (s *SchemaStorage) CreateSchema(ctx context.Context, schema entities.Schema) (*entities.Schema, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO schemas (name, description, retention_days, title, fields, kinds, tag_id) VALUES (?,?,?,?,?,?,?);"

	result, err := tx.ExecContext(ctx, query, schema.Name, schema.Description, schema.RetentionDays, schema.Title,
		schema.Fields, schema.Kinds, schema.TagID)

	if err != nil {
		return nil, fmt.Errorf("failed inserting schemas: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving schema last insert id: %w", err)
	}

	sm, err := s.GetById(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) UpdateSchemaByID(ctx context.Context, id int, schema entities.Schema) (*entities.Schema, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE schemas SET description=?, title=?, fields=?, kinds=?, retention_days=?, tag_id=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, schema.Description, schema.Title, schema.Fields, schema.Kinds, schema.RetentionDays, schema.TagID, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating schema: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetById(ctx, id)
}

func (s *SchemaStorage) GetById(ctx context.Context, id int) (*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var sm entities.Schema
	err = row.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning schema: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetByTagID(ctx context.Context, tagID int) ([]*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE tag_id=?;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, tagID)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer rows.Close()

	schemas := make([]*entities.Schema, 0)

	for rows.Next() {
		var sm entities.Schema

		err = rows.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
		if err != nil {
			return nil, fmt.Errorf("failed scanning schema: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) GetByName(ctx context.Context, name string) (*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE name=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return &entities.Schema{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, name)

	var sm entities.Schema
	err = row.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return &entities.Schema{}, fmt.Errorf("failed scanning schema: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetSchemasByTitleOrDescription(ctx context.Context, title string, description string) ([]*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE title LIKE ? OR description like ? ORDER BY id desc;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, qrx.Contains(title), qrx.Contains(description))
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer rows.Close()

	schemas := make([]*entities.Schema, 0)

	for rows.Next() {
		var sm entities.Schema

		err = rows.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
		if err != nil {
			return nil, fmt.Errorf("failed scanning schemas: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) GetAllSchemas(ctx context.Context) ([]*entities.Schema, error) {
	query := "SELECT * FROM schemas ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer rows.Close()

	schemas := make([]*entities.Schema, 0)

	for rows.Next() {
		var sm entities.Schema

		err = rows.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
		if err != nil {
			return nil, fmt.Errorf("failed scanning schema: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) DeleteSchemaByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM schemas WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting schema: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
