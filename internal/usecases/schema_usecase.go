package usecases

import (
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"strings"
)

type SchemaUseCase struct {
	schemaRepository *repositories.SchemaRepository
}

func NewSchemaUseCase(schemaRepository *repositories.SchemaRepository) *SchemaUseCase {
	return &SchemaUseCase{schemaRepository: schemaRepository}
}

func (uc *SchemaUseCase) CreateSchema(schema entities.Schema) (*entities.Schema, error) {
	existingSchema, err := uc.GetSchemaByName(normalizeName(schema.Name))
	if err != nil {
		return nil, fmt.Errorf("failed querying schema by name: %w", err)
	}

	// update relevant fields if schema by the given name already exists
	if existingSchema.ID != 0 {
		mergedSchema := mergeSchemaFields(*existingSchema, schema)

		return uc.schemaRepository.UpdateSchemaByID(existingSchema.ID, mergedSchema)
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

	return uc.schemaRepository.CreateSchema(schema)
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

	return uc.schemaRepository.UpdateSchemaByID(id, schema)
}

func (uc *SchemaUseCase) GetAllSchemas() ([]*entities.Schema, error) {
	return uc.schemaRepository.GetAllSchemas()
}

func (uc *SchemaUseCase) GetSchemaByID(id int) (*entities.Schema, error) {
	return uc.schemaRepository.GetById(id)
}

func (uc *SchemaUseCase) GetSchemaByName(name string) (*entities.Schema, error) {
	return uc.schemaRepository.GetByName(name)
}

func (uc *SchemaUseCase) GetSchemasByTitleOrDescription(title string, description string) ([]*entities.Schema, error) {
	return uc.schemaRepository.GetSchemasByTitleOrDescription(title, description)
}

func (uc *SchemaUseCase) DestroySchemaByID(id int) error {
	return uc.schemaRepository.DestroySchemaByID(id)
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
	if len(newSchema.Kinds) == 0 {
		newSchema.Kinds = existingSchema.Kinds
	}
	if len(newSchema.Fields) == 0 {
		newSchema.Fields = existingSchema.Fields
	}
	if newSchema.RetentionTime == 0 {
		newSchema.RetentionTime = existingSchema.RetentionTime
	}
	if len(newSchema.Tags) == 0 {
		newSchema.Tags = existingSchema.Tags
	}

	return newSchema
}
