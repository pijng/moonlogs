package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type IncidentStorage interface {
	CreateIncident(ctx context.Context, incident entities.Incident) (*entities.Incident, error)
	FindStaleIDs(ctx context.Context, threshold int64) ([]int, error)
	DeleteByIDs(ctx context.Context, ids []int) error
	GetAllIncidents(ctx context.Context) ([]*entities.Incident, error)
	GetIncidentsByKeys(ctx context.Context, keys entities.JSONMap[any]) ([]*entities.Incident, error)
}
