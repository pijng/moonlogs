package usecases

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
)

type TagUseCase struct {
	tagStorage storage.TagStorage
}

func NewTagUseCase(tagStorage storage.TagStorage) *TagUseCase {
	return &TagUseCase{tagStorage: tagStorage}
}

func (uc *TagUseCase) CreateTag(ctx context.Context, tag entities.Tag) (*entities.Tag, error) {
	if tag.Name == "" {
		return nil, fmt.Errorf("failed creating tag: `name` attribute is required")
	}

	if tag.ViewOrder == 0 {
		tag.ViewOrder = 1
	}

	return uc.tagStorage.CreateTag(ctx, tag)
}

func (uc *TagUseCase) GetAllTags(ctx context.Context) ([]*entities.Tag, error) {
	return uc.tagStorage.GetAllTags(ctx)
}

func (uc *TagUseCase) DeleteTagByID(ctx context.Context, id int) error {
	return uc.tagStorage.DeleteTagByID(ctx, id)
}

func (uc *TagUseCase) GetTagByID(ctx context.Context, id int) (*entities.Tag, error) {
	return uc.tagStorage.GetTagByID(ctx, id)
}

func (uc *TagUseCase) UpdateTagByID(ctx context.Context, id int, tag entities.Tag) (*entities.Tag, error) {
	return uc.tagStorage.UpdateTagByID(ctx, id, tag)
}
