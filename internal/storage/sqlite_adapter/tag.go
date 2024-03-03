package sqlite_adapter

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/persistence"
)

type TagStorage struct {
	ctx context.Context
	db  *sql.DB
}

func NewTagStorage(ctx context.Context) *TagStorage {
	return &TagStorage{
		ctx: ctx,
		db:  persistence.DB(),
	}
}

func (s *TagStorage) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	query := "INSERT INTO tags (name) VALUES (?);"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(s.ctx, tag.Name)

	if err != nil {
		return nil, fmt.Errorf("failed inserting tag: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving tag last insert id: %w", err)
	}

	t, err := s.GetTagByID(int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying tag: %w", err)
	}

	return t, nil
}

func (s *TagStorage) GetTagByID(id int) (*entities.Tag, error) {
	query := "SELECT * FROM tags WHERE id=? LIMIT 1;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
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
	query := "DELETE FROM tags WHERE id=?"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(s.ctx, id)
	if err != nil {
		return fmt.Errorf("failed deleting tag: %w", err)
	}

	return err
}

func (s *TagStorage) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	query := "UPDATE tags SET name=? WHERE id=?"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(s.ctx, tag.Name, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return s.GetTagByID(id)
}

func (s *TagStorage) GetAllTags() ([]*entities.Tag, error) {
	query := "SELECT * FROM tags ORDER BY id DESC;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
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
