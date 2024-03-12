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

type ApiTokenStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewApiTokenStorage(ctx context.Context) *ApiTokenStorage {
	return &ApiTokenStorage{
		ctx:        ctx,
		collection: persistence.MongoDB().Database("moonlogs").Collection("api_tokens"),
	}
}

func (s *ApiTokenStorage) CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error) {
	nextValue, err := getNextSequenceValue(s.ctx, persistence.MongoDB(), "api_tokens")
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	apiToken.ID = nextValue

	result, err := s.collection.InsertOne(s.ctx, apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed inserting api_token: %w", err)
	}

	var t entities.ApiToken
	err = s.collection.FindOne(s.ctx, bson.M{"_id": result.InsertedID}).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("failed querying inserted api token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	update := bson.M{"$set": bson.M{"name": apiToken.Name, "is_revoked": apiToken.IsRevoked}}

	_, err := s.collection.UpdateOne(s.ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	return s.GetApiTokenByID(id)
}

func (s *ApiTokenStorage) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	var t entities.ApiToken

	err := s.collection.FindOne(s.ctx, bson.M{"id": id}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying api token by id: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetApiTokenByDigest(digest string) (*entities.ApiToken, error) {
	var t entities.ApiToken

	err := s.collection.FindOne(s.ctx, bson.M{"token_digest": digest}).Decode(&t)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("failed querying api token by digest: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens() ([]*entities.ApiToken, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(s.ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying api tokens: %w", err)
	}
	defer cursor.Close(s.ctx)

	tokens := make([]*entities.ApiToken, 0)

	for cursor.Next(s.ctx) {
		var t entities.ApiToken
		if err := cursor.Decode(&t); err != nil {
			return nil, fmt.Errorf("failed decoding api token: %w", err)
		}

		tokens = append(tokens, &t)
	}

	return tokens, nil
}

func (s *ApiTokenStorage) DeleteApiTokenByID(id int) error {
	_, err := s.collection.DeleteOne(s.ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting api token: %w", err)
	}

	return err
}
