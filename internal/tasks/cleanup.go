package tasks

import (
	"context"
	"log"
	"moonlogs/internal/persistence"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"time"
)

func RunCleanupTask(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	schemaRepository := repositories.NewSchemaRepository(ctx)
	recordRepository := repositories.NewRecordRepository(ctx)
	recordUseCase := usecases.NewRecordUseCase(recordRepository)

	for range ticker.C {
		schemas, err := schemaRepository.GetAllSchemas()
		if err != nil {
			log.Printf("failed getting log schemas: %v", err)
			continue
		}

		for _, schema := range schemas {
			err = recordUseCase.DeleteStaleRecords(schema)
			if err != nil {
				log.Printf("failed cleaning up stale log records: %v", err)
			}

			_, err = persistence.DB().ExecContext(ctx, "ANALYZE;")
			if err != nil {
				log.Printf("failed optimizing db's query planner statistics: %v", err)
			}

			_, err = persistence.DB().ExecContext(ctx, "VACUUM;")
			if err != nil {
				log.Printf("failed vacuuming db: %v", err)
			}
		}
	}
}
