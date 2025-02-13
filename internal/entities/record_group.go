package entities

import (
	"time"
)

type AggregationGroup struct {
	Keys  JSONMap[any]
	Count int32
}

type RecordFilter struct {
	Level        Level
	SchemaIDs    []int
	SchemaFields []string
	SchemaKinds  []string
	From         time.Time
	To           time.Time
}

type RecordAggregation struct {
	GroupBy []string
	Type    AggregationType
}
