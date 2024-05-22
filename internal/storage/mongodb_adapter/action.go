package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/persistence"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ActionStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewActionStorage(ctx context.Context) *ActionStorage {
	return &ActionStorage{
		ctx:        ctx,
		collection: persistence.MongoDB().Database("moonlogs").Collection("actions"),
	}
}

func (s *ActionStorage) CreateAction(action entities.Action) (*entities.Action, error) {
	nextValue, err := getNextSequenceValue(s.ctx, persistence.MongoDB(), "actions")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	action.ID = nextValue

	result, err := s.collection.InsertOne(s.ctx, action)
	if err != nil {
		return nil, fmt.Errorf("failed inserting action: %w", err)
	}

	var t entities.Action
	err = s.collection.FindOne(s.ctx, bson.M{"_id": result.InsertedID}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted action: %w", err)
	}

	return &t, nil
}

func (s *ActionStorage) GetActionByID(id int) (*entities.Action, error) {
	var t entities.Action

	err := s.collection.FindOne(s.ctx, bson.M{"id": id}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying action by id: %w", err)
	}

	return &t, nil
}

func (s *ActionStorage) DeleteActionByID(id int) error {
	_, err := s.collection.DeleteOne(s.ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting action: %w", err)
	}

	return err
}

func (s *ActionStorage) UpdateActionByID(id int, action entities.Action) (*entities.Action, error) {
	update := bson.M{"$set": bson.M{
		"name":       action.Name,
		"pattern":    action.Pattern,
		"method":     action.Method,
		"conditions": action.Conditions,
		"disabled":   action.Disabled,
	}}

	_, err := s.collection.UpdateOne(s.ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating action: %w", err)
	}

	return s.GetActionByID(id)
}

func (s *ActionStorage) GetAllActions() ([]*entities.Action, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(s.ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying actions: %w", err)
	}
	defer cursor.Close(s.ctx)

	actions := make([]*entities.Action, 0)

	for cursor.Next(s.ctx) {
		var t entities.Action
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding action: %w", err)
		}

		actions = append(actions, &t)
	}

	return actions, nil
}
