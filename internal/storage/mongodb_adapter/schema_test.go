package mongodb_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaStorage(t *testing.T) {
	ctx := context.Background()

	mongoC, client, err := testutil.SetupMongoContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := testutil.TeardownMongoContainer(ctx, mongoC); err != nil {
			log.Fatal(err)
		}
	}()

	schemaStorage := &SchemaStorage{
		ctx:        ctx,
		client:     client,
		collection: client.Database("test_moonlogs").Collection("schemas"),
	}

	t.Run("CreateSchema", func(t *testing.T) {
		schema := entities.Schema{
			Title:         "Test Schema",
			Description:   "A schema for testing",
			Fields:        entities.Fields{{Title: "Field1", Name: "field1"}, {Title: "Field2", Name: "field2"}},
			Kinds:         entities.Kinds{{Title: "Kind1", Name: "kind1"}, {Title: "Kind2", Name: "kind2"}},
			RetentionDays: 30,
			TagID:         1,
		}
		createdSchema, err := schemaStorage.CreateSchema(schema)
		assert.NoError(t, err)
		assert.NotNil(t, createdSchema)
		assert.Equal(t, schema.Title, createdSchema.Title)
		assert.Equal(t, schema.Description, createdSchema.Description)
		assert.Equal(t, schema.Fields, createdSchema.Fields)
		assert.Equal(t, schema.Kinds, createdSchema.Kinds)
	})

	t.Run("GetById", func(t *testing.T) {
		schema, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema 2",
			Description:   "Another schema for testing",
			Fields:        entities.Fields{{Title: "Field3", Name: "field3"}, {Title: "Field4", Name: "field4"}},
			Kinds:         entities.Kinds{{Title: "Kind3", Name: "kind3"}, {Title: "Kind4", Name: "kind4"}},
			RetentionDays: 60,
			TagID:         2,
		})
		assert.NoError(t, err)

		fetchedSchema, err := schemaStorage.GetById(schema.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedSchema)
		assert.Equal(t, schema.Title, fetchedSchema.Title)
		assert.Equal(t, schema.Description, fetchedSchema.Description)
		assert.Equal(t, schema.Fields, fetchedSchema.Fields)
		assert.Equal(t, schema.Kinds, fetchedSchema.Kinds)
	})

	t.Run("UpdateSchemaByID", func(t *testing.T) {
		schema, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema 3",
			Description:   "Yet another schema for testing",
			Fields:        entities.Fields{{Title: "Field5", Name: "field5"}, {Title: "Field6", Name: "field6"}},
			Kinds:         entities.Kinds{{Title: "Kind5", Name: "kind5"}, {Title: "Kind6", Name: "kind6"}},
			RetentionDays: 90,
			TagID:         3,
		})
		assert.NoError(t, err)

		schema.Title = "Updated Test Schema 3"
		schema.Fields = entities.Fields{{Title: "Updated Field5", Name: "updated_field5"}}
		updatedSchema, err := schemaStorage.UpdateSchemaByID(schema.ID, *schema)
		assert.NoError(t, err)
		assert.NotNil(t, updatedSchema)
		assert.Equal(t, "Updated Test Schema 3", updatedSchema.Title)
		assert.Equal(t, schema.Fields, updatedSchema.Fields)
	})

	t.Run("GetByTagID", func(t *testing.T) {
		tagID := 10

		_, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema Tag",
			Description:   "A schema for testing by tag",
			Fields:        entities.Fields{{Title: "Field7", Name: "field7"}},
			Kinds:         entities.Kinds{{Title: "Kind7", Name: "kind7"}},
			RetentionDays: 180,
			TagID:         tagID,
		})
		assert.NoError(t, err)

		_, err = schemaStorage.CreateSchema(entities.Schema{
			Title:         "Another Test Schema Tag",
			Description:   "Another schema for testing by tag",
			Fields:        entities.Fields{{Title: "Field8", Name: "field8"}},
			Kinds:         entities.Kinds{{Title: "Kind8", Name: "kind8"}},
			RetentionDays: 365,
			TagID:         tagID,
		})
		assert.NoError(t, err)

		schemas, err := schemaStorage.GetByTagID(tagID)
		assert.NoError(t, err)
		assert.Len(t, schemas, 2)
	})

	t.Run("GetByName", func(t *testing.T) {
		schema, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema by Name",
			Description:   "A schema for testing by name",
			Name:          "unique-schema-name",
			Fields:        entities.Fields{{Title: "Field9", Name: "field9"}},
			Kinds:         entities.Kinds{{Title: "Kind9", Name: "kind9"}},
			RetentionDays: 90,
		})
		assert.NoError(t, err)

		fetchedSchema, err := schemaStorage.GetByName("unique-schema-name")
		assert.NoError(t, err)
		assert.NotNil(t, fetchedSchema)
		assert.Equal(t, schema.Name, fetchedSchema.Name)
	})

	t.Run("GetSchemasByTitleOrDescription", func(t *testing.T) {
		title := "Partial Title"
		description := "Partial Description"

		_, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Some Partial Title",
			Description:   "Some Partial Description",
			Fields:        entities.Fields{{Title: "Field10", Name: "field10"}},
			Kinds:         entities.Kinds{{Title: "Kind10", Name: "kind10"}},
			RetentionDays: 60,
		})
		assert.NoError(t, err)

		schemas, err := schemaStorage.GetSchemasByTitleOrDescription(title, description)
		assert.NoError(t, err)
		assert.Len(t, schemas, 1)
	})

	t.Run("GetAllSchemas", func(t *testing.T) {
		_, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema for All",
			Description:   "A schema for testing all",
			Fields:        entities.Fields{{Title: "Field11", Name: "field11"}},
			Kinds:         entities.Kinds{{Title: "Kind11", Name: "kind11"}},
			RetentionDays: 150,
		})
		assert.NoError(t, err)

		schemas, err := schemaStorage.GetAllSchemas()
		assert.NoError(t, err)
		assert.True(t, len(schemas) > 0)
	})

	t.Run("DeleteSchemaByID", func(t *testing.T) {
		schema, err := schemaStorage.CreateSchema(entities.Schema{
			Title:         "Test Schema to Delete",
			Description:   "Schema to be deleted",
			Fields:        entities.Fields{{Title: "Field12", Name: "field12"}},
			Kinds:         entities.Kinds{{Title: "Kind12", Name: "kind12"}},
			RetentionDays: 120,
			TagID:         4,
		})
		assert.NoError(t, err)

		err = schemaStorage.DeleteSchemaByID(schema.ID)
		assert.NoError(t, err)

		deletedSchema, err := schemaStorage.GetById(schema.ID)
		assert.NoError(t, err)
		assert.Nil(t, deletedSchema)
	})
}
