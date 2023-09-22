package main

import (
	"flag"
	"log"
	"moonlogs/api/server"
	"moonlogs/ent"
	"moonlogs/internal/config"
	"moonlogs/internal/schema"

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

	server.Serve()
}
