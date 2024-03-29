package sqlite_adapter

import (
	"context"
	"database/sql"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
	"strings"
	"time"
)

type RecordStorage struct {
	ctx     context.Context
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewRecordStorage(ctx context.Context) *RecordStorage {
	return &RecordStorage{
		ctx:     ctx,
		readDB:  persistence.SqliteReadDB(),
		writeDB: persistence.SqliteWriteDB(),
	}
}

func (s *RecordStorage) CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "INSERT INTO records (text, schema_name, schema_id, query, request, response, kind, group_hash, level, created_at) VALUES (?,?,?,?,?,?,?,?,?,?) RETURNING *;"

	row := tx.QueryRowContext(s.ctx, query, record.Text, record.SchemaName, schemaID, record.Query,
		record.Request, record.Response, record.Kind, groupHash, record.Level, record.CreatedAt)

	var lr entities.Record
	err = row.Scan(&lr.ID, &lr.Text, &lr.CreatedAt, &lr.SchemaName, &lr.SchemaID, &lr.Query, &lr.Kind, &lr.GroupHash, &lr.Level, &lr.Request, &lr.Response)
	if err != nil {
		return nil, fmt.Errorf("failed scanning record: %w", err)
	}

	return &lr, nil
}

func (s *RecordStorage) GetRecordByID(id int) (*entities.Record, error) {
	query := "SELECT * FROM records WHERE id = ? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var dest entities.Record
	err = row.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response)
	if err != nil {
		return nil, fmt.Errorf("failed scanning record: %w", err)
	}

	return &dest, nil
}

func (s *RecordStorage) GetRecordsByQuery(record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error) {
	var queryBuilder strings.Builder
	queryParams := make([]interface{}, 0)

	if record.ID != 0 {
		queryBuilder.WriteString("schema_id = ?")
		queryParams = append(queryParams, record.SchemaID)
	} else {
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
		queryBuilder.WriteString(fmt.Sprintf(" AND %s", qrx.MapLike(record.Query)))
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
			records.*
		FROM
			records
		WHERE %s`, countBuilder.String(), queryBuilder.String(),
	)

	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return make([]*entities.Record, 0), 0, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	totalParams := append(countParams, queryParams...)

	rows, err := stmt.QueryContext(s.ctx, totalParams...)
	if err != nil {
		return make([]*entities.Record, 0), 0, fmt.Errorf("failed querying record: %w", err)
	}
	defer rows.Close()

	var totalCount int
	lr := make([]*entities.Record, 0, limit)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&totalCount, &dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response)
		if err != nil {
			return make([]*entities.Record, 0), 0, fmt.Errorf("failed querying record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, totalCount, nil
}

func (s *RecordStorage) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	query := "SELECT * FROM records LIMIT ? OFFSET ?;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer rows.Close()

	lr := make([]*entities.Record, 0)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response)
		if err != nil {
			return nil, fmt.Errorf("failed querying record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	query := "SELECT * FROM records WHERE schema_name = ? AND group_hash = ? ORDER BY created_at ASC;"

	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx, schemaName, groupHash)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}
	defer rows.Close()

	lr := make([]*entities.Record, 0)

	for rows.Next() {
		var dest entities.Record

		err := rows.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response)
		if err != nil {
			return nil, fmt.Errorf("failed querying record: %w", err)
		}

		lr = append(lr, &dest)
	}

	return lr, nil
}

func (s *RecordStorage) GetAllRecordsCount() (int, error) {
	query := "SELECT COUNT(*) FROM records;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *RecordStorage) FindStaleIDs(schemaID int, threshold int64) ([]int, error) {
	// Count the number of rows before fetching the IDs to efficiently
	// pre-allocate array of ids for resulting query
	var rowsCount int
	countQuery := "SELECT COUNT(*) FROM records WHERE schema_id = ? AND created_at <= ?"
	err := s.readDB.QueryRowContext(s.ctx, countQuery, schemaID, threshold).Scan(&rowsCount)
	if err != nil {
		return nil, fmt.Errorf("failed querying count of stale records IDs: %w", err)
	}

	rows, err := s.readDB.QueryContext(s.ctx, "SELECT id FROM records WHERE schema_id = ? AND created_at <= ?", schemaID, threshold)
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

func (s *RecordStorage) DeleteByIDs(ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders, args := qrx.In(ids)

	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "DELETE FROM records WHERE id IN (%s);"
	query = fmt.Sprintf(query, placeholders)

	_, err = tx.ExecContext(s.ctx, query, args...)

	return err
}
