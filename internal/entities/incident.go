package entities

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"moonlogs/internal/lib/serialize"
	"slices"
	"strings"
	"text/template"
)

type Incident struct {
	ID       int          `json:"id" sql:"id" bson:"id"`
	RuleName string       `json:"rule_name" sql:"rule_name" bson:"rule_name"`
	RuleID   int          `json:"rule_id" sql:"rule_id" bson:"rule_id"`
	Keys     JSONMap[any] `json:"keys" sql:"keys" bson:"keys"`
	Count    int          `json:"count" sql:"count" bson:"count"`
	TTL      RecordTime   `json:"ttl" sql:"ttl" bson:"ttl"`
}

func (i *Incident) Message(payload string, timeWindow Duration) (string, error) {
	tmpl, err := template.New("incident").Parse(payload)
	if err != nil {
		return "", fmt.Errorf("failed to parse template for rule:'%v': %w", i.RuleName, err)
	}

	data := struct {
		RuleName   string
		Keys       string
		Count      int
		TimeWindow Duration
	}{
		RuleName:   i.RuleName,
		Keys:       mapToString(i.Keys),
		Count:      i.Count,
		TimeWindow: timeWindow,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template for rule:'%v': %w", i.RuleName, err)
	}

	return buf.String(), nil
}

func (i *Incident) Hash() (string, error) {
	keysJSON, err := serialize.JSONMarshal(i.Keys)
	if err != nil {
		return "", fmt.Errorf("failed to marshal incident keys: %w", err)
	}

	data := fmt.Sprintf("%s:%s", i.RuleName, string(keysJSON))
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:]), nil
}

func mapToString(m JSONMap[any]) string {
	parts := make([]string, 0, len(m))
	for k, v := range m {
		parts = append(parts, fmt.Sprintf("%s: %s", k, v))
	}
	slices.Sort(parts)
	return strings.Join(parts, ", ")
}
