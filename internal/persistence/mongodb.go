package persistence

import (
	"context"
	"fmt"
	"moonlogs/internal/shared"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB(dataSourceName string) (*mongo.Client, error) {
	connectCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(connectCtx, options.Client().ApplyURI(dataSourceName))
	if err != nil {
		return nil, fmt.Errorf("connect to MongoDB: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(pingCtx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping MongoDB: %w", err)
	}

	err = createIndexes(client)
	if err != nil {
		return nil, fmt.Errorf("create indexes: %w", err)
	}

	return client, err
}

func createIndexes(client *mongo.Client) error {
	collection := client.Database(MONGODB_DATABASE_NAME).Collection("records")
	indexNames := []string{"schema_name", "id", "group_hash"}

	return shared.CreateIndexes(collection, indexNames)
}
