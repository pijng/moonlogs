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

type TagStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewTagStorage(ctx context.Context) *TagStorage {
	return &TagStorage{
		ctx:        ctx,
		collection: persistence.MongoDB().Database("moonlogs").Collection("tags"),
	}
}

func (s *TagStorage) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	nextValue, err := getNextSequenceValue(s.ctx, persistence.MongoDB(), "tags")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	tag.ID = nextValue

	result, err := s.collection.InsertOne(s.ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("failed inserting tag: %w", err)
	}

	var t entities.Tag
	err = s.collection.FindOne(s.ctx, bson.M{"_id": result.InsertedID}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted tag: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) GetTagByID(id int) (*entities.Tag, error) {
	var t entities.Tag

	err := s.collection.FindOne(s.ctx, bson.M{"id": id}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying tag by id: %w", err)
	}

	return &t, nil
}

func (s *TagStorage) DeleteTagByID(id int) error {
	_, err := s.collection.DeleteOne(s.ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting tag: %w", err)
	}

	return err
}

func (s *TagStorage) UpdateTagByID(id int, tag entities.Tag) (*entities.Tag, error) {
	update := bson.M{"$set": bson.M{"name": tag.Name, "view_order": tag.ViewOrder}}

	_, err := s.collection.UpdateOne(s.ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating tag: %w", err)
	}

	return s.GetTagByID(id)
}

func (s *TagStorage) GetAllTags() ([]*entities.Tag, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(s.ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying tags: %w", err)
	}
	defer cursor.Close(s.ctx)

	tags := make([]*entities.Tag, 0)

	for cursor.Next(s.ctx) {
		var t entities.Tag
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding tag: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
