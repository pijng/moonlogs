package usecases

import (
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

func (uc *TagUseCase) CreateTag(name string) (*entities.Tag, error) {
	if name == "" {
		return nil, fmt.Errorf("failed creating tag: `name` attribute is required")
	}

	return uc.tagStorage.CreateTag(entities.Tag{Name: name})
}

func (uc *TagUseCase) GetAllTags() ([]*entities.Tag, error) {
	return uc.tagStorage.GetAllTags()
}

func (uc *TagUseCase) DestroyTagByID(id int) error {
	return uc.tagStorage.DestroyTagByID(id)
}

func (uc *TagUseCase) GetTagByID(id int) (*entities.Tag, error) {
	return uc.tagStorage.GetTagByID(id)
}

func (uc *TagUseCase) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	return uc.tagStorage.UpdateTagByID(id, tag)
}
