package usecases

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"slices"
	"strings"
)

type AlertingRuleUseCase struct {
	alertingRuleStorage storage.AlertingRuleStorage
}

func NewAlertingRuleUseCase(alertingRuleStorage storage.AlertingRuleStorage) *AlertingRuleUseCase {
	return &AlertingRuleUseCase{alertingRuleStorage: alertingRuleStorage}
}

func (uc *AlertingRuleUseCase) CreateRule(ctx context.Context, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	if rule.Name == "" {
		return nil, fmt.Errorf("failed creating rule: `name` attribute is required")
	}

	if rule.Description == "" {
		return nil, fmt.Errorf("failed creating rule: `description` attribute is required")
	}

	if (rule.Interval == entities.Duration{}) {
		return nil, fmt.Errorf("failed creating rule: `interval` attribute is required")
	}

	if rule.Severity == "" {
		return nil, fmt.Errorf("failed creating rule: `severity` attribute is required")
	}

	if err := validateSeverity(rule.Severity); err != nil {
		return nil, err
	}

	if rule.Threshold == 0 {
		return nil, fmt.Errorf("failed creating rule: `threshold` attribute is required and must be greater than 0")
	}

	if err := validateCondition(rule.Condition); err != nil {
		return nil, err
	}

	if err := validateFilter(rule); err != nil {
		return nil, err
	}

	if err := validateAggregation(rule); err != nil {
		return nil, err
	}

	return uc.alertingRuleStorage.CreateRule(ctx, rule)
}

func (uc *AlertingRuleUseCase) GetAllRules(ctx context.Context) ([]*entities.AlertingRule, error) {
	return uc.alertingRuleStorage.GetAllRules(ctx)
}

func (uc *AlertingRuleUseCase) DeleteRuleByID(ctx context.Context, id int) error {
	return uc.alertingRuleStorage.DeleteRuleByID(ctx, id)
}

func (uc *AlertingRuleUseCase) GetRuleByID(ctx context.Context, id int) (*entities.AlertingRule, error) {
	return uc.alertingRuleStorage.GetRuleByID(ctx, id)
}

func (uc *AlertingRuleUseCase) UpdateRuleByID(ctx context.Context, id int, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	return uc.alertingRuleStorage.UpdateRuleByID(ctx, id, rule)
}

func validateSeverity(serverity entities.Level) error {
	isValidSeverity := slices.Contains(entities.AppropriateLevels, string(serverity))
	if !isValidSeverity {
		appropriateLevels := strings.Join(entities.AppropriateLevels, ", ")
		return fmt.Errorf("failed creating rule: `severity` field should be one of: %v", appropriateLevels)
	}

	return nil
}

func validateCondition(condition entities.AlertCondition) error {
	isValidCondition := slices.Contains(entities.AppropriateAlertConditions, string(condition))
	if !isValidCondition {
		appropriateConditions := strings.Join(entities.AppropriateAlertConditions, ", ")
		return fmt.Errorf("failed creating rule: `condition` field should be one of: %v", appropriateConditions)
	}

	return nil
}

func validateFilter(rule entities.AlertingRule) error {
	isValidLevel := slices.Contains(entities.AppropriateLevels, string(rule.FilterLevel))
	if !isValidLevel {
		appropriateLevels := strings.Join(entities.AppropriateLevels, ", ")
		return fmt.Errorf("failed creating rule: `filter_level` field should be one of: %v", appropriateLevels)
	}

	return nil
}

func validateAggregation(rule entities.AlertingRule) error {
	isValidType := slices.Contains(entities.AppropriateAggregationTypes, string(rule.AggregationType))
	if !isValidType {
		appropriateTypes := strings.Join(entities.AppropriateAggregationTypes, ", ")
		return fmt.Errorf("failed creating rule: `aggregation_type` field should be one of: %v", appropriateTypes)
	}

	groupBy := []string(rule.AggregationGroupBy)

	for _, f := range rule.FilterSchemaFields {
		if !slices.Contains(groupBy, f) {
			appropriateGroups := strings.Join(groupBy, ", ")
			return fmt.Errorf("failed creating rule: `aggregation_group_by` field should be one of: %v", appropriateGroups)
		}
	}

	if (rule.AggregationTimeWindow == entities.Duration{}) {
		return fmt.Errorf("failed creating rule: `aggregation_time_window` attribute is required")
	}

	return nil
}
