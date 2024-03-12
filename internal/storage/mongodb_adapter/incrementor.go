package mongodb_adapter

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sequence struct {
	Name  string `bson:"_id"`
	Value int    `bson:"value"`
}

func getNextSequenceValue(ctx context.Context, client *mongo.Client, sequenceName string) (int, error) {
	sequences := client.Database("moonlogs").Collection("sequences")

	filter := map[string]interface{}{"_id": sequenceName}
	update := map[string]interface{}{"$inc": map[string]interface{}{"value": 1}}

	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var sequence Sequence
	err := sequences.FindOneAndUpdate(ctx, filter, update, options).Decode(&sequence)
	if err != nil {
		return 0, err
	}

	return sequence.Value, nil
}
