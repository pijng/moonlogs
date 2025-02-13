package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionIncident = "incidents"

type IncidentStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewIncidentStorage(db *mongo.Database) *IncidentStorage {
	return &IncidentStorage{
		db:         db,
		collection: db.Collection(collectionIncident),
	}
}

func (s *IncidentStorage) CreateIncident(ctx context.Context, incident entities.Incident) (*entities.Incident, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, collectionIncident)
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	incident.ID = nextValue

	result, err := s.collection.InsertOne(ctx, incident)
	if err != nil {
		return nil, fmt.Errorf("failed inserting incident: %w", err)
	}

	var inc entities.Incident
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&inc)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying inserted incident: %w", err)
	}

	return &inc, nil
}

func (s *IncidentStorage) GetAllIncidents(ctx context.Context) ([]*entities.Incident, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "ttl", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying incidents: %w", err)
	}
	defer cursor.Close(ctx)

	incidents := make([]*entities.Incident, 0)

	for cursor.Next(ctx) {
		var inc entities.Incident
		if err := cursor.Decode(&inc); err != nil {
			return nil, fmt.Errorf("failed decoding incident: %w", err)
		}

		incidents = append(incidents, &inc)
	}

	return incidents, nil
}

func (s *IncidentStorage) GetIncidentsByKeys(ctx context.Context, keys entities.JSONMap[any]) ([]*entities.Incident, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "ttl", Value: -1}})

	filter := bson.M{}

	if len(keys) != 0 {
		filter = qrx.KeysObject(filter, keys)
	}

	cursor, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying incidents: %w", err)
	}
	defer cursor.Close(ctx)

	incidents := make([]*entities.Incident, 0)

	for cursor.Next(ctx) {
		var inc entities.Incident
		if err := cursor.Decode(&inc); err != nil {
			return nil, fmt.Errorf("failed decoding incident: %w", err)
		}

		incidents = append(incidents, &inc)
	}

	return incidents, nil
}

func (s *IncidentStorage) FindStaleIDs(ctx context.Context, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	filter := bson.M{
		"ttl": bson.M{"$lte": threshold},
	}

	rowsCount, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("counting stale incidents: %w", err)
	}

	ids := make([]int, 0, rowsCount)

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("querying stale incidents's ids: %w", err)
	}

	for cursor.Next(ctx) {
		var inc entities.Incident
		if err := cursor.Decode(&inc); err != nil {
			return nil, fmt.Errorf("failed decoding incident's id: %w", err)
		}

		ids = append(ids, inc.ID)
	}

	return ids, nil
}

func (s *IncidentStorage) DeleteByIDs(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	filter := bson.M{"id": bson.M{"$in": ids}}
	_, err := s.collection.DeleteMany(ctx, filter)

	return err
}
