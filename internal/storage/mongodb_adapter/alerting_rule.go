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

const collectionAlertingRule = "alerting_rules"

type AlertingRuleStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewAlertingRuleStorage(db *mongo.Database) *AlertingRuleStorage {
	return &AlertingRuleStorage{
		db:         db,
		collection: db.Collection(collectionAlertingRule),
	}
}

func (s *AlertingRuleStorage) CreateRule(ctx context.Context, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	nextValue, err := getNextSequenceValue(ctx, s.db, collectionAlertingRule)
	if err != nil {
		return nil, fmt.Errorf("getting next sequence value: %w", err)
	}

	rule.ID = nextValue

	result, err := s.collection.InsertOne(ctx, rule)
	if err != nil {
		return nil, fmt.Errorf("failed inserting alerting rule: %w", err)
	}

	var ar entities.AlertingRule
	err = s.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&ar)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying inserted alerting rule: %w", err)
	}

	return &ar, nil
}

func (s *AlertingRuleStorage) GetRuleByID(ctx context.Context, id int) (*entities.AlertingRule, error) {
	var ar entities.AlertingRule

	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&ar)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying alerting rule by id: %w", err)
	}

	return &ar, nil
}

func (s *AlertingRuleStorage) DeleteRuleByID(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed deleting alerting rule: %w", err)
	}

	return err
}

func (s *AlertingRuleStorage) UpdateRuleByID(ctx context.Context, id int, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	update := bson.M{"$set": bson.M{
		"name":                    rule.Name,
		"description":             rule.Description,
		"enabled":                 rule.Enabled,
		"severity":                rule.Severity,
		"interval":                rule.Interval,
		"threshold":               rule.Threshold,
		"condition":               rule.Condition,
		"filter_level":            rule.FilterLevel,
		"filter_schema_ids":       rule.FilterSchemaIDs,
		"filter_schema_fields":    rule.FilterSchemaFields,
		"filter_schema_kinds":     rule.FilterSchemaKinds,
		"aggregation_type":        rule.AggregationType,
		"aggregation_group_by":    rule.AggregationGroupBy,
		"aggregation_time_window": rule.AggregationTimeWindow,
	}}

	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed updating alerting rule: %w", err)
	}

	return s.GetRuleByID(ctx, id)
}

func (s *AlertingRuleStorage) GetAllRules(ctx context.Context) ([]*entities.AlertingRule, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})

	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed querying alerting rule: %w", err)
	}
	defer cursor.Close(ctx)

	alertingRules := make([]*entities.AlertingRule, 0)

	for cursor.Next(ctx) {
		var ar entities.AlertingRule
		if err := cursor.Decode(&ar); err != nil {
			return nil, fmt.Errorf("failed decoding alerting rule: %w", err)
		}

		alertingRules = append(alertingRules, &ar)
	}

	return alertingRules, nil
}
