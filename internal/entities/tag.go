package entities

import (
	"database/sql/driver"
	"fmt"
	"moonlogs/internal/lib/serialize"
)

type Tag struct {
	ID   int    `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
}

type Tags []int

func (t *Tags) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		return serialize.JSONUnmarshal([]byte(v), t)
	case []byte:
		return serialize.JSONUnmarshal(v, t)
	default:
		return fmt.Errorf("unsupported type for Tags: %T", value)
	}
}

func (t Tags) Value() (driver.Value, error) {
	b, err := serialize.JSONMarshal(t)

	return string(b), err
}
