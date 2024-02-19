package sqlite_adapter

import (
	"context"
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
}

func NewRecordStorage(ctx context.Context) *RecordStorage {
	return &RecordStorage{
		ctx:     ctx,
		records: qrx.Scan(entities.Record{}).With(persistence.DB()).From("records"),
	}
}

func (s *RecordStorage) CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error) {
	lr, err := s.records.Create(s.ctx, map[string]interface{}{
		"text":        record.Text,
		"schema_name": record.SchemaName,
		"schema_id":   schemaID,
		"query":       record.Query,
		"kind":        record.Kind,
		"group_hash":  groupHash,
		"level":       record.Level,
		"created_at":  entities.RecordTime{Time: time.Now()},
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordByID(id int) (*entities.Record, error) {
	lr, err := s.records.Where("id = ?", id).First(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordsByQuery(record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, error) {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("(schema_id = ? OR schema_name = ?)")
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

	queryBuilder.WriteString(" ORDER BY id DESC LIMIT ? OFFSET ?")
	queryParams = append(queryParams, limit, offset)

	lr, err := s.records.Where(queryBuilder.String(), queryParams...).All(s.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	lr, err := s.records.All(s.ctx, "LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	lr, err := s.records.Where("schema_name = ? AND group_hash = ? ORDER BY id ASC", schemaName, groupHash).All(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (s *RecordStorage) GetRecordsCountByQuery(record entities.Record, from *time.Time, to *time.Time) (int, error) {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("(schema_id = ? OR schema_name = ?)")
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

	count, err := s.records.CountWhere(s.ctx, queryBuilder.String(), queryParams...)
	if err != nil {
		return 0, fmt.Errorf("failed querying record: %w", err)
	}

	return count, nil
}

func (s *RecordStorage) GetAllRecordsCount() (int, error) {
	count, err := s.records.Count(s.ctx)
	if err != nil {
		return 0, fmt.Errorf("failed querying record: %w", err)
	}

	return count, nil
}

func (s *RecordStorage) FindStale(schemaID int, threshold int64) ([]*entities.Record, error) {
	return s.records.Where("schema_id = ? AND created_at <= ?", schemaID, threshold).All(s.ctx)
}

func (s *RecordStorage) DeleteByIDs(ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders, args := qrx.In(ids)

	_, err := s.records.DeleteAll(s.ctx, fmt.Sprintf("id IN (%s)", strings.Join(placeholders, ",")), args...)

	return err
}
