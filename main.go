package main

import (
	"context"
	"embed"
	"log"
	"moonlogs/internal/api/server"
	"moonlogs/internal/config"
	"moonlogs/internal/persistence"
	"moonlogs/internal/tasks"
	"moonlogs/internal/usecases"
	"time"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	databases, err := persistence.InitDB(cfg.DBAdapter, cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tasks.Migrate(cfg.DBAdapter, databases, embedMigrations)
	if err != nil {
		log.Fatal(err)
	}

	storageInstances := persistence.InitStorages(cfg.DBAdapter, databases)
	usecaseInstances := usecases.InitUsecases(storageInstances)

	runCleanupTasks(context.Background(), usecaseInstances)

	err = server.ListenAndServe(
		usecaseInstances,
		server.WithPort(cfg.Port),
		server.WithReadTimeout(cfg.ReadTimeout),
		server.WithWriteTimeout(cfg.WriteTimeout),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func runCleanupTasks(ctx context.Context, uc *usecases.UseCases) {
	go tasks.RunRecordsCleanupTask(ctx, 1*time.Hour, uc)
}
