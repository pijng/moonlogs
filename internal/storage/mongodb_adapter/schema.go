package mongodb_adapter

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchemaStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewSchemaStorage(db *mongo.Database) *SchemaStorage {
	return &SchemaStorage{
		db:         db,
		collection: db.Collection("schemas"),
	}
}

func (s *SchemaStorage) CreateSchema(ctx context.Context, schema entities.Schema) (*entities.Schema, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "schemas")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	schema.ID = nextValue

	result, err := s.collection.InsertOne(ctx, schema)
	if err != nil {
		return nil, fmt.Errorf("failed inserting schema: %w", err)
	}

	var sm entities.Schema
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&sm)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying inserted schema: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) UpdateSchemaByID(ctx context.Context, id int, schema entities.Schema) (*entities.Schema, error) {
	update := bson.M{
		"description": schema.Description, "title": schema.Title, "fields": schema.Fields,
		"kinds": schema.Kinds, "retention_days": schema.RetentionDays,
	}

	if schema.TagID != 0 {
		update["tag_id"] = schema.TagID
	}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": update})
	if err != nil {
		return nil, fmt.Errorf("failed updating schema: %w", err)
	}

	return s.GetById(ctx, id)
}

func (s *SchemaStorage) GetById(ctx context.Context, id int) (*entities.Schema, error) {
	var sm entities.Schema

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&sm)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed querying schema by id: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetByTagID(ctx context.Context, tagID int) ([]*entities.Schema, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	filter := bson.M{"tag_id": tagID}
	cursor, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer cursor.Close(ctx)

	schemas := make([]*entities.Schema, 0)

	for cursor.Next(ctx) {
		var sm entities.Schema
		if err := cursor.Decode(&sm); err != nil {
			return nil, fmt.Errorf("failed decoding schema: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) GetByName(ctx context.Context, name string) (*entities.Schema, error) {
	var sm entities.Schema

	err := s.collection.FindOne(ctx, bson.M{"name": name}).Decode(&sm)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying schema by name: %w", err)
	}

	return &sm, nil
}

func (s *SchemaStorage) GetSchemasByTitleOrDescription(ctx context.Context, title string, description string) ([]*entities.Schema, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	filter := bson.M{
		"title":       bson.M{"$regex": primitive.Regex{Pattern: title, Options: "i"}},
		"description": bson.M{"$regex": primitive.Regex{Pattern: description, Options: "i"}},
	}

	cursor, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer cursor.Close(ctx)

	schemas := make([]*entities.Schema, 0)

	for cursor.Next(ctx) {
		var sm entities.Schema
		if err := cursor.Decode(&sm); err != nil {
			return nil, fmt.Errorf("failed decoding schema: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) GetAllSchemas(ctx context.Context) ([]*entities.Schema, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying schemas: %w", err)
	}
	defer cursor.Close(ctx)

	schemas := make([]*entities.Schema, 0)

	for cursor.Next(ctx) {
		var sm entities.Schema
		if err := cursor.Decode(&sm); err != nil {
			return nil, fmt.Errorf("failed decoding api schema: %w", err)
		}

		schemas = append(schemas, &sm)
	}

	return schemas, nil
}

func (s *SchemaStorage) DeleteSchemaByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting schema: %w", err)
	}

	return err
}
