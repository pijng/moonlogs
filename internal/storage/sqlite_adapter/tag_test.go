package sqlite_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"moonlogs/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagStorage(t *testing.T) {
	ctx := context.Background()

	writeDB, readDB, err := testutil.SetupSqlite()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := testutil.TeardownSqlite(); err != nil {
			log.Fatal(err)
		}
	}()

	tagStorage := NewTagStorage(readDB, writeDB)

	t.Run("CreateTag", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag", ViewOrder: 1}
		createdTag, err := tagStorage.CreateTag(ctx, tag)
		assert.NoError(t, err)
		assert.NotNil(t, createdTag)
		assert.NotNil(t, createdTag.ID)
		assert.Equal(t, "TestTag", createdTag.Name)
	})

	t.Run("GetTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag2", ViewOrder: 2}
		createdTag, err := tagStorage.CreateTag(ctx, tag)
		assert.NoError(t, err)

		foundTag, err := tagStorage.GetTagByID(ctx, createdTag.ID)
		assert.NoError(t, err)
		assert.NotNil(t, createdTag.ID)
		assert.Equal(t, "TestTag2", foundTag.Name)
	})

	t.Run("DeleteTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag3", ViewOrder: 3}
		createdTag, err := tagStorage.CreateTag(ctx, tag)
		assert.NoError(t, err)

		err = tagStorage.DeleteTagByID(ctx, createdTag.ID)
		assert.NoError(t, err)

		foundTag, err := tagStorage.GetTagByID(ctx, createdTag.ID)
		assert.Nil(t, foundTag)
		assert.ErrorIs(t, err, storage.ErrNotFound)
	})

	t.Run("UpdateTagByID", func(t *testing.T) {
		tag := entities.Tag{Name: "TestTag4", ViewOrder: 4}
		createdTag, err := tagStorage.CreateTag(ctx, tag)
		assert.NoError(t, err)

		updatedTag := entities.Tag{Name: "UpdatedTestTag4", ViewOrder: 5}
		updated, err := tagStorage.UpdateTagByID(ctx, createdTag.ID, updatedTag)
		assert.NoError(t, err)
		assert.Equal(t, "UpdatedTestTag4", updated.Name)
		assert.Equal(t, 5, updated.ViewOrder)
	})

	t.Run("GetAllTags", func(t *testing.T) {
		tags, err := tagStorage.GetAllTags(ctx)
		assert.NoError(t, err)
		assert.True(t, len(tags) > 0)
	})
}
