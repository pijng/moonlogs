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

type TagStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTagStorage(db *mongo.Database) *TagStorage {
	return &TagStorage{
		db:         db,
		collection: db.Collection("tags"),
	}
}

func (s *TagStorage) CreateTag(ctx context.Context, tag entities.Tag) (*entities.Tag, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "tags")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	tag.ID = nextValue

	result, err := s.collection.InsertOne(ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("failed inserting tag: %w", err)
	}

	var t entities.Tag
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying inserted tag: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) GetTagByID(ctx context.Context, id int) (*entities.Tag, error) {
	var t entities.Tag

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying tag by id: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) DeleteTagByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting tag: %w", err)
	}

	return err
}

func (s *TagStorage) UpdateTagByID(ctx context.Context, id int, tag entities.Tag) (*entities.Tag, error) {
	update := bson.M{"$set": bson.M{"name": tag.Name, "view_order": tag.ViewOrder}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return s.GetTagByID(ctx, id)
}

func (s *TagStorage) GetAllTags(ctx context.Context) ([]*entities.Tag, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}
	defer cursor.Close(ctx)

	tags := make([]*entities.Tag, 0)

	for cursor.Next(ctx) {
		var t entities.Tag
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding tag: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
