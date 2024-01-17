package usecases

import (
	"encoding/json"
	"fmt"
	"hash"
	"hash/fnv"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/shared"
	"slices"
	"strings"
	"sync"
	"time"
)

var hasherPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

type RecordUseCase struct {
	recordRepository *repositories.RecordRepository
}

func NewRecordUseCase(recordRepository *repositories.RecordRepository) *RecordUseCase {
	return &RecordUseCase{recordRepository: recordRepository}
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

	FNV64Hasher := hasherPool.Get().(hash.Hash64)
	defer hasherPool.Put(FNV64Hasher)

	FNV64Hasher.Write(bytes)
	hashSum := FNV64Hasher.Sum64()
	FNV64Hasher.Reset()

	groupHash := fmt.Sprint(hashSum)

	return uc.recordRepository.CreateRecord(record, schemaID, groupHash)
}

func (uc *RecordUseCase) DeleteStaleRecords(schema *entities.Schema) error {
	// Treat 0 retention-time as infinite
	if schema.RetentionTime == 0 {
		return nil
	}

	threshold := time.Now().Add(-time.Duration(schema.RetentionTime) * 24 * time.Hour).Unix()

	staleRecords, err := uc.recordRepository.FindStale(schema.ID, threshold)
	if err != nil {
		return fmt.Errorf("DeleteStaleRecords: failed to query stale log records: %w", err)
	}

	var staleRecordIDs []int
	for _, record := range staleRecords {
		staleRecordIDs = append(staleRecordIDs, record.ID)
	}

	recordIDsBatches := shared.BatchSlice(staleRecordIDs, 500)

	for _, recordIDs := range recordIDsBatches {
		err = uc.recordRepository.DestroyByIDs(recordIDs)

		if err != nil {
			return fmt.Errorf("DeleteStatelecords: failed to delete stale log records: %w", err)
		}
	}

	return nil
}

func (uc *RecordUseCase) GetAllRecords(limit int, offset int) ([]*entities.Record, error) {
	return uc.recordRepository.GetAllRecords(limit, offset)
}

func (uc *RecordUseCase) GetAllRecordsCount() (int, error) {
	return uc.recordRepository.GetAllRecordsCount()
}

func (uc *RecordUseCase) GetRecordByID(id int) (*entities.Record, error) {
	return uc.recordRepository.GetRecordByID(id)
}

func (uc *RecordUseCase) GetRecordsByQuery(record entities.Record, limit int, offset int) ([]*entities.Record, error) {
	return uc.recordRepository.GetRecordsByQuery(record, limit, offset)
}

func (uc *RecordUseCase) GetRecordsCountByQuery(record entities.Record) (int, error) {
	return uc.recordRepository.GetRecordsCountByQuery(record)
}

func (uc *RecordUseCase) GetRecordsByGroupHash(schemaName string, groupHash string) ([]*entities.Record, error) {
	return uc.recordRepository.GetRecordsByGroupHash(schemaName, groupHash)
}
