package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"
	"time"
)

type Record struct {
	ID         int        `json:"id" sql:"id"`
	Text       string     `json:"text" sql:"text"`
	CreatedAt  RecordTime `json:"created_at" sql:"created_at"`
	SchemaName string     `json:"schema_name,omitempty" sql:"schema_name"`
	SchemaID   int        `json:"schema_id,omitempty" sql:"schema_id"`
	Query      JSONMap    `json:"query,omitempty" sql:"query"`
	Kind       string     `json:"kind,omitempty" sql:"kind"`
	GroupHash  string     `json:"group_hash,omitempty" sql:"group_hash"`
	Level      Level      `json:"level,omitempty" sql:"level"`
}

type JSONMap map[string]interface{}

func (jm *JSONMap) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return serialize.JSONUnmarshal(v, jm)
	case string:
		return serialize.JSONUnmarshal([]byte(v), jm)
	default:
		return fmt.Errorf("unsupported type for Query.Scan")
	}
}

func (jm JSONMap) Value() (driver.Value, error) {
	if jm == nil {
		jm = make(JSONMap)
	}

	b, err := serialize.JSONMarshal(jm)
	return string(b), err
}

type RecordTime struct {
	time.Time
}

func (t *RecordTime) Scan(value interface{}) error {
	if v, ok := value.(int64); ok {
		t.Time = time.Unix(v, 0)
	}

	return nil
}

func (t RecordTime) Value() (driver.Value, error) {
	return t.Unix(), nil
}

type Level string

const (
	TraceLevel Level = "Trace"
	DebugLevel Level = "Debug"
	InfoLevel  Level = "Info"
	WarnLevel  Level = "Warn"
	ErrorLevel Level = "Error"
	FatalLevel Level = "Fatal"
)

var AppropriateLevels = []string{
	string(TraceLevel),
	string(DebugLevel),
	string(InfoLevel),
	string(WarnLevel),
	string(ErrorLevel),
	string(FatalLevel),
}
