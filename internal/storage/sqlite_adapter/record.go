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
	records *qrx.TableQuerier[entities.Record]
	db      *sql.DB
}

func NewRecordStorage(ctx context.Context) *RecordStorage {
	return &RecordStorage{
		ctx:     ctx,
		records: qrx.Scan(entities.Record{}).With(persistence.DB()).From("records"),
		db:      persistence.DB(),
	}
}

func (s *RecordStorage) CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error) {
	query := "INSERT INTO records (text, schema_name, schema_id, query, request, response, kind, group_hash, level, created_at) VALUES (?,?,?,?,?,?,?,?,?,?);"
	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	result, err := stmt.ExecContext(s.ctx, record.Text, record.SchemaName, schemaID, record.Query,
		record.Request, record.Response, record.Kind, groupHash, record.Level, entities.RecordTime{Time: time.Now()})

	if err != nil {
		return nil, fmt.Errorf("failed inserting record: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving record last insert id: %w", err)
	}

	lr, err := s.GetRecordByID(int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordByID(id int) (*entities.Record, error) {
	query := "SELECT * FROM records WHERE id = ? LIMIT 1;"
	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	row := stmt.QueryRowContext(s.ctx, id)

	var dest entities.Record
	err = row.Scan(&dest.ID, &dest.Text, &dest.CreatedAt, &dest.SchemaName, &dest.SchemaID, &dest.Query, &dest.Kind, &dest.GroupHash, &dest.Level, &dest.Request, &dest.Response)
	if err != nil {
		return nil, fmt.Errorf("failed scanning record: %w", err)
	}

	return &dest, nil
}

func (s *RecordStorage) GetRecordsByQuery(record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, error) {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("SELECT * FROM records WHERE (schema_id = ? OR schema_name = ?)")
	queryParams := []interface{}{record.SchemaID, record.SchemaName}

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

	queryBuilder.WriteString(fmt.Sprintf(" AND %s", qrx.MapLike(record.Query)))
	queryBuilder.WriteString(fmt.Sprintf(" AND created_at BETWEEN %s", qrx.Between(from, to)))

	queryBuilder.WriteString(" ORDER BY id DESC LIMIT ? OFFSET ?;")
	queryParams = append(queryParams, limit, offset)

	stmt, err := qrx.CachedStmt(s.ctx, s.db, queryBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	rows, err := stmt.QueryContext(s.ctx, queryParams...)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

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

func (s *RecordStorage) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	query := "SELECT * FROM records LIMIT ? OFFSET ?;"
	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	rows, err := stmt.QueryContext(s.ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}

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
	query := "SELECT * FROM records WHERE schema_name = ? AND group_hash = ? ORDER BY id ASC;"

	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	rows, err := stmt.QueryContext(s.ctx, schemaName, groupHash)
	if err != nil {
		return nil, fmt.Errorf("failed querying records: %w", err)
	}

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

func (s *RecordStorage) GetRecordsCountByQuery(record entities.Record, from *time.Time, to *time.Time) (int, error) {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("SELECT COUNT(*) FROM records WHERE (schema_id = ? OR schema_name = ?)")
	queryParams := []interface{}{record.SchemaID, record.SchemaName}

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

	queryBuilder.WriteString(fmt.Sprintf(" AND %s", qrx.MapLike(record.Query)))
	queryBuilder.WriteString(fmt.Sprintf(" AND created_at BETWEEN %s;", qrx.Between(from, to)))

	stmt, err := qrx.CachedStmt(s.ctx, s.db, queryBuilder.String())
	if err != nil {
		return 0, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

	row := stmt.QueryRowContext(s.ctx, queryParams...)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *RecordStorage) GetAllRecordsCount() (int, error) {
	query := "SELECT COUNT(*) FROM records;"
	stmt, err := qrx.CachedStmt(s.ctx, s.db, query)
	if err != nil {
		return 0, fmt.Errorf("failed retrieving cached statement: %w", err)
	}

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
	err := s.db.QueryRowContext(s.ctx, countQuery, schemaID, threshold).Scan(&rowsCount)
	if err != nil {
		return nil, fmt.Errorf("failed querying count of stale records IDs: %w", err)
	}

	rows, err := s.db.QueryContext(s.ctx, "SELECT id FROM records WHERE schema_id = ? AND created_at <= ?", schemaID, threshold)
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

	_, err := s.records.DeleteAll(s.ctx, fmt.Sprintf("id IN (%s)", strings.Join(placeholders, ",")), args...)

	return err
}
