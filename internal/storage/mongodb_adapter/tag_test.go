package mongodb_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagStorage(t *testing.T) {
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

	tagStorage := &TagStorage{
		ctx:        ctx,
		client:     client,
		collection: client.Database("test_moonlogs").Collection("tags"),
	}

	t.Run("CreateTag", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag", ViewOrder: 1}
		createdTag, err := tagStorage.CreateTag(tag)
		assert.NoError(t, err)
		assert.NotNil(t, createdTag)
		assert.NotNil(t, createdTag.ID)
		assert.Equal(t, "TestTag", createdTag.Name)
	})

	t.Run("GetTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag2", ViewOrder: 2}
		createdTag, err := tagStorage.CreateTag(tag)
		assert.NoError(t, err)

		foundTag, err := tagStorage.GetTagByID(createdTag.ID)
		assert.NoError(t, err)
		assert.NotNil(t, createdTag.ID)
		assert.Equal(t, "TestTag2", foundTag.Name)
	})

	t.Run("DeleteTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag3", ViewOrder: 3}
		createdTag, err := tagStorage.CreateTag(tag)
		assert.NoError(t, err)

		err = tagStorage.DeleteTagByID(createdTag.ID)
		assert.NoError(t, err)

		foundTag, err := tagStorage.GetTagByID(createdTag.ID)
		assert.Nil(t, foundTag)
		assert.NoError(t, err)
	})

	t.Run("UpdateTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag4", ViewOrder: 4}
		createdTag, err := tagStorage.CreateTag(tag)
		assert.NoError(t, err)

		updatedTag := entities.Tag{Name: "UpdatedTestTag4", ViewOrder: 5}
		updated, err := tagStorage.UpdateTagByID(createdTag.ID, updatedTag)
		assert.NoError(t, err)
		assert.Equal(t, "UpdatedTestTag4", updated.Name)
		assert.Equal(t, 5, updated.ViewOrder)
	})

	t.Run("GetAllTags", func(t *testing.T) {
		tags, err := tagStorage.GetAllTags()
		assert.NoError(t, err)
		assert.True(t, len(tags) > 0)
	})
}
