package entities

import (
	"database/sql/driver"
	"fmt"
)

type ApiToken struct {
	ID          int       `json:"id" sql:"id" bson:"id"`
	Token       string    `json:"token" sql:"token" bson:"token"`
	TokenDigest string    `json:"-" sql:"token_digest" bson:"token_digest"`
	Name        string    `json:"name" sql:"name" bson:"name"`
	IsRevoked   BoolAsInt `json:"is_revoked" sql:"is_revoked" bson:"is_revoked"`
}

type BoolAsInt bool

func (b *BoolAsInt) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*b = BoolAsInt(v != 0)
	case nil:
		*b = false
	default:
		return fmt.Errorf("unexpected type %T for BoolAsInt", value)
	}

	return nil
}

func (b BoolAsInt) Value() (driver.Value, error) {
	if b {
		return int64(1), nil
	}

	return int64(0), nil
}
