package usecases

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"slices"
	"strings"
	"time"
)

type RecordUseCase struct {
	recordStorage storage.RecordStorage
}

func NewRecordUseCase(recordStorage storage.RecordStorage) *RecordUseCase {
	return &RecordUseCase{recordStorage: recordStorage}
}

func (uc *RecordUseCase) CreateRecord(ctx context.Context, record entities.Record, schemaID int) (*entities.Record, error) {

	if len(record.Query) == 0 {
		return nil, fmt.Errorf("failed creating record: `query` attribute is required")
	}

	if len(record.Level) > 0 {
		isValidLevel := slices.Contains(entities.AppropriateLevels, string(record.Level))
		if !isValidLevel {
			appropriateLevels := strings.Join(entities.AppropriateLevels, ", ")
			return nil, fmt.Errorf("failed creating record: `level` field should be one of: %v", appropriateLevels)
		}
	} else {
		record.Level = entities.InfoLevel
	}

	if record.CreatedAt.Equal(time.Time{}) {
		record.CreatedAt = entities.RecordTime{Time: time.Now()}
	}

	groupHash, err := shared.HashQuery(record.Query)
	if err != nil {
		return nil, fmt.Errorf("failed calculating record query hash: %w", err)
	}

	record.SchemaID = schemaID
	record.GroupHash = groupHash

	return uc.recordStorage.CreateRecord(ctx, record)
}

func (uc *RecordUseCase) DeleteStaleRecords(ctx context.Context, schema *entities.Schema) error {
	// Treat 0 retention days as infinite
	if schema.RetentionDays == 0 {
		return nil
	}

	threshold := time.Now().Add(-time.Duration(schema.RetentionDays) * 24 * time.Hour).UnixMilli()

	staleRecordIDs, err := uc.recordStorage.FindStaleIDs(ctx, schema.ID, threshold)
	if err != nil {
		return fmt.Errorf("DeleteStaleRecords: failed to query stale log records: %w", err)
	}

	staleRecordIDsBatches := shared.BatchSlice(staleRecordIDs, 950)

	for _, ids := range staleRecordIDsBatches {
		err = uc.recordStorage.DeleteByIDs(ctx, ids)

		if err != nil {
			return fmt.Errorf("DeleteStateleRecords: failed to delete stale log records: %w", err)
		}
	}

	return nil
}

func (uc *RecordUseCase) GetAllRecords(ctx context.Context, limit int, offset int) ([]*entities.Record, error) {
	return uc.recordStorage.GetAllRecords(ctx, limit, offset)
}

func (uc *RecordUseCase) GetAllRecordsCount(ctx context.Context) (int, error) {
	return uc.recordStorage.GetAllRecordsCount(ctx)
}

func (uc *RecordUseCase) GetRecordByID(ctx context.Context, id int) (*entities.Record, error) {
	return uc.recordStorage.GetRecordByID(ctx, id)
}

func (uc *RecordUseCase) GetRecordsByQuery(ctx context.Context, record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error) {
	return uc.recordStorage.GetRecordsByQuery(ctx, record, from, to, limit, offset)
}

func (uc *RecordUseCase) GetRecordsByGroupHash(ctx context.Context, schemaName string, groupHash string) ([]*entities.Record, error) {
	return uc.recordStorage.GetRecordsByGroupHash(ctx, schemaName, groupHash)
}

func (uc *RecordUseCase) AggregateRecords(ctx context.Context, filter entities.RecordFilter, aggregation entities.RecordAggregation) ([]*entities.AggregationGroup, error) {
	return uc.recordStorage.AggregateRecords(ctx, filter, aggregation)
}
