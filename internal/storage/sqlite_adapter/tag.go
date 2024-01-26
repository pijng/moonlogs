package sqlite_adapter

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type TagStorage struct {
	ctx  context.Context
	tags *qrx.TableQuerier[entities.Tag]
}

func NewTagStorage(ctx context.Context) *TagStorage {
	return &TagStorage{
		ctx:  ctx,
		tags: qrx.Scan(entities.Tag{}).With(persistence.DB()).From("tags"),
	}
}

func (r *TagStorage) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	u, err := r.tags.Create(r.ctx, map[string]interface{}{
		"name": tag.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating tag: %w", err)
	}

	return u, nil
}

func (s *TagStorage) GetTagByID(id int) (*entities.Tag, error) {
	u, err := s.tags.Where("id = ?", id).First(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying tag: %w", err)
	}

	return u, nil
}

func (r *TagStorage) DestroyTagByID(id int) error {
	_, err := r.tags.DeleteOne(r.ctx, "id=?", id)

	return err
}

func (r *TagStorage) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	data := map[string]interface{}{
		"name": tag.Name,
	}

	u, err := r.tags.Where("id = ?", id).UpdateOne(r.ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return u, nil
}

func (r *TagStorage) GetAllTags() ([]*entities.Tag, error) {
	u, err := r.tags.All(r.ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}

	return u, nil
}
