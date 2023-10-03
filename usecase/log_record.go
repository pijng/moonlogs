package usecase

import (
	"fmt"
	"moonlogs/internal/repository"
	"moonlogs/shared"
	"time"
)

type LogRecordUseCase struct {
	logRecordRepository *repository.LogRecordRepository
	logSchemaRepository *repository.LogSchemaRepository
}

func NewLogRecordsUseCase(logRecordRepository *repository.LogRecordRepository, logSchemaRepository *repository.LogSchemaRepository) *LogRecordUseCase {
	return &LogRecordUseCase{logRecordRepository: logRecordRepository, logSchemaRepository: logSchemaRepository}
}

func (uc *LogRecordUseCase) DeleteStateLogRecords() error {
	schemas, err := uc.logSchemaRepository.GetAll()
	if err != nil {
		return fmt.Errorf("DeleteStateLogRecords: failed to query schemas: %w", err)
	}

	for _, schema := range schemas {
		// Treat 0 retention-time as infinite
		if schema.RetentionTime == 0 {
			continue
		}

		threshold := time.Now().Add(-time.Duration(schema.RetentionTime) * 24 * time.Hour)

		staleLogRecords, err := uc.logRecordRepository.FindStale(schema.ID, threshold)
		if err != nil {
			return fmt.Errorf("DeleteStateLogRecords: failed to query stale log records: %w", err)
		}

		var staleLogRecordIDs []int
		for _, logRecord := range staleLogRecords {
			staleLogRecordIDs = append(staleLogRecordIDs, logRecord.ID)
		}

		logRecordIDsBatches := shared.BatchSlice(staleLogRecordIDs, 500)

		for _, logRecordIDs := range logRecordIDsBatches {
			err = uc.logRecordRepository.DeleteByIDs(logRecordIDs)

			if err != nil {
				return fmt.Errorf("DeleteStateLogRecords: failed to delete stale log records: %w", err)
			}
		}
	}

	return nil
}
