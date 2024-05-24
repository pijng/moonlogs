package sqlite_adapter

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type ActionStorage struct {
	ctx     context.Context
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewActionStorage(ctx context.Context) *ActionStorage {
	return &ActionStorage{
		ctx:     ctx,
		writeDB: persistence.SqliteWriteDB(),
		readDB:  persistence.SqliteReadDB(),
	}
}

func (s *ActionStorage) CreateAction(action entities.Action) (*entities.Action, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO actions (name, pattern, method, conditions, schema_ids, disabled) VALUES (?, ?, ?, ?, ?, ?);"

	result, err := tx.ExecContext(s.ctx, query, action.Name, action.Pattern, action.Method, action.Conditions, action.SchemaIDs, action.Disabled)
	if err != nil {
		return nil, fmt.Errorf("failed inserting action: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving action last insert id: %w", err)
	}

	t, err := s.GetActionByID(int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying action: %w", err)
	}

	return t, nil
}

func (s *ActionStorage) GetActionByID(id int) (*entities.Action, error) {
	query := "SELECT id, name, pattern, method, conditions, schema_ids, disabled FROM actions WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var a entities.Action
	err = row.Scan(&a.ID, &a.Name, &a.Pattern, &a.Method, &a.Conditions, &a.SchemaIDs, &a.Disabled)
	if err != nil {
		return nil, fmt.Errorf("failed scanning action: %w", err)
	}

	return &a, nil
}

func (s *ActionStorage) DeleteActionByID(id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "DELETE FROM actions WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting action: %w", err)
	}

	return err
}

func (s *ActionStorage) UpdateActionByID(id int, action entities.Action) (*entities.Action, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "UPDATE actions SET name=?, pattern=?, method=?, conditions=?, disabled=?, schema_ids=? WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, action.Name, action.Pattern, action.Method, action.Conditions, action.Disabled, action.SchemaIDs, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating action: %w", err)
	}

	return s.GetActionByID(id)
}

func (s *ActionStorage) GetAllActions() ([]*entities.Action, error) {
	query := "SELECT id, name, pattern, method, conditions, schema_ids, disabled FROM actions ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying actions: %w", err)
	}
	defer rows.Close()

	actions := make([]*entities.Action, 0)

	for rows.Next() {
		var a entities.Action

		err = rows.Scan(&a.ID, &a.Name, &a.Pattern, &a.Method, &a.Conditions, &a.SchemaIDs, &a.Disabled)
		if err != nil {
			return nil, fmt.Errorf("failed scanning action: %w", err)
		}

		actions = append(actions, &a)
	}

	return actions, nil
}
