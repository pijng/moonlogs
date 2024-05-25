package persistence

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB(dataSourceName string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dataSourceName))
	if err != nil {
		return nil, fmt.Errorf("connecting mongo db: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("pinging mongo db: %w", err)
	}

	err = createIndexes(client)
	if err != nil {
		return nil, fmt.Errorf("creating indexes: %w", err)
	}

	return client, err
}

func createIndexes(client *mongo.Client) error {
	collection := client.Database("moonlogs").Collection("records")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "schema_name", Value: 1}},
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return fmt.Errorf("index `schema_name` on `records` collection: %w", err)
	}

	return nil
}
