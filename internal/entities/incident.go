package entities

type Incident struct {
	ID     int          `json:"id" sql:"id" bson:"id"`
	RuleID int          `json:"rule_id" sql:"rule_id" bson:"rule_id"`
	Keys   JSONMap[any] `json:"keys" sql:"keys" bson:"keys"`
	Count  int          `json:"count" sql:"count" bson:"count"`
	TTL    RecordTime   `json:"ttl" sql:"ttl" bson:"ttl"`
}
