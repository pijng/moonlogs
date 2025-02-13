package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type AlertingRuleStorage interface {
	CreateRule(ctx context.Context, rule entities.AlertingRule) (*entities.AlertingRule, error)
	DeleteRuleByID(ctx context.Context, id int) error
	GetAllRules(ctx context.Context) ([]*entities.AlertingRule, error)
	GetRuleByID(ctx context.Context, id int) (*entities.AlertingRule, error)
	UpdateRuleByID(ctx context.Context, id int, rule entities.AlertingRule) (*entities.AlertingRule, error)
}
