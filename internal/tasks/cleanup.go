package tasks

import (
	"context"
	"log"
	"moonlogs/internal/config"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"time"
)

func RunRecordsCleanupTask(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	schemaStorage := storage.NewSchemaStorage(ctx, config.Get().DBAdapter)
	recordStorage := storage.NewRecordStorage(ctx, config.Get().DBAdapter)
	recordUseCase := usecases.NewRecordUseCase(recordStorage)

	for range ticker.C {
		schemas, err := schemaStorage.GetAllSchemas()
		if err != nil {
			log.Printf("failed getting log schemas: %v", err)
			continue
		}

		for _, schema := range schemas {
			err = recordUseCase.DeleteStaleRecords(schema)
			if err != nil {
				log.Printf("failed cleaning up stale log records: %v", err)
				continue
			}

			_, err = persistence.DB().ExecContext(ctx, "ANALYZE;")
			if err != nil {
				log.Printf("failed optimizing db's query planner statistics: %v", err)
				continue
			}

			_, err = persistence.DB().ExecContext(ctx, "VACUUM;")
			if err != nil {
				log.Printf("failed vacuuming db: %v", err)
			}
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
			log.Printf("cached statement cleanup task canceled")
		}
	}
}
