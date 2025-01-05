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
	SchemaIDs  SchemaIDs  `json:"schema_ids" sql:"schema_ids" bson:"schema_ids"`
	Disabled   BoolAsInt  `json:"disabled" sql:"disabled" bson:"disabled"`
}

type SchemaIDs []int

func (s *SchemaIDs) Scan(value interface{}) error {
	if value == nil {
		*s = make(SchemaIDs, 0)
		return nil
	}

	switch v := value.(type) {
	case string:
		return serialize.JSONUnmarshal([]byte(v), s)
	case []byte:
		return serialize.JSONUnmarshal(v, s)
	default:
		return fmt.Errorf("unsupported type for SchemaIDs: %T", value)
	}
}

func (s SchemaIDs) Value() (driver.Value, error) {
	b, err := serialize.JSONMarshal(s)

	return string(b), err
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

type ConditionOperation string

const (
	EQ     ConditionOperation = "=="
	NE     ConditionOperation = "!="
	LT     ConditionOperation = "<"
	LTE    ConditionOperation = "<="
	GT     ConditionOperation = ">"
	GTE    ConditionOperation = ">="
	EXISTS ConditionOperation = "EXISTS"
	EMPTY  ConditionOperation = "EMPTY"
)

var AppropriateOperations = []string{
	string(EQ),
	string(NE),
	string(LT),
	string(LTE),
	string(GT),
	string(GTE),
	string(EXISTS),
	string(EMPTY),
}

var AppropriateOperationsInfo = strings.Join(AppropriateOperations, ", ")
