package tasks

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/usecases"
	"time"
)

func RunRecordsCleanupTask(ctx context.Context, interval time.Duration, uc *usecases.UseCases) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		schemas, err := uc.SchemaUseCase.GetAllSchemas(ctx, &entities.User{})
		if err != nil {
			log.Printf("failed getting log schemas: %v\n", err)
			continue
		}

		for _, schema := range schemas {
			err = uc.RecordUseCase.DeleteStaleRecords(ctx, schema)
			if err != nil {
				log.Printf("failed cleaning up stale log records: %v\n", err)
				continue
			}
		}
	}
}

func RunIncidentsCleanupTask(ctx context.Context, interval time.Duration, uc *usecases.UseCases) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		err := uc.IncidentUseCase.DeleteStaleIncidents(ctx)
		if err != nil {
			log.Printf("failed cleaning up stale incidents: %v\n", err)
			continue
		}
	}
}

func RunStatementsCleanupTask(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			qrx.CleanCachedStatements()
		case <-ctx.Done():
			log.Println("cached statement cleanup task canceled")
			return
		}
	}
}
