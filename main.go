package main

import (
	"context"
	"log"
	"moonlogs/api/server"
	"moonlogs/internal/config"
	"moonlogs/internal/persistence"
	"moonlogs/internal/tasks"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = persistence.InitDB(config.DB_PATH)
	if err != nil {
		log.Fatal(err)
	}

	runTasks(context.Background())

	err = server.ListenAndServe(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func runTasks(ctx context.Context) {
	go tasks.RunRecordsCleanupTask(ctx, 1*time.Hour)
	go tasks.RunStatementsCleanupTask(ctx, 15*time.Minute)
}
