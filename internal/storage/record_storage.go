package storage

import (
	"context"
	"moonlogs/internal/entities"
	"time"
)

type RecordStorage interface {
	CreateRecord(ctx context.Context, record entities.Record) (*entities.Record, error)
	DeleteByIDs(ctx context.Context, ids []int) error
	FindStaleIDs(ctx context.Context, schemaID int, threshold int64) ([]int, error)
	GetAllRecords(ctx context.Context, limit int, offset int) ([]*entities.Record, error)
	GetAllRecordsCount(ctx context.Context) (int, error)
	GetRecordByID(ctx context.Context, id int) (*entities.Record, error)
	GetRecordsByGroupHash(ctx context.Context, schemaName string, groupHash string) ([]*entities.Record, error)
	GetRecordsByQuery(ctx context.Context, record entities.Record, from *time.Time, to *time.Time, limit int, offset int) ([]*entities.Record, int, error)
	AggregateRecords(ctx context.Context, filter entities.RecordFilter, aggregation entities.RecordAggregation) ([]*entities.AggregationGroup, error)
}
