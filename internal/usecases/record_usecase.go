package usecases

import (
	"encoding/json"
	"fmt"
	"hash"
	"hash/fnv"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"slices"
	"strings"
	"sync"
	"time"
)

var FNVHasherPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

type RecordUseCase struct {
	recordStorage storage.RecordStorage
}

func NewRecordUseCase(recordStorage storage.RecordStorage) *RecordUseCase {
	return &RecordUseCase{recordStorage: recordStorage}
}

func (uc *RecordUseCase) CreateRecord(record entities.Record, schemaID int) (*entities.Record, error) {
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

	bytes, err := json.Marshal(record.Query)
	if err != nil {
		return nil, fmt.Errorf("failed creating record: %v", err)
	}

	FNV64Hasher := FNVHasherPool.Get().(hash.Hash64)
	defer FNVHasherPool.Put(FNV64Hasher)

	FNV64Hasher.Write(bytes)
	hashSum := FNV64Hasher.Sum64()
	FNV64Hasher.Reset()

	groupHash := fmt.Sprint(hashSum)

	return uc.recordStorage.CreateRecord(record, schemaID, groupHash)
}

func (uc *RecordUseCase) DeleteStaleRecords(schema *entities.Schema) error {
	// Treat 0 retention days as infinite
	if schema.RetentionDays == 0 {
		return nil
	}

	threshold := time.Now().Add(-time.Duration(schema.RetentionDays) * 24 * time.Hour).Unix()

	staleRecords, err := uc.recordStorage.FindStale(schema.ID, threshold)
	if err != nil {
		return fmt.Errorf("DeleteStaleRecords: failed to query stale log records: %w", err)
	}

	var staleRecordIDs []int
	for _, record := range staleRecords {
		staleRecordIDs = append(staleRecordIDs, record.ID)
	}

	recordIDsBatches := shared.BatchSlice(staleRecordIDs, 500)

	for _, recordIDs := range recordIDsBatches {
		err = uc.recordStorage.DestroyByIDs(recordIDs)

		if err != nil {
			return fmt.Errorf("DeleteStatelecords: failed to delete stale log records: %w", err)
		}
	}

	return nil
}

func (uc *RecordUseCase) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	return uc.recordStorage.GetAllRecords(limit, offset)
}

func (uc *RecordUseCase) GetAllRecordsCount() (int, error) {
	return uc.recordStorage.GetAllRecordsCount()
}

func (uc *RecordUseCase) GetRecordByID(id int) (*entities.Record, error) {
	return uc.recordStorage.GetRecordByID(id)
}

func (uc *RecordUseCase) GetRecordsByQuery(record entities.Record, limit int, offset int) ([]*entities.Record, error) {
	return uc.recordStorage.GetRecordsByQuery(record, limit, offset)
}

func (uc *RecordUseCase) GetRecordsCountByQuery(record entities.Record) (int, error) {
	return uc.recordStorage.GetRecordsCountByQuery(record)
}

func (uc *RecordUseCase) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	return uc.recordStorage.GetRecordsByGroupHash(schemaName, groupHash)
}
