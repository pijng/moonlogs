package main

import (
	"context"
	"flag"
	"log"
	"moonlogs/api/server"
	"moonlogs/ent"
	"moonlogs/internal/config"
	"moonlogs/internal/repository"
	"moonlogs/internal/schema"
	"moonlogs/internal/tasks"
	"moonlogs/usecase"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	developmentFlag = flag.Bool("development", true, "Development mode")
)

func main() {
	flag.Parse()

	client, err := ent.Open("sqlite3", "file:./database.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if *developmentFlag {
		schema.Generate(client)
	}

	config.SetClient(client)

	runTasks(context.Background())

	server.Serve()
}

func runTasks(ctx context.Context) {
	logRecordRepository := repository.NewLogRecordRepository(ctx)
	logSchemaRepository := repository.NewLogSchemaRepository(ctx)
	logRecordUseCase := usecase.NewLogRecordsUseCase(logRecordRepository, logSchemaRepository)

	cleanupInterval := 1 * time.Hour
	go tasks.RunCleanupTask(logRecordUseCase, cleanupInterval)
}
