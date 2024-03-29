package main

import (
	"context"
	"embed"
	"log"
	"moonlogs/internal/api/server"
	"moonlogs/internal/config"
	"moonlogs/internal/persistence"
	"moonlogs/internal/services"
	"moonlogs/internal/tasks"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.PyroscopeProfiling {
		err = services.StartPyroscope(cfg.PyroscopeAddress)
		if err != nil {
			log.Fatal(err)
		}
	}

	var nrapp *newrelic.Application
	if cfg.NewrelicProfiling {
		nrapp, err = services.StartNewrelic(cfg.NewrelicLicense)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = persistence.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = tasks.Migrate(cfg.DBAdapter, embedMigrations)
	if err != nil {
		log.Fatal(err)
	}

	runCleanupTasks(context.Background())

	err = server.ListenAndServe(cfg, nrapp)
	if err != nil {
		log.Fatal(err)
	}
}

func runCleanupTasks(ctx context.Context) {
	go tasks.RunRecordsCleanupTask(ctx, 1*time.Hour)
	// Not used anymore
	// go tasks.RunStatementsCleanupTask(ctx, 15*time.Minute)
}
