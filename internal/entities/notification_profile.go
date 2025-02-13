package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type NotificationProfile struct {
	ID          int       `json:"id" sql:"id" bson:"id"`
	Name        string    `json:"name" sql:"name" bson:"name"`
	Description string    `json:"description" sql:"description" bson:"description"`
	RuleIDs     RuleIDs   `json:"rule_ids" sql:"rule_ids" bson:"rule_ids"`
	Enabled     BoolAsInt `json:"enabled" sql:"enabled" bson:"enabled"`
	URL         string    `json:"url" sql:"url" bson:"url"`
	Method      string    `json:"method" sql:"method" bson:"method"`
	Headers     Headers   `json:"headers" sql:"headers" bson:"headers"`
	Payload     string    `json:"payload" sql:"payload" bson:"payload"`
}
type RuleIDs []int

func (s *RuleIDs) Scan(value interface{}) error {
	if value == nil {
		*s = make(RuleIDs, 0)
		return nil
	}

	switch v := value.(type) {
	case string:
		return serialize.JSONUnmarshal([]byte(v), s)
	case []byte:
		return serialize.JSONUnmarshal(v, s)
	default:
		return fmt.Errorf("unsupported type for RuleID: %T", value)
	}
}

func (s RuleIDs) Value() (driver.Value, error) {
	b, err := serialize.JSONMarshal(s)

	return string(b), err
}

type Header struct {
	Key   string `json:"key" sql:"key" bson:"key"`
	Value string `json:"value" sql:"value" bson:"value"`
}

type Headers []Header

func (fs *Headers) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return serialize.JSONUnmarshal(v, fs)
	case string:
		return serialize.JSONUnmarshal([]byte(v), fs)
	default:
		return fmt.Errorf("unsupported type for Conditions.Scan")
	}
}

func (fs Headers) Value() (driver.Value, error) {
	if fs == nil {
		fs = make([]Header, 0)
	}

	b, err := serialize.JSONMarshal(fs)
	return string(b), err
}

func (c Headers) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if c == nil {
		c = make(Headers, 0)
	}
	return bson.MarshalValue([]Header(c))
}
