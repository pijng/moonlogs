package usecases

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"slices"
	"strings"
)

type SchemaUseCase struct {
	schemaStorage storage.SchemaStorage
}

func NewSchemaUseCase(schemaStorage storage.SchemaStorage) *SchemaUseCase {
	return &SchemaUseCase{schemaStorage: schemaStorage}
}

func (uc *SchemaUseCase) CreateSchema(ctx context.Context, schema entities.Schema) (*entities.Schema, error) {
	existingSchema, err := uc.GetSchemaByName(ctx, normalizeName(schema.Name))
	if err != nil && !errors.Is(err, storage.ErrNotFound) {
		return nil, fmt.Errorf("failed querying schema by name: %w", err)
	}

	// update relevant fields if schema by the given name already exists
	if existingSchema.ID != 0 {
		mergedSchema := mergeSchemaFields(*existingSchema, schema)

		return uc.schemaStorage.UpdateSchemaByID(ctx, existingSchema.ID, mergedSchema)
	}

	if len(schema.Fields) == 0 {
		return nil, fmt.Errorf("failed creating schema: `fields` attribute is required")
	}

	schema.Name = normalizeName(schema.Name)

	var formattedFields entities.Fields

	for _, field := range schema.Fields {
		var formattedField entities.Field

		if field.Title == "" || field.Name == "" {
			return nil, fmt.Errorf("failed creating schema: `title` and `name` attributes must be present for each `fields` object")
		}

		formattedField.Title = field.Title
		formattedField.Name = normalizeName(field.Name)

		formattedFields = append(formattedFields, formattedField)
	}

	schema.Fields = formattedFields

	var formattedKinds entities.Kinds

	for _, kind := range schema.Kinds {
		var formattedKind entities.Kind

		if kind.Title == "" || kind.Name == "" {
			return nil, fmt.Errorf("failed creating schema: `title` and `name` attributes must be present for each `kind` object")
		}

		formattedKind.Title = kind.Title
		formattedKind.Name = normalizeName(kind.Name)

		formattedKinds = append(formattedKinds, formattedKind)
	}

	schema.Kinds = formattedKinds

	return uc.schemaStorage.CreateSchema(ctx, schema)
}

func (uc *SchemaUseCase) UpdateSchemaByID(ctx context.Context, id int, schema entities.Schema) (*entities.Schema, error) {
	var formattedFields entities.Fields

	for _, field := range schema.Fields {
		var formattedField entities.Field

		if field.Title == "" || field.Name == "" {
			return nil, fmt.Errorf("failed creating schema: `title` and `name` attributes must be present for each `fields` object")
		}

		formattedField.Title = field.Title
		formattedField.Name = normalizeName(field.Name)

		formattedFields = append(formattedFields, formattedField)
	}

	schema.Name = normalizeName(schema.Name)
	schema.Fields = formattedFields

	return uc.schemaStorage.UpdateSchemaByID(ctx, id, schema)
}

func (uc *SchemaUseCase) GetAllSchemas(ctx context.Context, user *entities.User) ([]*entities.Schema, error) {
	schemas, err := uc.schemaStorage.GetAllSchemas(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying all schemas")
	}

	if user == nil {
		return nil, fmt.Errorf("cannot get user for tag check")
	}

	if len(user.Tags) > 0 {
		var filteredSchemas []*entities.Schema
		for _, schema := range schemas {
			if schema.TagID != 0 && slices.Contains(user.Tags, schema.TagID) {
				filteredSchemas = append(filteredSchemas, schema)
			}
		}

		return filteredSchemas, nil
	}

	return schemas, nil
}

func (uc *SchemaUseCase) GetSchemaByID(ctx context.Context, id int) (*entities.Schema, error) {
	return uc.schemaStorage.GetById(ctx, id)
}

func (uc *SchemaUseCase) GetSchemaByTagID(ctx context.Context, tagID int) ([]*entities.Schema, error) {
	return uc.schemaStorage.GetByTagID(ctx, tagID)
}

func (uc *SchemaUseCase) GetSchemaByName(ctx context.Context, name string) (*entities.Schema, error) {
	return uc.schemaStorage.GetByName(ctx, name)
}

func (uc *SchemaUseCase) GetSchemasByTitleOrDescription(ctx context.Context, title string, description string) ([]*entities.Schema, error) {
	return uc.schemaStorage.GetSchemasByTitleOrDescription(ctx, title, description)
}

func (uc *SchemaUseCase) DeleteSchemaByID(ctx context.Context, id int) error {
	return uc.schemaStorage.DeleteSchemaByID(ctx, id)
}

func normalizeName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "_")
}

func mergeSchemaFields(existingSchema entities.Schema, newSchema entities.Schema) entities.Schema {
	if newSchema.Title == "" {
		newSchema.Title = existingSchema.Title
	}
	if newSchema.Name == "" {
		newSchema.Name = existingSchema.Name
	}
	if newSchema.Description == "" {
		newSchema.Description = existingSchema.Description
	}
	if len(newSchema.Fields) == 0 {
		newSchema.Fields = existingSchema.Fields
	}
	if newSchema.TagID == 0 {
		newSchema.TagID = existingSchema.TagID
	}

	return newSchema
}
