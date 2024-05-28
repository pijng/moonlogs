package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type ActionStorage interface {
	CreateAction(ctx context.Context, action entities.Action) (*entities.Action, error)
	DeleteActionByID(ctx context.Context, id int) error
	GetAllActions(ctx context.Context) ([]*entities.Action, error)
	GetActionByID(ctx context.Context, id int) (*entities.Action, error)
	UpdateActionByID(ctx context.Context, id int, action entities.Action) (*entities.Action, error)
}
