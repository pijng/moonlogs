package repositories

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type TagRepository struct {
	ctx  context.Context
	tags *qrx.TableQuerier[entities.Tag]
}

func NewTagRepository(ctx context.Context) *TagRepository {
	return &TagRepository{
		ctx:  ctx,
		tags: qrx.Scan(entities.Tag{}).With(persistence.DB()).From("tags"),
	}
}

func (r *TagRepository) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	u, err := r.tags.Create(r.ctx, map[string]interface{}{
		"name": tag.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating tag: %w", err)
	}

	return u, nil
}

func (r *TagRepository) GetTagByID(id int) (*entities.Tag, error) {
	u, err := r.tags.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying tag: %w", err)
	}

	return u, nil
}

func (r *TagRepository) DestroyTagByID(id int) error {
	_, err := r.tags.DeleteOne(r.ctx, "id=?", id)

	return err
}

func (r *TagRepository) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	data := map[string]interface{}{
		"name": tag.Name,
	}

	u, err := r.tags.Where("id = ?", id).UpdateOne(r.ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return u, nil
}

func (r *TagRepository) GetAllTags() ([]*entities.Tag, error) {
	u, err := r.tags.All(r.ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}

	return u, nil
}
