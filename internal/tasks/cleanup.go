package tasks

import (
	"log"
	"moonlogs/usecase"
	"time"
)

func RunCleanupTask(logRecordUseCase *usecase.LogRecordUseCase, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		err := logRecordUseCase.DeleteStateLogRecords()

		if err != nil {
			log.Printf("Error cleaning up stale log records: %v", err)
		}
	}
}
