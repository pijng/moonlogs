package usecases

import (
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"slices"
	"strings"
)

var cachedSchemas = make(map[string]*entities.Schema)

type SchemaUseCase struct {
	schemaStorage storage.SchemaStorage
}

func NewSchemaUseCase(schemaStorage storage.SchemaStorage) *SchemaUseCase {
	return &SchemaUseCase{schemaStorage: schemaStorage}
}

func (uc *SchemaUseCase) CreateSchema(schema entities.Schema) (*entities.Schema, error) {
	existingSchema, err := uc.GetSchemaByName(normalizeName(schema.Name))
	if err != nil {
		return nil, fmt.Errorf("failed querying schema by name: %w", err)
	}

	// update relevant fields if schema by the given name already exists
	if existingSchema.ID != 0 {
		mergedSchema := mergeSchemaFields(*existingSchema, schema)

		return uc.schemaStorage.UpdateSchemaByID(existingSchema.ID, mergedSchema)
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

	return uc.schemaStorage.CreateSchema(schema)
}

func (uc *SchemaUseCase) UpdateSchemaByID(id int, schema entities.Schema) (*entities.Schema, error) {
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

	return uc.schemaStorage.UpdateSchemaByID(id, schema)
}

func (uc *SchemaUseCase) GetAllSchemas(user *entities.User) ([]*entities.Schema, error) {
	schemas, err := uc.schemaStorage.GetAllSchemas()
	if err != nil {
		return nil, fmt.Errorf("failed querying all schemas")
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

func (uc *SchemaUseCase) GetSchemaByID(id int) (*entities.Schema, error) {
	return uc.schemaStorage.GetById(id)
}

func (uc *SchemaUseCase) GetSchemaByTagID(tagID int) ([]*entities.Schema, error) {
	return uc.schemaStorage.GetByTagID(tagID)
}

func (uc *SchemaUseCase) GetSchemaByName(name string) (*entities.Schema, error) {
	schema, ok := cachedSchemas[name]
	if ok {
		return schema, nil
	}

	schema, err := uc.schemaStorage.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("getting schema by name: %w", err)
	}

	cachedSchemas[schema.Name] = schema

	return schema, nil
}

func (uc *SchemaUseCase) GetSchemasByTitleOrDescription(title string, description string) ([]*entities.Schema, error) {
	return uc.schemaStorage.GetSchemasByTitleOrDescription(title, description)
}

func (uc *SchemaUseCase) DeleteSchemaByID(id int) error {
	return uc.schemaStorage.DeleteSchemaByID(id)
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
