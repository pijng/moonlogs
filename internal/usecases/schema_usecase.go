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

	// if schema by the given name alrady exists – update the existing one
	if existingSchema.ID > 0 {
		return uc.UpdateSchemaByID(existingSchema.ID, schema)
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
