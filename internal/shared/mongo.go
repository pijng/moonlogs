package shared

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(collection *mongo.Collection, indexNames [][]string) error {
	for _, names := range indexNames {
		keys := bson.D{}
		for _, key := range names {
			keys = append(keys, bson.E{Key: key, Value: 1})
		}

		indexModel := mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(false),
		}

		_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
		if err != nil {
			return fmt.Errorf("creating index `%v` on `records` collection: %w", keys, err)
		}
	}

	return nil
}
