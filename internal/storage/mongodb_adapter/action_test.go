package mongodb_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActionStorage(t *testing.T) {
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

	actionStorage := &ActionStorage{
		ctx:        ctx,
		client:     client,
		collection: client.Database("test_moonlogs").Collection("actions"),
	}

	t.Run("CreateAction", func(t *testing.T) {
		action := entities.Action{
			Name:       "Test Action",
			Pattern:    "/test",
			Method:     string(entities.GETActionMethod),
			Conditions: entities.Conditions{{Attribute: "attr", Operation: "op", Value: "val"}},
			SchemaIDs:  entities.SchemaIDs{1, 2, 3},
			Disabled:   false,
		}
		createdAction, err := actionStorage.CreateAction(action)
		assert.NoError(t, err)
		assert.NotNil(t, createdAction)
		assert.Equal(t, action.Name, createdAction.Name)
	})

	t.Run("GetActionByID", func(t *testing.T) {
		action := entities.Action{
			Name:       "Test Action By ID",
			Pattern:    "/test/id",
			Method:     string(entities.GETActionMethod),
			Conditions: entities.Conditions{{Attribute: "attr", Operation: "op", Value: "val"}},
			SchemaIDs:  entities.SchemaIDs{1, 2, 3},
			Disabled:   false,
		}
		createdAction, err := actionStorage.CreateAction(action)
		assert.NoError(t, err)

		fetchedAction, err := actionStorage.GetActionByID(createdAction.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedAction)
		assert.Equal(t, createdAction.Name, fetchedAction.Name)
	})

	t.Run("DeleteActionByID", func(t *testing.T) {
		action := entities.Action{
			Name:       "Test Action To Delete",
			Pattern:    "/test/delete",
			Method:     string(entities.GETActionMethod),
			Conditions: entities.Conditions{{Attribute: "attr", Operation: "op", Value: "val"}},
			SchemaIDs:  entities.SchemaIDs{1, 2, 3},
			Disabled:   false,
		}
		createdAction, err := actionStorage.CreateAction(action)
		assert.NoError(t, err)

		err = actionStorage.DeleteActionByID(createdAction.ID)
		assert.NoError(t, err)

		deletedAction, err := actionStorage.GetActionByID(createdAction.ID)
		assert.NoError(t, err)
		assert.Nil(t, deletedAction)
	})

	t.Run("UpdateActionByID", func(t *testing.T) {
		action := entities.Action{
			Name:       "Test Action To Update",
			Pattern:    "/test/update",
			Method:     string(entities.GETActionMethod),
			Conditions: entities.Conditions{{Attribute: "attr", Operation: "op", Value: "val"}},
			SchemaIDs:  entities.SchemaIDs{1, 2, 3},
			Disabled:   false,
		}
		createdAction, err := actionStorage.CreateAction(action)
		assert.NoError(t, err)

		updatedData := entities.Action{
			Name:       "Updated Action Name",
			Pattern:    "/updated",
			Method:     string(entities.GETActionMethod),
			Conditions: entities.Conditions{{Attribute: "new_attr", Operation: "new_op", Value: "new_val"}},
			SchemaIDs:  entities.SchemaIDs{4, 5, 6},
			Disabled:   true,
		}
		updatedAction, err := actionStorage.UpdateActionByID(createdAction.ID, updatedData)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAction)
		assert.Equal(t, updatedData.Name, updatedAction.Name)
		assert.Equal(t, updatedData.Pattern, updatedAction.Pattern)
	})

	t.Run("GetAllActions", func(t *testing.T) {
		actions, err := actionStorage.GetAllActions()
		assert.NoError(t, err)
		assert.True(t, len(actions) > 0)
	})
}
