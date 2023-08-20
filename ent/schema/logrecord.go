package schema

import (
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Meta map[string]interface{}

// LogRecord holds the schema definition for the LogRecord entity.
type LogRecord struct {
	ent.Schema
}

// Fields of the LogRecord.
func (LogRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("field is required")
			}
			return nil
		}).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.String("schema_name").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("field is required")
			}
			return nil
		}).Immutable(),
		field.Int("schema_id").Validate(func(i int) error {
			if i == 0 {
				return errors.New("field is required")
			}
			return nil
		}).Immutable(),
		field.JSON("meta", Meta{}),
	}
}

// Edges of the LogRecord.
func (LogRecord) Edges() []ent.Edge {
	return nil
}
