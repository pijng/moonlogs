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

func (uc *TagUseCase) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	if tag.Name == "" {
		return nil, fmt.Errorf("failed creating tag: `name` attribute is required")
	}

	if tag.ViewOrder == 0 {
		tag.ViewOrder = 1
	}

	return uc.tagStorage.CreateTag(tag)
}

func (uc *TagUseCase) GetAllTags() ([]*entities.Tag, error) {
	return uc.tagStorage.GetAllTags()
}

func (uc *TagUseCase) DeleteTagByID(id int) error {
	return uc.tagStorage.DeleteTagByID(id)
}

func (uc *TagUseCase) GetTagByID(id int) (*entities.Tag, error) {
	return uc.tagStorage.GetTagByID(id)
}

func (uc *TagUseCase) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	return uc.tagStorage.UpdateTagByID(id, tag)
}
