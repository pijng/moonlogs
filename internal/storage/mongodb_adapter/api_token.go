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

type ApiTokenStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewApiTokenStorage(db *mongo.Database) *ApiTokenStorage {
	return &ApiTokenStorage{
		db:         db,
		collection: db.Collection("api_tokens"),
	}
}

func (s *ApiTokenStorage) CreateApiToken(ctx context.Context, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, "api_tokens")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	apiToken.ID = nextValue

	result, err := s.collection.InsertOne(ctx, apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed inserting api_token: %w", err)
	}

	var t entities.ApiToken
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("failed querying inserted api token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) UpdateApiTokenByID(ctx context.Context, id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	update := bson.M{"$set": bson.M{"name": apiToken.Name, "is_revoked": apiToken.IsRevoked}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	return s.GetApiTokenByID(ctx, id)
}

func (s *ApiTokenStorage) GetApiTokenByID(ctx context.Context, id int) (*entities.ApiToken, error) {
	var t entities.ApiToken

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying api token by id: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetApiTokenByDigest(ctx context.Context, digest string) (*entities.ApiToken, error) {
	var t entities.ApiToken

	err := s.collection.FindOne(ctx, bson.M{"token_digest": digest}).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying api token by id: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens(ctx context.Context) ([]*entities.ApiToken, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying api tokens: %w", err)
	}
	defer cursor.Close(ctx)

	tokens := make([]*entities.ApiToken, 0)

	for cursor.Next(ctx) {
		var t entities.ApiToken
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding api token: %w", err)
		}

		tokens = append(tokens, &t)
	}

	return tokens, nil
}

func (s *ApiTokenStorage) DeleteApiTokenByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting api token: %w", err)
	}

	return err
}
