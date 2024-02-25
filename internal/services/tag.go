package services

import (
	"fmt"
	"moonlogs/internal/usecases"
)

type TagService struct {
	tagUseCase    *usecases.TagUseCase
	schemaUseCase *usecases.SchemaUseCase
	userUseCase   *usecases.UserUseCase
}

func NewTagService(tuc *usecases.TagUseCase, suc *usecases.SchemaUseCase, uuc *usecases.UserUseCase) *TagService {
	return &TagService{tagUseCase: tuc, schemaUseCase: suc, userUseCase: uuc}
}

func (m *TagService) DestroyTagByID(id int) error {
	err := m.tagUseCase.DeleteTagByID(id)
	if err != nil {
		return fmt.Errorf("destroying tag by id: %w", err)
	}

	schemas, err := m.schemaUseCase.GetSchemaByTagID(id)
	if err != nil {
		return fmt.Errorf("querying schemas by tag_id: %w", err)
	}

	for _, schema := range schemas {
		schema.TagID = 0

		_, err = m.schemaUseCase.UpdateSchemaByID(schema.ID, *schema)
		if err != nil {
			return fmt.Errorf("updating schema tag_id: %w", err)
		}
	}

	users, err := m.userUseCase.GetUsersByTagID(id)
	if err != nil {
		return fmt.Errorf("querying users by tag_id: %w", err)
	}

	for _, user := range users {
		tagIDS := make([]int, len(user.Tags)-1)
		for _, tagID := range user.Tags {
			if tagID != id {
				tagIDS = append(tagIDS, tagID)
			}
		}

		user.Tags = tagIDS

		_, err = m.userUseCase.UpdateUserByID(user.ID, *user)
		if err != nil {
			return fmt.Errorf("updating user tag_id: %w", err)
		}
	}

	return nil
}
