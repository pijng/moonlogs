package usecases

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"time"
)

type IncidentUseCase struct {
	incidentStorage storage.IncidentStorage
}

func NewIncidentUseCase(incidentStorage storage.IncidentStorage) *IncidentUseCase {
	return &IncidentUseCase{incidentStorage: incidentStorage}
}

func (uc *IncidentUseCase) CreateIncident(ctx context.Context, incident entities.Incident) (*entities.Incident, error) {
	existingIncidents, err := uc.incidentStorage.GetIncidentsByKeys(ctx, incident.Keys)
	if err != nil && !errors.Is(err, storage.ErrNotFound) {
		return nil, fmt.Errorf("failed querying incident by keys: %w", err)
	}

	if len(existingIncidents) > 0 {
		return existingIncidents[0], nil
	}

	return uc.incidentStorage.CreateIncident(ctx, incident)
}

func (uc *IncidentUseCase) DeleteStaleIncidents(ctx context.Context) error {
	threshold := time.Now()

	staleIncidentIDs, err := uc.incidentStorage.FindStaleIDs(ctx, threshold.UnixMilli())
	if err != nil {
		return fmt.Errorf("DeleteStaleIncidents: failed to query stale incidents: %w", err)
	}

	staleIncidentIDsBatches := shared.BatchSlice(staleIncidentIDs, 950)

	for _, ids := range staleIncidentIDsBatches {
		err = uc.incidentStorage.DeleteByIDs(ctx, ids)

		if err != nil {
			return fmt.Errorf("DeleteStaleIncidents: failed to delete stale incidents: %w", err)
		}
	}

	return nil
}

func (uc *IncidentUseCase) GetAllIncidents(ctx context.Context) ([]*entities.Incident, error) {
	return uc.incidentStorage.GetAllIncidents(ctx)
}

func (uc *IncidentUseCase) GetIncidentsByKeys(ctx context.Context, keys entities.JSONMap[any]) ([]*entities.Incident, error) {
	return uc.incidentStorage.GetIncidentsByKeys(ctx, keys)
}
