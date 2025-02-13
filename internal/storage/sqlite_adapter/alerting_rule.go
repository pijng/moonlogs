package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
)

type AlertingRuleStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewAlertingRuleStorage(readDB *sql.DB, writeDB *sql.DB) *AlertingRuleStorage {
	return &AlertingRuleStorage{
		writeDB: readDB,
		readDB:  writeDB,
	}
}

func (s *AlertingRuleStorage) CreateRule(ctx context.Context, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := `INSERT INTO alerting_rules
		(name, description, enabled,
		severity, interval, threshold,
		condition, filter_level, filter_schema_ids,
		filter_schema_fields, filter_schema_kinds, aggregation_type,
		aggregation_group_by, aggregation_time_window)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := tx.ExecContext(ctx, query,
		rule.Name,
		rule.Description,
		rule.Enabled,
		rule.Severity,
		rule.Interval,
		rule.Threshold,
		rule.Condition,
		rule.FilterLevel,
		rule.FilterSchemaIDs,
		rule.FilterSchemaFields,
		rule.FilterSchemaKinds,
		rule.AggregationType,
		rule.AggregationGroupBy,
		rule.AggregationTimeWindow)
	if err != nil {
		return nil, fmt.Errorf("failed inserting alerting rule: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving alerting rule last insert id: %w", err)
	}

	t, err := s.GetRuleByID(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying alerting rule: %w", err)
	}

	return t, nil
}

func (s *AlertingRuleStorage) GetRuleByID(ctx context.Context, id int) (*entities.AlertingRule, error) {
	query := `SELECT id, name, description, enabled, severity,
		interval, threshold, condition, filter_level, filter_schema_ids,
		filter_schema_fields, filter_schema_kinds, aggregation_type,
		aggregation_group_by, aggregation_time_window FROM alerting_rules WHERE id=? LIMIT 1;`
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var ar entities.AlertingRule
	err = row.Scan(
		&ar.ID,
		&ar.Name,
		&ar.Description,
		&ar.Enabled,
		&ar.Severity,
		&ar.Interval,
		&ar.Threshold,
		&ar.Condition,
		&ar.FilterLevel,
		&ar.FilterSchemaIDs,
		&ar.FilterSchemaFields,
		&ar.FilterSchemaKinds,
		&ar.AggregationType,
		&ar.AggregationGroupBy,
		&ar.AggregationTimeWindow)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning alerting rule: %w", err)
	}

	return &ar, nil
}

func (s *AlertingRuleStorage) DeleteRuleByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM alerting_rules WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting alerting rule: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *AlertingRuleStorage) UpdateRuleByID(ctx context.Context, id int, rule entities.AlertingRule) (*entities.AlertingRule, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE alerting_rules SET name=?, description=?, enabled=?, severity=?, interval=?, threshold=?, condition=?, filter_level=?, filter_schema_ids=?, filter_schema_fields=?, filter_schema_kinds=?, aggregation_type=?, aggregation_group_by=?, aggregation_time_window=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query,
		rule.Name,
		rule.Description,
		rule.Enabled,
		rule.Severity,
		rule.Interval,
		rule.Threshold,
		rule.Condition,
		rule.FilterLevel,
		rule.FilterSchemaIDs,
		rule.FilterSchemaFields,
		rule.FilterSchemaKinds,
		rule.AggregationType,
		rule.AggregationGroupBy,
		rule.AggregationTimeWindow,
		id)
	if err != nil {
		return nil, fmt.Errorf("failed updating alerting rule: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetRuleByID(ctx, id)
}

func (s *AlertingRuleStorage) GetAllRules(ctx context.Context) ([]*entities.AlertingRule, error) {
	query := `SELECT id, name, description, enabled, severity,
		interval, threshold, condition, filter_level, filter_schema_ids,
		filter_schema_fields, filter_schema_kinds, aggregation_type,
		aggregation_group_by, aggregation_time_window FROM alerting_rules ORDER BY id DESC;`
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying alerting rules: %w", err)
	}
	defer rows.Close()

	alertingRules := make([]*entities.AlertingRule, 0)

	for rows.Next() {
		var ar entities.AlertingRule

		err = rows.Scan(
			&ar.ID,
			&ar.Name,
			&ar.Description,
			&ar.Enabled,
			&ar.Severity,
			&ar.Interval,
			&ar.Threshold,
			&ar.Condition,
			&ar.FilterLevel,
			&ar.FilterSchemaIDs,
			&ar.FilterSchemaFields,
			&ar.FilterSchemaKinds,
			&ar.AggregationType,
			&ar.AggregationGroupBy,
			&ar.AggregationTimeWindow)
		if err != nil {
			return nil, fmt.Errorf("failed scanning alerting rule: %w", err)
		}

		alertingRules = append(alertingRules, &ar)
	}

	return alertingRules, nil
}
