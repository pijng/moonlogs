package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
	"strings"
)

type IncidentStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewIncidentStorage(readDB *sql.DB, writeDB *sql.DB) *IncidentStorage {
	return &IncidentStorage{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (s *IncidentStorage) CreateIncident(ctx context.Context, incident entities.Incident) (*entities.Incident, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO incidents (rule_name, rule_id, keys, count, ttl) VALUES (?,?,?,?,?) RETURNING id, rule_name, rule_id, keys, count, ttl;"

	row := tx.QueryRowContext(ctx, query, incident.RuleName, incident.RuleID, incident.Keys, incident.Count, incident.TTL)

	var inc entities.Incident
	err = row.Scan(&inc.ID, &inc.RuleName, &inc.RuleID, &inc.Keys, &inc.Count, &inc.TTL)
	if err != nil {
		return nil, fmt.Errorf("failed scanning incident: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &inc, nil
}

func (s *IncidentStorage) GetAllIncidents(ctx context.Context) ([]*entities.Incident, error) {
	query := "SELECT id, rule_name, rule_id, keys, count, ttl FROM incidents ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying incidents: %w", err)
	}
	defer rows.Close()

	incidents := make([]*entities.Incident, 0)

	for rows.Next() {
		var dest entities.Incident

		err := rows.Scan(&dest.ID, &dest.RuleName, &dest.RuleID, &dest.Keys, &dest.Count, &dest.TTL)
		if err != nil {
			return nil, fmt.Errorf("failed scanning incident: %w", err)
		}

		incidents = append(incidents, &dest)
	}

	return incidents, nil
}

func (s *IncidentStorage) GetIncidentsByKeys(ctx context.Context, keys entities.JSONMap[any]) ([]*entities.Incident, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, rule_name, rule_id, keys, count, ttl FROM incidents ")

	if len(keys) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(qrx.KeysMap(keys))
	}

	queryBuilder.WriteString(" ORDER BY id DESC;")

	stmt, err := s.readDB.PrepareContext(ctx, queryBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying incidents: %w", err)
	}
	defer rows.Close()

	incidents := make([]*entities.Incident, 0)

	for rows.Next() {
		var dest entities.Incident

		err := rows.Scan(&dest.ID, &dest.RuleName, &dest.RuleID, &dest.Keys, &dest.Count, &dest.TTL)
		if err != nil {
			return nil, fmt.Errorf("failed scanning incident: %w", err)
		}

		incidents = append(incidents, &dest)
	}

	return incidents, nil
}

func (s *IncidentStorage) FindStaleIDs(ctx context.Context, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	var rowsCount int
	countQuery := "SELECT COUNT(*) FROM incidents WHERE ttl <= ?"
	err := s.readDB.QueryRowContext(ctx, countQuery, threshold).Scan(&rowsCount)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying count of stale incidents IDs: %w", err)
	}

	rows, err := s.readDB.QueryContext(ctx, "SELECT id FROM incidents WHERE ttl <= ?", threshold)
	if err != nil {
		return nil, fmt.Errorf("failed querying stale incidents IDs: %w", err)
	}
	defer rows.Close()

	ids := make([]int, 0, rowsCount)

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed scanning stale incidents ID: %w", err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *IncidentStorage) DeleteByIDs(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders, args := qrx.In(ids)

	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM incidents WHERE id IN (%s);"
	query = fmt.Sprintf(query, placeholders)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed deleting incident: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
