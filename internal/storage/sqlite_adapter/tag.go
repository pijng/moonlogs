package sqlite_adapter

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type TagStorage struct {
	ctx     context.Context
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewTagStorage(ctx context.Context) *TagStorage {
	return &TagStorage{
		ctx:     ctx,
		writeDB: persistence.SqliteWriteDB(),
		readDB:  persistence.SqliteReadDB(),
	}
}

func (s *TagStorage) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO tags (name) VALUES (?);"

	result, err := tx.ExecContext(s.ctx, query, tag.Name)
	if err != nil {
		return nil, fmt.Errorf("failed inserting tag: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving tag last insert id: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	t, err := s.GetTagByID(int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying tag: %w", err)
	}

	return t, nil
}

func (s *TagStorage) GetTagByID(id int) (*entities.Tag, error) {
	query := "SELECT * FROM tags WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var t entities.Tag
	err = row.Scan(&t.ID, &t.Name)
	if err != nil {
		return nil, fmt.Errorf("failed scanning tag: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) DeleteTagByID(id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "DELETE FROM tags WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting tag: %w", err)
	}

	return err
}

func (s *TagStorage) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "UPDATE tags SET name=? WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, tag.Name, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return s.GetTagByID(id)
}

func (s *TagStorage) GetAllTags() ([]*entities.Tag, error) {
	query := "SELECT * FROM tags ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}
	defer rows.Close()

	tags := make([]*entities.Tag, 0)

	for rows.Next() {
		var t entities.Tag

		err = rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return nil, fmt.Errorf("failed scanning tag: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
