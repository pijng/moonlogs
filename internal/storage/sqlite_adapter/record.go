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
	"time"
)

const defaultSchemaGroupField = "schema_name"

type RecordStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewRecordStorage(readDB *sql.DB, writeDB *sql.DB) *RecordStorage {
	return &RecordStorage{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (s *RecordStorage) CreateRecord(ctx context.Context, record entities.Record) (*entities.Record, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO records (text, schema_name, schema_id, query, request, response, changes, kind, group_hash, level, created_at, old_value, new_value) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) RETURNING id, text, created_at, schema_name, schema_id, query, kind, group_hash, level, request, response, changes, old_value, new_value;"

	row := tx.QueryRowContext(ctx, query, record.Text, record.SchemaName, record.SchemaID, record.Query,
		record.Request, record.Response, record.Changes, record.Kind, record.GroupHash, record.Level, record.CreatedAt,
		record.OldValue, record.NewValue)

	var lr entities.Record
	err = row.Scan(&lr.ID, &lr.Text, &lr.CreatedAt, &lr.SchemaName, &lr.SchemaID, &lr.Query, &lr.Kind, &lr.GroupHash, &lr.Level, &lr.Request, &lr.Response, &lr.Changes, &lr.OldValue, &lr.NewValue)
	if err != nil {
		return nil, fmt.Errorf("failed scanning record: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordByID(ctx context.Context, id int) (*entities.Record, error) {
	query := "SELECT id, text, created_at, schema_name, schema_id, query, kind, group_hash, level, request, response, changes, old_value, new_value FROM records WHERE id = ? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var dest entities.Record
	err = row.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID,
		&dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response,
		&dest.Changes, &dest.OldValue, &dest.NewValue)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning record: %w", err)
	}

	return &dest, nil
}

func (s *RecordStorage) GetRecordsByQuery(ctx context.Context, record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error) {
	var queryBuilder strings.Builder
	queryParams := make([]interface{}, 0)

	if record.ID != 0 {
		queryBuilder.WriteString("schema_id = ?")
		queryParams = append(queryParams, record.SchemaID)
	}

	if record.SchemaName != "" {
		queryBuilder.WriteString("schema_name = ?")
		queryParams = append(queryParams, record.SchemaName)
	}

	if record.Text != "" {
		queryBuilder.WriteString(" AND text LIKE ?")
		queryParams = append(queryParams, qrx.Contains(record.Text))
	}
	if record.Kind != "" {
		queryBuilder.WriteString(" AND kind LIKE ?")
		queryParams = append(queryParams, qrx.Contains(record.Kind))
	}
	if record.Level != "" {
		queryBuilder.WriteString(" AND level LIKE ?")
		queryParams = append(queryParams, qrx.Contains(record.Level))
	}

	if len(record.Query) != 0 {
		queryBuilder.WriteString(fmt.Sprintf(" AND %s", qrx.QueryMap(record.Query)))
	}

	if from != nil || to != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND created_at BETWEEN %s", qrx.Between(from, to)))
	}

	countBuilder := queryBuilder
	countParams := queryParams

	queryBuilder.WriteString(" ORDER BY created_at DESC LIMIT ? OFFSET ?")
	queryParams = append(queryParams, limit, offset)

	query := fmt.Sprintf(`
		SELECT
			(SELECT COUNT(*) FROM records WHERE %s) AS total_count,
			records.id, records.text, records.created_at, records.schema_name, records.schema_id, records.query, records.kind, records.group_hash, records.level, records.request, records.response, records.changes, records.old_value, records.new_value
		FROM
			records
		WHERE %s`, countBuilder.String(), queryBuilder.String(),
	)

	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return make([]*entities.Record, 0), 0, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	totalParams := append(countParams, queryParams...)

	rows, err := stmt.QueryContext(ctx, totalParams...)
	if err != nil {
		return make([]*entities.Record, 0), 0, fmt.Errorf("failed querying record: %w", err)
	}
	defer rows.Close()

	var totalCount int
	lr := make([]*entities.Record, 0, limit)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&totalCount, &dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query,
			&dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response, &dest.Changes, &dest.OldValue, &dest.NewValue)
		if err != nil {
			return make([]*entities.Record, 0), 0, fmt.Errorf("failed scanning record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, totalCount, nil
}

func (s *RecordStorage) AggregateRecords(ctx context.Context, recordsFilter entities.RecordFilter, aggregation entities.RecordAggregation) ([]*entities.AggregationGroup, error) {
	query := `SELECT`

	groupByClauses := []string{}
	selectFields := []string{}
	for _, field := range aggregation.GroupBy {
		if field == defaultSchemaGroupField {
			selectFields = append(selectFields, fmt.Sprintf(`%s AS "%s"`, field, field))
			groupByClauses = append(groupByClauses, field)
		} else {
			selectFields = append(selectFields, fmt.Sprintf(`json_extract(query, '$.%s') AS "%s"`, field, field))
			groupByClauses = append(groupByClauses, fmt.Sprintf(`json_extract(query, '$.%s')`, field))
		}
	}

	query += " " + strings.Join(selectFields, ", ") + ", COUNT(*) AS count FROM records"

	whereClauses := []string{}
	args := []interface{}{}

	if len(recordsFilter.SchemaIDs) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("schema_id IN (%s)", qrx.Placeholders(len(recordsFilter.SchemaIDs))))
		for _, id := range recordsFilter.SchemaIDs {
			args = append(args, id)
		}
	}

	if len(recordsFilter.SchemaKinds) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("kind IN (%s)", qrx.Placeholders(len(recordsFilter.SchemaKinds))))
		for _, kind := range recordsFilter.SchemaKinds {
			args = append(args, kind)
		}
	}

	if len(recordsFilter.Level) > 0 {
		whereClauses = append(whereClauses, "level = ?")
		args = append(args, recordsFilter.Level)
	}

	whereClauses = append(whereClauses, fmt.Sprintf("created_at BETWEEN %s", qrx.Between(&recordsFilter.From, &recordsFilter.To)))
	args = append(args, recordsFilter.From, recordsFilter.To)

	if len(recordsFilter.SchemaFields) > 0 {
		for _, field := range recordsFilter.SchemaFields {
			whereClauses = append(whereClauses, fmt.Sprintf("json_extract(query, '$.%s') IS NOT NULL", field))
		}
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if len(groupByClauses) > 0 {
		query += " GROUP BY " + strings.Join(groupByClauses, ", ")
	}

	rows, err := s.readDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed aggregating records: %w", err)
	}
	defer rows.Close()

	var aggregationGroups []*entities.AggregationGroup
	for rows.Next() {
		keys := make(entities.JSONMap[any])
		values := make([]interface{}, len(aggregation.GroupBy)+1)
		valuePtrs := make([]interface{}, len(values))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed decoding aggregated records: %w", err)
		}

		for i, field := range aggregation.GroupBy {
			if strVal, ok := values[i].(string); ok {
				keys[field] = strVal
			} else {
				keys[field] = fmt.Sprintf("%v", values[i])
			}
		}

		aggregationGroups = append(aggregationGroups, &entities.AggregationGroup{
			Keys:  keys,
			Count: int32(values[len(values)-1].(int64)),
		})
	}

	return aggregationGroups, nil
}

func (s *RecordStorage) GetAllRecords(ctx context.Context, limit int, offset int) ([]*entities.Record, error) {
	query := "SELECT id, text, created_at, schema_name, schema_id, query, kind, group_hash, level, request, response, changes, old_value, new_value FROM records LIMIT ? OFFSET ?;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer rows.Close()

	lr := make([]*entities.Record, 0)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID,
			&dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response,
			&dest.Changes, &dest.OldValue, &dest.NewValue)
		if err != nil {
			return nil, fmt.Errorf("failed querying record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordsByGroupHash(ctx context.Context, schemaName string, groupHash string) ([]*entities.Record, error) {
	query := "SELECT id, text, created_at, schema_name, schema_id, query, kind, group_hash, level, request, response, changes, old_value, new_value FROM records WHERE schema_name = ? AND group_hash = ? ORDER BY created_at ASC;"

	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, schemaName, groupHash)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer rows.Close()

	lr := make([]*entities.Record, 0)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query,
			&dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response, &dest.Changes, &dest.OldValue,
			&dest.NewValue)
		if err != nil {
			return nil, fmt.Errorf("failed querying record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, nil
}

func (s *RecordStorage) GetAllRecordsCount(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM records;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *RecordStorage) FindStaleIDs(ctx context.Context, schemaID int, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	var rowsCount int
	countQuery := "SELECT COUNT(*) FROM records WHERE schema_id = ? AND created_at <= ?"
	err := s.readDB.QueryRowContext(ctx, countQuery, schemaID, threshold).Scan(&rowsCount)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed querying count of stale records IDs: %w", err)
	}

	rows, err := s.readDB.QueryContext(ctx, "SELECT id FROM records WHERE schema_id = ? AND created_at <= ?", schemaID, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed querying stale records IDs: %w", err)
	}
	defer rows.Close()

	ids := make([]int, 0, rowsCount)

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed scanning stale records ID: %w", err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *RecordStorage) DeleteByIDs(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders, args := qrx.In(ids)

	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM records WHERE id IN (%s);"
	query = fmt.Sprintf(query, placeholders)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed deleting schema: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
