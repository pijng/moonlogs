package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"moonlogs/internal/lib/serialize"
	"time"
)

type AlertingRule struct {
	ID                    int             `json:"id" sql:"id" bson:"id"`
	Name                  string          `json:"name" sql:"name" bson:"name"`
	Description           string          `json:"description" sql:"description" bson:"description"`
	Enabled               BoolAsInt       `json:"enabled" sql:"enabled" bson:"enabled"`
	Severity              Level           `json:"severity" sql:"severity" bson:"severity"`
	Interval              Duration        `json:"interval" sql:"interval" bson:"interval"`
	Threshold             int             `json:"threshold" sql:"threshold" bson:"threshold"`
	Condition             AlertCondition  `json:"condition" sql:"condition" bson:"condition"`
	FilterLevel           Level           `json:"filter_level,omitempty" sql:"filter_level" bson:"filter_level"`
	FilterSchemaIDs       SchemaIDs       `json:"filter_schema_ids" sql:"filter_schema_ids" bson:"filter_schema_ids"`
	FilterSchemaFields    StringArray     `json:"filter_schema_fields" sql:"filter_schema_fields" bson:"filter_schema_fields"`
	FilterSchemaKinds     StringArray     `json:"filter_schema_kinds" sql:"filter_schema_kinds" bson:"filter_schema_kinds"`
	AggregationType       AggregationType `json:"aggregation_type" sql:"aggregation_type" bson:"aggregation_type"`
	AggregationGroupBy    StringArray     `json:"aggregation_group_by" sql:"aggregation_group_by" bson:"aggregation_group_by"`
	AggregationTimeWindow Duration        `json:"aggregation_time_window" sql:"aggregation_time_window" bson:"aggregation_time_window"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		sd := string(b[1 : len(b)-1])
		d.Duration, err = time.ParseDuration(sd)
		return
	}

	var id int64
	id, err = json.Number(string(b)).Int64()
	d.Duration = time.Duration(id)

	return
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}

type AggregationType string

var AppropriateAggregationTypes = []string{
	string(CountAggregationType),
}

const (
	CountAggregationType AggregationType = "count"
)

type StringArray []string

func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		return serialize.JSONUnmarshal([]byte(v), sa)
	case []byte:
		return serialize.JSONUnmarshal(v, sa)
	default:
		return fmt.Errorf("unsupported type for SchemaFields: %T", value)
	}
}

func (sa StringArray) Value() (driver.Value, error) {
	b, err := serialize.JSONMarshal(sa)

	return string(b), err
}

type AlertCondition string

var AppropriateAlertConditions = []string{
	string(LTAlertCondition),
	string(GTAlertCondition),
	string(EQAlertCondition),
}

const (
	LTAlertCondition AlertCondition = "<"
	GTAlertCondition AlertCondition = ">"
	EQAlertCondition AlertCondition = "=="
)
