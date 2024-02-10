package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"
)

type Schema struct {
	ID            int    `json:"id" sql:"id"`
	Title         string `json:"title" sql:"title"`
	Description   string `json:"description" sql:"description"`
	Name          string `json:"name" sql:"name"`
	Fields        Fields `json:"fields" sql:"fields"`
	Kinds         Kinds  `json:"kinds" sql:"kinds"`
	TagID         int    `json:"tag_id,omitempty" sql:"tag_id"`
	RetentionDays int64  `json:"retention_days,omitempty" sql:"retention_days"`
}

type Field struct {
	Title string `json:"title" sql:"title"`
	Name  string `json:"name" sql:"name"`
}
type Fields []Field

type Kind struct {
	Title string `json:"title" sql:"title"`
	Name  string `json:"name" sql:"name"`
}
type Kinds []Kind

func (fs *Fields) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return serialize.JSONUnmarshal(v, fs)
	case string:
		return serialize.JSONUnmarshal([]byte(v), fs)
	default:
		return fmt.Errorf("unsupported type for Fields.Scan")
	}
}

func (fs Fields) Value() (driver.Value, error) {
	if fs == nil {
		fs = make([]Field, 0)
	}

	b, err := serialize.JSONMarshal(fs)
	return string(b), err
}

func (fs *Kinds) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return serialize.JSONUnmarshal(v, fs)
	case string:
		return serialize.JSONUnmarshal([]byte(v), fs)
	default:
		return fmt.Errorf("unsupported type for Kinds.Scan")
	}
}

func (fs Kinds) Value() (driver.Value, error) {
	if fs == nil {
		fs = make([]Kind, 0)
	}

	b, err := serialize.JSONMarshal(fs)
	return string(b), err
}
