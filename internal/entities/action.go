package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type Action struct {
	ID         int        `json:"id" sql:"id" bson:"id"`
	Name       string     `json:"name" sql:"name" bson:"name"`
	Pattern    string     `json:"pattern" sql:"pattern" bson:"pattern"`
	Method     string     `json:"method" sql:"method" bson:"method"`
	Conditions Conditions `json:"conditions" sql:"conditions" bson:"conditions"`
	SchemaName string     `json:"schema_name" sql:"schema_name" bson:"schema_name"`
	SchemaID   int        `json:"schema_id,omitempty" sql:"schema_id" bson:"schema_id"`
	Disabled   BoolAsInt  `json:"disabled" sql:"disabled" bson:"disabled"`
}

type Condition struct {
	Attribute string `json:"attribute" sql:"attribute" bson:"attribute"`
	Operation string `json:"operation" sql:"operation" bson:"operation"`
	Value     string `json:"value" sql:"value" bson:"value"`
}

type Conditions []Condition

func (fs *Conditions) Scan(value interface{}) error {
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

func (fs Conditions) Value() (driver.Value, error) {
	if fs == nil {
		fs = make([]Condition, 0)
	}

	b, err := serialize.JSONMarshal(fs)
	return string(b), err
}

func (c Conditions) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if c == nil {
		c = make(Conditions, 0)
	}
	return bson.MarshalValue([]Condition(c))
}

type ActionMethod string

const (
	GETActionMethod ActionMethod = "GET"
)

var AppropriateActions = []string{
	string(GETActionMethod),
}

var AppropriateActionsInfo = strings.Join(AppropriateActions, ", ")
