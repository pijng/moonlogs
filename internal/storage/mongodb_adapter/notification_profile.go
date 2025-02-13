package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionNotificationProfile = "notification_profiles"

type NotificationProfileStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewNotificationProfileStorage(db *mongo.Database) *NotificationProfileStorage {
	return &NotificationProfileStorage{
		db:         db,
		collection: db.Collection(collectionNotificationProfile),
	}
}

func (s *NotificationProfileStorage) CreateNotificationProfile(ctx context.Context, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, collectionNotificationProfile)
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	profile.ID = nextValue

	result, err := s.collection.InsertOne(ctx, profile)
	if err != nil {
		return nil, fmt.Errorf("failed inserting notitication profile: %w", err)
	}

	var np entities.NotificationProfile
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&np)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying inserted notitication profile: %w", err)
	}

	return &np, nil
}

func (s *NotificationProfileStorage) GetNotificationProfileByID(ctx context.Context, id int) (*entities.NotificationProfile, error) {
	var t entities.NotificationProfile

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying notitication profile by id: %w", err)
	}

	return &t, nil
}

func (s *NotificationProfileStorage) DeleteNotificationProfileByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting notitication profile: %w", err)
	}

	return err
}

func (s *NotificationProfileStorage) UpdateNotificationProfileByID(ctx context.Context, id int, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	update := bson.M{"$set": bson.M{
		"name":        profile.Name,
		"description": profile.Description,
		"rule_ids":    profile.RuleIDs,
		"enabled":     profile.Enabled,
		"url":         profile.URL,
		"method":      profile.Method,
		"headers":     profile.Headers,
		"payload":     profile.Payload,
	}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating notitication profile: %w", err)
	}

	return s.GetNotificationProfileByID(ctx, id)
}

func (s *NotificationProfileStorage) GetAllNotificationProfiles(ctx context.Context) ([]*entities.NotificationProfile, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying notitication profiles: %w", err)
	}
	defer cursor.Close(ctx)

	profiles := make([]*entities.NotificationProfile, 0)

	for cursor.Next(ctx) {
		var np entities.NotificationProfile
		if err := cursor.Decode(&np); err != nil {
			return nil, fmt.Errorf("failed decoding notitication profile: %w", err)
		}

		profiles = append(profiles, &np)
	}

	return profiles, nil
}
