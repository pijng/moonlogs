package repository

import (
	"context"
	"fmt"
	"moonlogs/ent"
	"moonlogs/ent/logrecord"
	"moonlogs/ent/predicate"
	"moonlogs/ent/schema"
	"moonlogs/internal/config"

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

func (r *LogRecordRepository) Create(logRecord ent.LogRecord, logSchemaId int) (*ent.LogRecord, error) {
	lr, err := r.client.LogRecord.
		Create().
		SetText(logRecord.Text).
		SetSchemaName(logRecord.SchemaName).
		SetSchemaID(logSchemaId).
		SetMeta(logRecord.Meta).
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

func (r *LogRecordRepository) GetBySchemaAndMeta(schemaId int, schemaName string, meta schema.Meta, limit int, offset int) ([]*ent.LogRecord, error) {
	lr, err := r.client.Debug().LogRecord.
		Query().
		Where(logrecord.Or(logrecord.SchemaID(schemaId), logrecord.SchemaName(schemaName))).
		Where(predicate.LogRecord(func(s *sql.Selector) {
			for name, value := range meta {
				s.Where(sqljson.ValueContains(logrecord.FieldMeta, value, sqljson.Path(name)))
			}
		})).
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
		Query().Limit(limit).Offset(offset).All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying log_record: %w", err)
	}

	return lr, nil
}
