package repositories

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/persistence"
	"moonlogs/lib/qrx"
	"strings"
	"time"
)

type RecordRepository struct {
	ctx     context.Context
	records *qrx.TableQuerier[entities.Record]
}

func NewRecordRepository(ctx context.Context) *RecordRepository {
	return &RecordRepository{
		ctx:     ctx,
		records: qrx.Scan(entities.Record{}).With(persistence.DB()).From("records"),
	}
}

func (r *RecordRepository) CreateRecord(record entities.Record, schemaID int, groupHash string) (*entities.Record, error) {
	lr, err := r.records.Create(r.ctx, map[string]interface{}{
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

func (r *RecordRepository) GetRecordByID(id int) (*entities.Record, error) {
	lr, err := r.records.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (r *RecordRepository) GetRecordsByQuery(record entities.Record, limit int, offset int) ([]*entities.Record, error) {
	lr, err := r.records.Where(
		fmt.Sprintf(
			`(schema_id = ? OR schema_name = ?) AND text LIKE ? AND kind LIKE ? AND %s ORDER BY created_at DESC LIMIT ? OFFSET ?`,
			qrx.MapLike(record.Query),
		),
		record.SchemaID, record.SchemaName, qrx.Contains(record.Text), qrx.Contains(record.Kind), limit, offset,
	).All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (r *RecordRepository) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	lr, err := r.records.All(r.ctx, "LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (r *RecordRepository) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	lr, err := r.records.Where("schema_name = ? AND group_hash = ? ORDER BY created_at ASC", schemaName, groupHash).All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying record: %w", err)
	}

	return lr, nil
}

func (r *RecordRepository) GetRecordsCountByQuery(record entities.Record) (int, error) {
	count, err := r.records.CountWhere(
		r.ctx,
		fmt.Sprintf("schema_name = ? AND kind LIKE ? AND text LIKE ? AND %s", qrx.MapLike(record.Query)),
		record.SchemaName, qrx.Contains(record.Kind), qrx.Contains(record.Text))
	if err != nil {
		return 0, fmt.Errorf("failed querying record: %w", err)
	}

	return count, nil
}

func (r *RecordRepository) GetAllRecordsCount() (int, error) {
	count, err := r.records.Count(r.ctx)
	if err != nil {
		return 0, fmt.Errorf("failed querying record: %w", err)
	}

	return count, nil
}

func (r *RecordRepository) FindStale(schemaID int, threshold int64) ([]*entities.Record, error) {
	return r.records.Where("schema_id = ? AND created_at <= ?", schemaID, threshold).All(r.ctx)
}

func (r *RecordRepository) DestroyByIDs(ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	_, err := r.records.DeleteAll(r.ctx, fmt.Sprintf("id IN (%s)", strings.Join(placeholders, ",")), args...)

	return err
}
