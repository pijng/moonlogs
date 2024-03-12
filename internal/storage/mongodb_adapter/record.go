package mongodb_adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecordStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewRecordStorage(ctx context.Context) *RecordStorage {
	return &RecordStorage{
		ctx:        ctx,
		collection: persistence.MongoDB().Database("moonlogs").Collection("records"),
	}
}

func (s *RecordStorage) CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error) {
	nextValue, err := getNextSequenceValue(s.ctx, persistence.MongoDB(), "records")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	txn := newrelic.FromContext(s.ctx)
	defer txn.StartSegment("storage.sqlite_adapter.CreateRecord").End()

	formattedQuery := make(map[string]string)
	for k, v := range record.Query {
		var vStr string

		switch v := v.(type) {
		case string:
			vStr = v
		default:
			mv, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("converting query key to string: %w", err)
			}

			vStr = string(mv)
		}

		formattedQuery[k] = vStr
	}

	document := bson.M{
		"id": nextValue, "text": record.Text, "schema_name": record.SchemaName, "schema_id": schemaID, "query": formattedQuery,
		"request": record.Request, "response": record.Response, "kind": record.Kind, "group_hash": groupHash,
		"level": record.Level, "created_at": entities.RecordTime{Time: time.Now()},
	}
	result, err := s.collection.InsertOne(s.ctx, document)
	if err != nil {
		return nil, fmt.Errorf("failed inserting record: %w", err)
	}

	var lr entities.Record
	err = s.collection.FindOne(s.ctx, bson.M{"_id": result.InsertedID}).Decode(&lr)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted record: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordByID(id int) (*entities.Record, error) {
	var lr entities.Record

	err := s.collection.FindOne(s.ctx, bson.M{"id": id}).Decode(&lr)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying record by id: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordsByQuery(record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error) {
	filter := bson.M{}

	if record.ID != 0 {
		filter["schema_id"] = record.SchemaID
	} else {
		filter["schema_name"] = record.SchemaName
	}

	if record.Text != "" {
		filter["text"] = bson.M{"$regex": primitive.Regex{Pattern: record.Text, Options: "i"}}
	}
	if record.Kind != "" {
		filter["kind"] = record.Kind
	}
	if record.Level != "" {
		filter["level"] = record.Level
	}
	if len(record.Query) != 0 {
		filter = qrx.QueryObject(filter, record.Query)
	}
	if from != nil || to != nil {
		filter["created_at"] = bson.M{"$gte": qrx.From(from), "$lte": qrx.To(to)}
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	totalCount, err := s.collection.CountDocuments(s.ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed counting records by query: %w", err)
	}

	cursor, err := s.collection.Find(s.ctx, filter, findOptions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(s.ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(s.ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, 0, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, int(totalCount), nil
}

func (s *RecordStorage) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(s.ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(s.ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(s.ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, nil
}

func (s *RecordStorage) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	filter := bson.M{"schema_name": schemaName, "group_hash": groupHash}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: 1}})

	cursor, err := s.collection.Find(s.ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(s.ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(s.ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, nil
}

func (s *RecordStorage) GetAllRecordsCount() (int, error) {
	totalCount, err := s.collection.CountDocuments(s.ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("counting all records: %w", err)
	}

	return int(totalCount), nil
}

func (s *RecordStorage) FindStaleIDs(schemaID int, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	filter := bson.M{
		"schema_id":  schemaID,
		"created_at": bson.M{"$lte": time.Unix(threshold, 0)},
	}

	rowsCount, err := s.collection.CountDocuments(s.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("counting stale records: %w", err)
	}

	ids := make([]int, 0, rowsCount)

	cursor, err := s.collection.Find(s.ctx, filter, options.Find().SetProjection(bson.M{"id": 1}))
	if err != nil {
		return nil, fmt.Errorf("querying stale record's ids: %w", err)
	}

	for cursor.Next(s.ctx) {
		var id int
		if err := cursor.Decode(&id); err != nil {
			return nil, fmt.Errorf("failed decoding record's id: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (s *RecordStorage) DeleteByIDs(ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	filter := bson.M{"id": bson.M{"$in": ids}}
	_, err := s.collection.DeleteMany(s.ctx, filter)

	return err
}
