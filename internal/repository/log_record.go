package repository

import (
	"context"
	"fmt"
	"moonlogs/ent"
	"moonlogs/ent/logrecord"
	"moonlogs/ent/predicate"
	"moonlogs/ent/schema"
	"moonlogs/internal/config"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
)

type LogRecordRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewLogRecordRepository(ctx context.Context) *LogRecordRepository {
	return &LogRecordRepository{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *LogRecordRepository) Create(logRecord ent.LogRecord, logSchemaId int, groupHash string) (*ent.LogRecord, error) {
	lr, err := r.client.LogRecord.
		Create().
		SetText(logRecord.Text).
		SetSchemaName(logRecord.SchemaName).
		SetSchemaID(logSchemaId).
		SetQuery(logRecord.Query).
		SetGroupHash(groupHash).
		SetLevel(logRecord.Level).
		Save(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating log_record: %w", err)
	}

	return lr, nil
}

func (r *LogRecordRepository) GetById(id int) (*ent.LogRecord, error) {
	lr, err := r.client.LogRecord.
		Query().
		Where(logrecord.ID(id)).First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_record: %w", err)
	}

	return lr, nil
}

func (r *LogRecordRepository) GetBySchemaAndQuery(schemaId int, schemaName string, text string, query schema.Query, limit int, offset int) ([]*ent.LogRecord, error) {
	lr, err := r.client.Debug().LogRecord.
		Query().
		Where(logrecord.Or(logrecord.SchemaID(schemaId), logrecord.SchemaName(schemaName)), logrecord.TextContains(text)).
		Where(predicate.LogRecord(func(s *sql.Selector) {
			for name, value := range query {
				s.Where(sqljson.ValueContains(logrecord.FieldQuery, value, sqljson.Path(name)))
			}
		})).
		Order(ent.Desc(logrecord.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_record: %w", err)
	}

	return lr, nil
}

func (r *LogRecordRepository) GetAll(limit int, offset int) ([]*ent.LogRecord, error) {
	lr, err := r.client.LogRecord.
		Query().Order(ent.Desc(logrecord.FieldCreatedAt)).Limit(limit).Offset(offset).All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_record: %w", err)
	}

	return lr, nil
}

func (r *LogRecordRepository) GetByGroupHash(schemaName string, groupHash string) ([]*ent.LogRecord, error) {
	lr, err := r.client.Debug().LogRecord.
		Query().
		Where(logrecord.SchemaName(schemaName)).
		Where(logrecord.GroupHash(groupHash)).
		Order(ent.Desc(logrecord.FieldCreatedAt)).
		All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_record: %w", err)
	}

	return lr, nil
}

func (r *LogRecordRepository) GetCountBySchemaAndQuery(schemaName string, text string, query schema.Query) (int, error) {
	count, err := r.client.LogRecord.
		Query().
		Where(logrecord.SchemaName(schemaName), logrecord.TextContains(text)).
		Where(predicate.LogRecord(func(s *sql.Selector) {
			for name, value := range query {
				s.Where(sqljson.ValueContains(logrecord.FieldQuery, value, sqljson.Path(name)))
			}
		})).
		Count(r.ctx)

	if err != nil {
		return 0, fmt.Errorf("failed querying log_record: %w", err)
	}

	return count, nil
}

func (r *LogRecordRepository) GetCountAll() (int, error) {
	count, err := r.client.LogRecord.
		Query().Count(r.ctx)

	if err != nil {
		return 0, fmt.Errorf("failed querying log_record: %w", err)
	}

	return count, nil
}

func (r *LogRecordRepository) FindStale(schemaID int, threshold time.Time) ([]*ent.LogRecord, error) {
	return r.client.LogRecord.Query().Where(logrecord.SchemaID(schemaID), logrecord.CreatedAtLT(threshold)).All(r.ctx)
}

func (r *LogRecordRepository) DeleteByIDs(ids []int) error {
	_, err := r.client.LogRecord.Delete().Where(logrecord.IDIn(ids...)).Exec(r.ctx)

	return err
}
