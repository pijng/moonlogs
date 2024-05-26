package mongodb_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecordStorage(t *testing.T) {
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

	recordStorage := &RecordStorage{
		ctx:        ctx,
		client:     client,
		collection: client.Database("test_moonlogs").Collection("records"),
	}

	t.Run("CreateRecord", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record",
			SchemaName: "Test Schema",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "Test Group",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		createdRecord, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)
		assert.NotNil(t, createdRecord)
		assert.Equal(t, record.Text, createdRecord.Text)
	})

	t.Run("GetRecordByID", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record By ID",
			SchemaName: "Test Schema By ID",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "Test Group",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		createdRecord, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)

		fetchedRecord, err := recordStorage.GetRecordByID(createdRecord.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedRecord)
		assert.Equal(t, createdRecord.Text, fetchedRecord.Text)
	})

	t.Run("GetRecordsByQuery", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record By Query",
			SchemaName: "Test Schema By Query",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "Test Group",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		_, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)

		from := time.Now().Add(-time.Hour)
		to := time.Now()
		records, totalCount, err := recordStorage.GetRecordsByQuery(record, &from, &to, 10, 0)
		assert.NoError(t, err)
		assert.True(t, totalCount > 0)
		assert.NotNil(t, records)
	})

	t.Run("GetAllRecords", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record For All",
			SchemaName: "Test Schema For All",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "Test Group",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		_, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)

		records, err := recordStorage.GetAllRecords(10, 0)
		assert.NoError(t, err)
		assert.True(t, len(records) > 0)
	})

	t.Run("GetRecordsByGroupHash", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record By Group Hash",
			SchemaName: "Test Schema By Group Hash",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "groupHash",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		_, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)

		records, err := recordStorage.GetRecordsByGroupHash(record.SchemaName, "groupHash")
		assert.NoError(t, err)
		assert.True(t, len(records) > 0)
	})

	t.Run("GetAllRecordsCount", func(t *testing.T) {
		count, err := recordStorage.GetAllRecordsCount()
		assert.NoError(t, err)
		assert.True(t, count > 0)
	})

	t.Run("FindStaleIDs", func(t *testing.T) {
		threshold := time.Now().Add(-time.Hour).Unix()
		ids, err := recordStorage.FindStaleIDs(1, threshold)
		assert.NoError(t, err)
		assert.NotNil(t, ids)
	})

	t.Run("DeleteByIDs", func(t *testing.T) {
		record := entities.Record{
			Text:       "Test Record To Delete",
			SchemaName: "Test Schema To Delete",
			SchemaID:   1,
			Query:      entities.JSONMap{"key": "value"},
			Request:    entities.JSONMap{"req": "data"},
			Response:   entities.JSONMap{"res": "data"},
			Kind:       "Test Kind",
			GroupHash:  "Test Group",
			Level:      entities.InfoLevel,
			CreatedAt:  entities.RecordTime{Time: time.Now()},
		}
		createdRecord, err := recordStorage.CreateRecord(record)
		assert.NoError(t, err)

		err = recordStorage.DeleteByIDs([]int{createdRecord.ID})
		assert.NoError(t, err)

		deletedRecord, err := recordStorage.GetRecordByID(createdRecord.ID)
		assert.NoError(t, err)
		assert.Nil(t, deletedRecord)
	})
}
