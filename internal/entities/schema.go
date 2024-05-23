package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type Schema struct {
	ID            int    `json:"id" sql:"id" bson:"id"`
	Title         string `json:"title" sql:"title" bson:"title"`
	Description   string `json:"description" sql:"description" bson:"description"`
	Name          string `json:"name" sql:"name" bson:"name"`
	Fields        Fields `json:"fields" sql:"fields" bson:"fields"`
	Kinds         Kinds  `json:"kinds" sql:"kinds" bson:"kinds"`
	TagID         int    `json:"tag_id,omitempty" sql:"tag_id" bson:"tag_id"`
	RetentionDays int64  `json:"retention_days,omitempty" sql:"retention_days" bson:"retention_days"`
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

func (fs Fields) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if fs == nil {
		fs = make([]Field, 0)
	}
	return bson.MarshalValue([]Field(fs))
}

func (k Kinds) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if k == nil {
		k = make([]Kind, 0)
	}
	return bson.MarshalValue([]Kind(k))
}
