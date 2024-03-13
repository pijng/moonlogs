package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type SchemaStorage struct {
	ctx     context.Context
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewSchemaStorage(ctx context.Context) *SchemaStorage {
	return &SchemaStorage{
		ctx:     ctx,
		readDB:  persistence.SqliteReadDB(),
		writeDB: persistence.SqliteWriteDB(),
	}
}

func (s *SchemaStorage) CreateSchema(schema entities.Schema) (*entities.Schema, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO schemas (name, description, retention_days, title, fields, kinds, tag_id) VALUES (?,?,?,?,?,?,?);"

	result, err := tx.ExecContext(s.ctx, query, schema.Name, schema.Description, schema.RetentionDays, schema.Title,
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

	sm, err := s.GetById(int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying schema: %w", err)
	}

	return sm, nil
}

func (s *SchemaStorage) UpdateSchemaByID(id int, schema entities.Schema) (*entities.Schema, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "UPDATE schemas SET description=?, title=?, fields=?, kinds=?, retention_days=?, tag_id=? WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, schema.Description, schema.Title, schema.Fields, schema.Kinds, schema.RetentionDays, schema.TagID, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating schema: %w", err)
	}

	return s.GetById(id)
}

func (s *SchemaStorage) GetById(id int) (*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var sm entities.Schema
	err = row.Scan(&sm.ID, &sm.Title, &sm.Description, &sm.Name, &sm.Fields, &sm.Kinds, &sm.TagID, &sm.RetentionDays)
	if err != nil {
		return nil, fmt.Errorf("failed scanning schema: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetByTagID(tagID int) ([]*entities.Schema, error) {
	query := "SELECT * FROM schemas WHERE tag_id=?;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx, tagID)
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

func (s *SchemaStorage) GetByName(name string) (*entities.Schema, error) {
	txn := newrelic.FromContext(s.ctx)
	defer txn.StartSegment("storage.sqlite_adapter.GetByName").End()

	query := "SELECT * FROM schemas WHERE name=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return &entities.Schema{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

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
	query := "SELECT * FROM schemas WHERE title LIKE ? OR description lile ? ORDER BY id desc;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx, qrx.Contains(title), qrx.Contains(description))
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

func (s *SchemaStorage) GetAllSchemas() ([]*entities.Schema, error) {
	query := "SELECT * FROM schemas ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx)
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

func (s *SchemaStorage) DeleteSchemaByID(id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "DELETE FROM schemas WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting schema: %w", err)
	}

	return err
}
