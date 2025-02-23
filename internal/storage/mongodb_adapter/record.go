package mongodb_adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultSchemaGroupField = "schema_name"

type RecordStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRecordStorage(db *mongo.Database) *RecordStorage {
	return &RecordStorage{
		db:         db,
		collection: db.Collection("records"),
	}
}

func (s *RecordStorage) CreateRecord(ctx context.Context, record entities.Record) (*entities.Record, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "records")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

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
		"id": nextValue, "text": record.Text, "schema_name": record.SchemaName, "schema_id": record.SchemaID, "query": formattedQuery,
		"request": record.Request, "response": record.Response, "changes": record.Changes, "kind": record.Kind, "group_hash": record.GroupHash,
		"level": record.Level, "created_at": record.CreatedAt, "old_value": record.OldValue, "new_value": record.NewValue,
	}
	result, err := s.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("failed inserting record: %w", err)
	}

	var lr entities.Record
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&lr)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying inserted record: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordByID(ctx context.Context, id int) (*entities.Record, error) {
	var lr entities.Record

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&lr)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying record by id: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordsByQuery(ctx context.Context, record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error) {
	filter := bson.M{}

	if record.ID != 0 {
		filter["schema_id"] = record.SchemaID
	}
	if record.SchemaName != "" {
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
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	totalCount, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed counting records by query: %w", err)
	}

	cursor, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, 0, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, int(totalCount), nil
}

func (s *RecordStorage) AggregateRecords(ctx context.Context, recordsFilter entities.RecordFilter, aggregation entities.RecordAggregation) ([]*entities.AggregationGroup, error) {
	filter := bson.M{}

	if len(recordsFilter.SchemaIDs) > 0 {
		filter["schema_id"] = bson.M{
			"$in": recordsFilter.SchemaIDs,
		}
	}

	if len(recordsFilter.SchemaKinds) > 0 {
		filter["kind"] = bson.M{
			"$in": recordsFilter.SchemaKinds,
		}
	}

	if len(recordsFilter.Level) > 0 {
		filter["level"] = recordsFilter.Level
	}

	filter["created_at"] = bson.M{"$gte": qrx.From(&recordsFilter.From), "$lte": qrx.To(&recordsFilter.To)}

	if len(recordsFilter.SchemaFields) > 0 {
		for _, field := range recordsFilter.SchemaFields {
			key := "query." + field
			filter[key] = bson.M{"$exists": true}
		}
	}

	idGroups := bson.M{}
	for _, field := range aggregation.GroupBy {
		var val string

		// special case for default schema_name grouping
		if field == defaultSchemaGroupField {
			val = field
		} else {
			val = "query." + field
		}

		idGroups[field] = "$" + val
	}

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: filter},
		},
		{
			{
				Key: "$group",
				Value: bson.M{
					"_id":   idGroups,
					"count": bson.M{"$sum": 1},
				},
			},
		},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed aggregating records: %w", err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed decoding aggregated records: %w", err)
	}

	aggregationGroups := make([]*entities.AggregationGroup, 0)
	for _, result := range results {
		keys := make(entities.JSONMap[any])
		if id, ok := result["_id"].(bson.M); ok {
			for key, value := range id {
				if strVal, isString := value.(string); isString {
					keys[key] = strVal
				} else {
					keys[key] = fmt.Sprintf("%v", value)
				}
			}
		}

		aggregationGroups = append(aggregationGroups, &entities.AggregationGroup{Keys: keys, Count: result["count"].(int32)})
	}

	return aggregationGroups, nil
}

func (s *RecordStorage) GetAllRecords(ctx context.Context, limit int, offset int) ([]*entities.Record, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, nil
}

func (s *RecordStorage) GetRecordsByGroupHash(ctx context.Context, schemaName string, groupHash string) ([]*entities.Record, error) {
	filter := bson.M{"schema_name": schemaName, "group_hash": groupHash}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: 1}})

	cursor, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer cursor.Close(ctx)

	records := make([]*entities.Record, 0)

	for cursor.Next(ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, fmt.Errorf("failed decoding record: %w", err)
		}

		records = append(records, &lr)
	}

	return records, nil
}

func (s *RecordStorage) GetAllRecordsCount(ctx context.Context) (int, error) {
	totalCount, err := s.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("counting all records: %w", err)
	}

	return int(totalCount), nil
}

func (s *RecordStorage) FindStaleIDs(ctx context.Context, schemaID int, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	filter := bson.M{
		"schema_id":  schemaID,
		"created_at": bson.M{"$lte": threshold},
	}

	rowsCount, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("counting stale records: %w", err)
	}

	ids := make([]int, 0, rowsCount)

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("querying stale record's ids: %w", err)
	}

	for cursor.Next(ctx) {
		var lr entities.Record
		if err := cursor.Decode(&lr); err != nil {
			return nil, fmt.Errorf("failed decoding record's id: %w", err)
		}

		ids = append(ids, lr.ID)
	}

	return ids, nil
}

func (s *RecordStorage) DeleteByIDs(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	filter := bson.M{"id": bson.M{"$in": ids}}
	_, err := s.collection.DeleteMany(ctx, filter)

	return err
}
