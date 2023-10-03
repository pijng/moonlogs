package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Field struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

// LogSchema holds the schema definition for the LogSchema entity.
type LogSchema struct {
	ent.Schema
}

// Fields of the LogSchema.
func (LogSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("field is required")
			}
			return nil
		}),
		field.String("description"),
		field.String("name").Validate(func(s string) error {
			if len(s) == 0 {
				return errors.New("field is required")
			}
			return nil
		}).Unique().Immutable(),
		field.JSON("fields", []Field{}),
		field.Int64("retention_time").Optional(),
	}
}

// Edges of the LogSchema.
func (LogSchema) Edges() []ent.Edge {
	return nil
}
