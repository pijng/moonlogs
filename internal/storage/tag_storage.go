package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type TagStorage interface {
	CreateTag(ctx context.Context, tag entities.Tag) (*entities.Tag, error)
	DeleteTagByID(ctx context.Context, id int) error
	GetAllTags(ctx context.Context) ([]*entities.Tag, error)
	GetTagByID(ctx context.Context, id int) (*entities.Tag, error)
	UpdateTagByID(ctx context.Context, id int, tag entities.Tag) (*entities.Tag, error)
}
