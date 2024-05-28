package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ActionStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewActionStorage(db *mongo.Database) *ActionStorage {
	return &ActionStorage{
		db:         db,
		collection: db.Collection("actions"),
	}
}

func (s *ActionStorage) CreateAction(ctx context.Context, action entities.Action) (*entities.Action, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "actions")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	action.ID = nextValue

	result, err := s.collection.InsertOne(ctx, action)
	if err != nil {
		return nil, fmt.Errorf("failed inserting action: %w", err)
	}

	var t entities.Action
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted action: %w", err)
	}

	return &t, nil
}

func (s *ActionStorage) GetActionByID(ctx context.Context, id int) (*entities.Action, error) {
	var t entities.Action

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&t)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed querying action by id: %w", err)
	}

	return &t, nil
}

func (s *ActionStorage) DeleteActionByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting action: %w", err)
	}

	return err
}

func (s *ActionStorage) UpdateActionByID(ctx context.Context, id int, action entities.Action) (*entities.Action, error) {
	update := bson.M{"$set": bson.M{
		"name":       action.Name,
		"pattern":    action.Pattern,
		"method":     action.Method,
		"conditions": action.Conditions,
		"disabled":   action.Disabled,
		"schema_ids": action.SchemaIDs,
	}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating action: %w", err)
	}

	return s.GetActionByID(ctx, id)
}

func (s *ActionStorage) GetAllActions(ctx context.Context) ([]*entities.Action, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying actions: %w", err)
	}
	defer cursor.Close(ctx)

	actions := make([]*entities.Action, 0)

	for cursor.Next(ctx) {
		var t entities.Action
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding action: %w", err)
		}

		actions = append(actions, &t)
	}

	return actions, nil
}
