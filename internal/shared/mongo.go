package shared

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(collection *mongo.Collection, indexNames []string) error {
	for _, name := range indexNames {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: name, Value: 1}},
			Options: options.Index().SetUnique(false),
		}

		_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
		if err != nil {
			return fmt.Errorf("creating index `%s` on `records` collection: %w", name, err)
		}
	}

	return nil
}
