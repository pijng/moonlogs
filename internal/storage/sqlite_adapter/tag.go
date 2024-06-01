package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
)

type TagStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewTagStorage(readDB *sql.DB, writeDB *sql.DB) *TagStorage {
	return &TagStorage{
		writeDB: readDB,
		readDB:  writeDB,
	}
}

func (s *TagStorage) CreateTag(ctx context.Context, tag entities.Tag) (*entities.Tag, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO tags (name, view_order) VALUES (?, ?);"

	result, err := tx.ExecContext(ctx, query, tag.Name, tag.ViewOrder)
	if err != nil {
		return nil, fmt.Errorf("failed inserting tag: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving tag last insert id: %w", err)
	}

	t, err := s.GetTagByID(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying tag: %w", err)
	}

	return t, nil
}

func (s *TagStorage) GetTagByID(ctx context.Context, id int) (*entities.Tag, error) {
	query := "SELECT * FROM tags WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var t entities.Tag
	err = row.Scan(&t.ID, &t.Name, &t.ViewOrder)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed scanning tag: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) DeleteTagByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM tags WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting tag: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *TagStorage) UpdateTagByID(ctx context.Context, id int, tag entities.Tag) (*entities.Tag, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE tags SET name=?, view_order=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, tag.Name, tag.ViewOrder, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetTagByID(ctx, id)
}

func (s *TagStorage) GetAllTags(ctx context.Context) ([]*entities.Tag, error) {
	query := "SELECT * FROM tags ORDER BY view_order ASC, id DESC;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}
	defer rows.Close()

	tags := make([]*entities.Tag, 0)

	for rows.Next() {
		var t entities.Tag

		err = rows.Scan(&t.ID, &t.Name, &t.ViewOrder)
		if err != nil {
			return nil, fmt.Errorf("failed scanning tag: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
