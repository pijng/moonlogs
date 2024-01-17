package usecases

import (
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
)

type TagUseCase struct {
	tagRepository *repositories.TagRepository
}

func NewTagUseCase(tagRepository *repositories.TagRepository) *TagUseCase {
	return &TagUseCase{tagRepository: tagRepository}
}

func (uc *TagUseCase) CreateTag(name string) (*entities.Tag, error) {
	if name == "" {
		return nil, fmt.Errorf("failed creating tag: `name` attribute is required")
	}

	return uc.tagRepository.CreateTag(entities.Tag{Name: name})
}

func (uc *TagUseCase) GetAllTags() ([]*entities.Tag, error) {
	return uc.tagRepository.GetAllTags()
}

func (uc *TagUseCase) DestroyTagByID(id int) error {
	return uc.tagRepository.DestroyTagByID(id)
}

func (uc *TagUseCase) GetTagByID(id int) (*entities.Tag, error) {
	return uc.tagRepository.GetTagByID(id)
}

func (uc *TagUseCase) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	return uc.tagRepository.UpdateTagByID(id, tag)
}
